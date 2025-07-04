package models

import (
	"github.com/og-game/glib/stores/gormx"
)

const TableNameUserBalance = "user_balance"

// UserBalance 用户游戏余额表
type UserBalance struct {
	BalanceId           int64   `json:"balance_id" gorm:"balance_id;primaryKey"`          // 余额记录主键ID
	UserId              int64   `json:"user_id" gorm:"user_id"`                           // 用户ID
	MerchantId          int64   `json:"merchant_id" gorm:"merchant_id"`                   // 商户ID
	PlatformId          int64   `json:"platform_id" gorm:"platform_id"`                   // 平台ID
	Balance             float64 `json:"balance" gorm:"balance"`                           // 游戏内余额[用户实时的账户余额]
	FrozenBalance       float64 `json:"frozen_balance" gorm:"frozen_balance"`             // 游戏内冻结余额[用户转入时的余额-不会变]
	TransferringBalance float64 `json:"transferring_balance" gorm:"transferring_balance"` // 转账中余额[进行提现转出或者转入的余额]
	Version             int     `json:"version"`
	gormx.Model                 // 删除时间戳（软删除）
}

// TableName 表名称
func (*UserBalance) TableName() string {
	return TableNameUserBalance
}
