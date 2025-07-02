package models

import "github.com/og-game/glib/stores/gormx"

const TableNameUserTransferRecord = "user_transfer_record"

// UserTransferRecord 玩家转账记录表
type UserTransferRecord struct {
	RecordId             int64   `json:"record_id" gorm:"record_id"`
	MerchantId           int64   `json:"merchant_id" gorm:"merchant_id"`                       // 商户id
	UserId               string  `json:"user_id" gorm:"user_id"`                               // 玩家id
	PlatformId           int64   `json:"platform_id" gorm:"platform_id"`                       // 游戏平台id
	GameId               int64   `json:"game_id" gorm:"game_id"`                               // 游戏id
	Amount               float64 `json:"amount" gorm:"amount"`                                 // 操作金额
	BalanceAfterTransfer float64 `json:"balance_after_transfer" gorm:"balance_after_transfer"` // 转账后余额
	TransferType         int     `json:"transfer_type" gorm:"transfer_type"`                   // 转账类型 1 转入 2转出
	SuccessAt            int64   `json:"success_at" gorm:"success_at"`                         // 转账成功时间
	gormx.Model
}

// TableName 表名称
func (*UserTransferRecord) TableName() string {
	return TableNameUserTransferRecord
}
