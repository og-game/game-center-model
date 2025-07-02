package models

const TableNameUserBalance = "user_balance"

// UserBalance 用户游戏余额表
type UserBalance struct {
	BalanceId           int64   `json:"balance_id" gorm:"balance_id"`                     // 余额记录主键ID
	UserId              string  `json:"user_id" gorm:"user_id"`                           // 用户ID
	MerchantId          int64   `json:"merchant_id" gorm:"merchant_id"`                   // 商户ID
	GameId              int64   `json:"game_id" gorm:"game_id"`                           // 游戏ID
	PlatformId          int64   `json:"platform_id" gorm:"platform_id"`                   // 平台ID
	Balance             float64 `json:"balance" gorm:"balance"`                           // 游戏内余额(微)
	FrozenBalance       float64 `json:"frozen_balance" gorm:"frozen_balance"`             // 游戏内冻结余额
	TransferringBalance float64 `json:"transferring_balance" gorm:"transferring_balance"` // 转账中余额
	Version             int64   `json:"version" gorm:"version"`                           // 乐观锁版本号
	CreatedAt           int64   `json:"created_at" gorm:"created_at"`                     // 创建时间戳
	UpdatedAt           int64   `json:"updated_at" gorm:"updated_at"`                     // 更新时间戳
	DeletedAt           int64   `json:"deleted_at" gorm:"deleted_at"`                     // 删除时间戳（软删除）
}

// TableName 表名称
func (*UserBalance) TableName() string {
	return TableNameUserBalance
}
