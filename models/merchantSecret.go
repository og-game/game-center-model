package models

const TableNameMerchantSecret = "merchant_secret"

// MerchantSecret 商户密钥管理表
type MerchantSecret struct {
	SecretId     int64  `json:"secret_id" gorm:"secret_id"`
	MerchantId   int64  `json:"merchant_id" gorm:"merchant_id"`     // 商户id
	SecretKey    string `json:"secret_key" gorm:"secret_key"`       // 密钥
	ExpireAt     int64  `json:"expire_at" gorm:"expire_at"`         // 过期时间
	ValidityType int    `json:"validity_type" gorm:"validity_type"` // 有效期类型 1=1个月 2=3个月 3=6个月 4=12个月
	CreatedAt    int64  `json:"created_at" gorm:"created_at"`       // 创建时间
	UpdatedAt    int64  `json:"updated_at" gorm:"updated_at"`       // 更新时间
	DeletedAt    int64  `json:"deleted_at" gorm:"deleted_at"`       // 删除时间
}

// TableName 表名称
func (*MerchantSecret) TableName() string {
	return TableNameMerchantSecret
}
