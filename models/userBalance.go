package models

import (
	"github.com/og-game/glib/stores/gormx"
	"gorm.io/plugin/optimisticlock"
)

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
	Version             optimisticlock.Version
	gormx.Model         // 删除时间戳（软删除）
}

// TableName 表名称
func (*UserBalance) TableName() string {
	return TableNameUserBalance
}
