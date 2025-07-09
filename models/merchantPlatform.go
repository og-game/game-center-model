package models

import "github.com/shopspring/decimal"

const TableNameMerchantPlatform = "merchant_platform"

// MerchantPlatform undefined
type MerchantPlatform struct {
	MerchantPlatformId int64           `json:"merchant_platform_id" gorm:"merchant_platform_id,primaryKey"`
	MerchantId         int64           `json:"merchant_id" gorm:"merchant_id"` // 商户id
	PlatformId         int64           `json:"platform_id" gorm:"platform_id"` // 平台id
	Rate               decimal.Decimal `json:"rate" gorm:"rate"`               // 费率
	Status             int             `json:"status" gorm:"status"`           // 状态 1 启用 2禁用
	CreatedAt          int64           `json:"created_at" gorm:"created_at"`
	UpdatedAt          int64           `json:"updated_at" gorm:"updated_at"`
	DeletedAt          int64           `json:"deleted_at" gorm:"deleted_at"`
}

// TableName 表名称
func (*MerchantPlatform) TableName() string {
	return TableNameMerchantPlatform
}
