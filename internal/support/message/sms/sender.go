package sms

import (
	"culture/internal/model/message"
	"culture/internal/support/db"
	"encoding/json"
	"errors"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"log"
)

type Sender interface {
	// 发送短信
	Send() (bool, error)
}

// 阿里云短信
type AliCloudSender struct {
	phoneNumber   string
	signName      string
	templateCode  string
	templateParam string
	accessKey     string
	accessSecret  string
}

func NewAliCloudSender(phoneNumber string, templateCode string, templateParam map[string]string) *AliCloudSender {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()

	var smsConfig message.SmsConfig
	db.DB().First(&smsConfig)
	if smsConfig.Id <= 0 {
		log.Println("AliCloudSender: 未配置短信")
		panic("AliCloudSender: 未配置短信")
	}

	sender := new(AliCloudSender)

	sender.accessKey = smsConfig.AppKey
	sender.accessSecret = smsConfig.SecretKey
	sender.phoneNumber = phoneNumber
	sender.signName = smsConfig.FreeSignName
	sender.templateCode = templateCode
	param, err := json.Marshal(templateParam)
	if err != nil {
		log.Println("AliCloudSender: " + err.Error())
		panic(err)
	}
	sender.templateParam = string(param)

	return sender
}

// 发送短信
func (aliCloud *AliCloudSender) Send() (bool, error) {
	client, err := sdk.NewClientWithAccessKey("default", aliCloud.accessKey, aliCloud.accessSecret)
	if err != nil {
		return false, err
	}

	request := requests.NewCommonRequest()
	request.Method = "POST"
	request.Scheme = "https"
	request.Domain = "dysmsapi.aliyuncs.com"
	request.Version = "2017-05-25"
	request.ApiName = "SendSms"

	params := make(map[string]string)
	params["PhoneNumbers"] = aliCloud.phoneNumber
	params["SignName"] = aliCloud.signName
	params["TemplateCode"] = aliCloud.templateCode
	params["TemplateParam"] = aliCloud.templateParam

	request.QueryParams = params

	response, err := client.ProcessCommonRequest(request)
	if err != nil {
		log.Println("AliCloudSender->Send: " + err.Error())
		return false, errors.New("短信发送失败")
	}

	r := response.GetHttpContentString()
	log.Println("send success")
	log.Println(r)

	return true, nil
}
