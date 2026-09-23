package models

import (
	"github.com/og-game/glib/stores/gormx"
	"gorm.io/datatypes"
)

const TableNameBrand = "brand"

// Brand 商户品牌表
type Brand struct {
	BrandId         int64                       `json:"brand_id" gorm:"brand_id;primaryKey"`      // 品牌ID
	ParentId        int64                       `json:"parent_id" gorm:"parent_id"`               // 上级品牌ID
	CountryCode     string                      `json:"country_code" gorm:"country_code"`         // 国家code
	CurrencyCode    datatypes.JSONSlice[string] `json:"currency_code" gorm:"currency_code"`       // 币种code
	Name            string                      `json:"name" gorm:"name"`                         // 商户名称
	CompanyName     string                      `json:"company_name" gorm:"company_name"`         // 商户公司名称
	Email           string                      `json:"email" gorm:"email"`                       // 邮箱
	AppId           string                      `json:"app_id" gorm:"app_id"`                     // appid
	Phone           string                      `json:"phone" gorm:"phone"`                       // 联系方式
	BusinessLicense string                      `json:"business_license" gorm:"business_license"` // 营业职照
	IPWhitelist     datatypes.JSONSlice[string] `json:"ip_whitelist" gorm:"ip_whitelist"`         // IP 白名单
	Status          int                         `json:"status" gorm:"status"`                     // 状态 1启用 2禁用 3 审核中 4 审核失败
	WalletType      int                         `json:"wallet_type" gorm:"wallet_type"`           // 钱包类型 1 单一钱包 2 转账钱包 默认转账钱包
	CallbackUrl     string                      `json:"callback_url" gorm:"callback_url"`         // 回调地址
	NumOfSubBrand   int64                       `json:"num_of_sub_brand" gorm:"num_of_sub_brand"` // 下级商户品牌数量
	gormx.Model
}

// TableName 表名称
func (*Brand) TableName() string {
	return TableNameBrand
}
