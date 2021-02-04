package message

import (
	"culture/internal/config"
	"culture/internal/support/db"
)

// 短信配置
type SmsConfig struct {
	Id              int64       `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	SmsType         string      `gorm:"type:varchar(150); not null" json:"sms_type"`
	AppKey          string      `gorm:"type:varchar(150); not null" json:"app_key"`
	SecretKey       string      `gorm:"type:varchar(150); not null" json:"secret_key"`
	FreeSignName    string      `gorm:"type:varchar(150); not null" json:"free_sign_name"`
	ValidCodeExpire int64       `gorm:"type:int(10); not null" json:"valid_code_expire"`
	CreatedAt       db.JSONTime `gorm:"default null" json:"created_at"`
	UpdatedAt       db.JSONTime `gorm:"default null" json:"updated_at"`
}

func (SmsConfig) TableName() string {
	return config.Config.DataBase.Prefix + "sms_config"
}

// 短信模版
type SmsTemplate struct {
	Id           int64       `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	TemplateName string      `gorm:"type:varchar(150); not null" json:"template_name"`
	Template     string      `gorm:"type:varchar(150); not null" json:"template"`
	TemplateId   string      `gorm:"type:varchar(150); not null" json:"template_id"`
	CreatedAt    db.JSONTime `gorm:"default null" json:"created_at"`
	UpdatedAt    db.JSONTime `gorm:"default null" json:"updated_at"`
}

func (SmsTemplate) TableName() string {
	return config.Config.DataBase.Prefix + "sms_template"
}
