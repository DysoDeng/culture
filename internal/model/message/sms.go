package message

import (
	"culture/cloud/base/internal/config"
	"culture/cloud/base/internal/support/db"
)

// 短信配置
type SmsConfig struct {
	Id              uint64      `gorm:"primary_key;autoIncrement" json:"id"`
	SmsType         string      `gorm:"type:varchar(150); not null;default:'';comment:短信服务商类型 ali_cloud-阿里云" json:"sms_type"`
	AppKey          string      `gorm:"type:varchar(150); not null;default:'';comment:短信AppKey" json:"app_key"`
	SecretKey       string      `gorm:"type:varchar(150); not null;default:'';comment:短信SecretKey" json:"secret_key"`
	FreeSignName    string      `gorm:"type:varchar(150); not null;default:'';comment:短信签名" json:"free_sign_name"`
	ValidCodeExpire uint        `gorm:"type:int(10); not null;default:0;comment:短信验证码过期时间，单位分钟" json:"valid_code_expire"`
	CreatedAt       db.JSONTime `gorm:"default null" json:"created_at"`
	UpdatedAt       db.JSONTime `gorm:"default null" json:"updated_at"`
}

func (SmsConfig) TableName() string {
	return config.Config.DataBase.Prefix + "sms_config"
}

// 短信模版
type SmsTemplate struct {
	Id           uint64      `gorm:"primary_key;autoIncrement" json:"id"`
	TemplateName string      `gorm:"type:varchar(150); not null; default:'';comment:模版名称" json:"template_name"`
	Template     string      `gorm:"type:varchar(150); not null; default:'';comment:短信模版类型" json:"template"`
	TemplateId   string      `gorm:"type:varchar(150); not null; default:'';comment:短信模版ID" json:"template_id"`
	CreatedAt    db.JSONTime `gorm:"default null" json:"created_at"`
	UpdatedAt    db.JSONTime `gorm:"default null" json:"updated_at"`
}

func (SmsTemplate) TableName() string {
	return config.Config.DataBase.Prefix + "sms_template"
}
