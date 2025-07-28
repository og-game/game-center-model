package models

import "github.com/og-game/glib/stores/gormx"

const TableNameMerchant = "merchant"

// Merchant 商户表
type Merchant struct {
	MerchantId        int64  `json:"merchant_id" gorm:"merchant_id;primaryKey"`
	ParentId          int64  `json:"parent_id" gorm:"parent_id"`                       // 上级id
	CountryCode       string `json:"country_code" gorm:"country_code"`                 // 国家code
	CurrencyCode      string `json:"currency_code" gorm:"currency_code"`               // 币种code
	Name              string `json:"name" gorm:"name"`                                 // 商户名称
	CompanyName       string `json:"company_name" gorm:"company_name"`                 // 商户公司名称
	Email             string `json:"email" gorm:"email"`                               // 邮箱
	AppId             string `json:"app_id" gorm:"app_id"`                             // appid
	AppSecret         string `json:"app_secret" gorm:"app_secret"`                     // 当前的 app 密钥
	AppSecretExpireAt int64  `json:"app_secret_expire_at" gorm:"app_secret_expire_at"` // app密钥过期时间（Unix 时间戳）
	LastAppSecret     string `json:"last_app_secret" gorm:"last_app_secret"`           // 上一次 app 密钥
	Phone             string `json:"phone" gorm:"phone"`                               // 联系方式
	BusinessLicense   string `json:"business_license" gorm:"business_license"`         // 营业职照
	Status            int    `json:"status" gorm:"status"`                             // 状态 1启用 2禁用 3 审核中 4 审核失败
	gormx.Model
}

// TableName 表名称
func (*Merchant) TableName() string {
	return TableNameMerchant
}
