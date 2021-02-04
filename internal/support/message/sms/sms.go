package sms

import (
	"culture/internal/model/message"
	"culture/internal/support/db"
	"culture/internal/support/redis"
	"culture/internal/util"
	"errors"
	"log"
	"strconv"
	"time"
)

// 验证码结构
type Code struct {
	Code   string `redis:"code"`
	Expire int64  `redis:"expire"`
	Time   int64  `redis:"time"`
}

// 发送验证码
func SendSmsCode(phoneNumber string, template string) error {
	var smsConfig message.SmsConfig
	db.DB.First(&smsConfig)
	if smsConfig.Id <= 0 {
		log.Println("SendSmsCode: 短信未配置")
		return errors.New("短信未配置")
	}

	var templateConfig message.SmsTemplate
	db.DB.Where("template=?", template).First(&templateConfig)
	if templateConfig.Id <= 0 {
		return errors.New("短信模版不存在")
	}

	templateCode := templateConfig.TemplateId
	templateParam := make(map[string]string)

	templateParam["code"] = util.GenValidateCode(6) // 验证码

	// 验证码缓存
	key := redis.Key("sms_code_" + template + ":" + phoneNumber)

	redisClient := redis.Client()

	redisClient.Del(key)

	smsCode := Code{
		Code:   templateParam["code"],
		Time:   time.Now().Unix(),
		Expire: int64(smsConfig.ValidCodeExpire),
	}

	redisClient.HMSet(key, map[string]interface{}{"Code": smsCode.Code, "Time": smsCode.Time, "Expire": smsCode.Expire})
	redisClient.Expire(key, time.Duration(smsConfig.ValidCodeExpire*60)*time.Second)

	var sender Sender

	switch smsConfig.SmsType {
	case "ali_cloud":
		sender = NewAliCloudSender(phoneNumber, templateCode, templateParam)
		break
	default:
		return errors.New("sms storage error:" + smsConfig.SmsType)
	}

	_, err := sender.Send()
	if err != nil {
		return err
	}

	return nil
}

// 验证短信验证码
func ValidSmsCode(phoneNumber string, template string, smsCode string) error {
	// 验证码缓存
	key := redis.Key("sms_code_" + template + ":" + phoneNumber)

	redisClient := redis.Client()

	code, err := redisClient.HGet(key, "Code").Result()
	if err != nil {
		log.Println(err)
		return errors.New("验证码已过期，请重新获取")
	}
	expire, _ := redisClient.HGet(key, "Expire").Result()
	codeTime, _ := redisClient.HGet(key, "Time").Result()

	expireInt, err := strconv.ParseInt(expire, 10, 64)
	if err != nil {
		log.Println(err)
		return errors.New("验证码已过期，请重新获取")
	}
	codeTimeInt, err := strconv.ParseInt(codeTime, 10, 64)
	if err != nil {
		log.Println(err)
		return errors.New("验证码已过期，请重新获取")
	}

	if codeTimeInt+expireInt*60 > time.Now().Unix() {
		if code != smsCode {
			return errors.New("验证码错误")
		}

		redisClient.Del(key)
	} else {
		return errors.New("验证码已过期，请重新获取")
	}

	return nil
}

// 发送普通短信消息
func SendSmsMessage(phoneNumber string, template string, param map[string]string) error {
	var smsConfig message.SmsConfig
	db.DB.First(&smsConfig)
	if smsConfig.Id <= 0 {
		log.Println("SendSmsCode: 短信未配置")
		return errors.New("短信未配置")
	}

	var templateConfig message.SmsTemplate
	db.DB.Where("template=?", template).First(&templateConfig)
	if templateConfig.Id <= 0 {
		return errors.New("短信模版不存在")
	}

	templateCode := templateConfig.TemplateId

	var sender Sender

	switch smsConfig.SmsType {
	case "ali_cloud":
		sender = NewAliCloudSender(phoneNumber, templateCode, param)
		break
	default:
		return errors.New("sms storage error:" + smsConfig.SmsType)
	}

	_, err := sender.Send()
	if err != nil {
		return err
	}

	return nil
}
