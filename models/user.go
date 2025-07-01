package models

const TableNameUser = "user"

// User 用户表
type User struct {
	UserId           int64  `json:"user_id" gorm:"user_id"`                     // 中台用户主键ID
	MerchantId       int64  `json:"merchant_id" gorm:"merchant_id"`             // 商户ID
	MerchantUserId   string `json:"merchant_user_id" gorm:"merchant_user_id"`   // 下游系统用户ID（唯一）
	MerchantUsername string `json:"merchant_username" gorm:"merchant_username"` // 下游系统用户名
	CurrencyCode     string `json:"currency_code" gorm:"currency_code"`         // 币种code
	AvatarUrl        string `json:"avatar_url" gorm:"avatar_url"`               // 头像URL
	Status           int    `json:"status" gorm:"status"`                       // 用户状态 1:正常 2:冻结 3:禁用 4:注销
	AccountType      int    `json:"account_type" gorm:"account_type"`           // 账号类型 1:正式账号 2:测试账号
	RiskLevel        int    `json:"risk_level" gorm:"risk_level"`               // 风险等级 1:正常 2:低风险 3:中风险 4:高风险
	CountryCode      string `json:"country_code" gorm:"country_code"`           // 国家代码
	Region           string `json:"region" gorm:"region"`                       // 地区
	LastLoginIp      string `json:"last_login_ip" gorm:"last_login_ip"`         // 最后登录IP
	LastDeviceId     string `json:"last_device_id" gorm:"last_device_id"`       // 最后使用设备ID
	CreatedAt        int64  `json:"created_at" gorm:"created_at"`               // 创建时间戳
	UpdatedAt        int64  `json:"updated_at" gorm:"updated_at"`               // 更新时间戳
	DeletedAt        int64  `json:"deleted_at" gorm:"deleted_at"`               // 删除时间戳（软删除）
}

// TableName 表名称
func (*User) TableName() string {
	return TableNameUser
}
