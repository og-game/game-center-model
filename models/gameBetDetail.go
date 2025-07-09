package models

import (
	"github.com/og-game/glib/stores/gormx"
	"github.com/shopspring/decimal"
)

const TableNameGameBetDetail = "game_bet_detail"

// GameBetDetail 投注/结算明细记录表
type GameBetDetail struct {
	DetailID        int64           `json:"detail_id" gorm:"detail_id;primaryKey"`      // 主键
	RecordType      int             `json:"record_type" gorm:"record_type"`             // '记录类型：1-投注，2-结算，3-投注取消，4-结算取消'
	UserID          int64           `json:"user_id" gorm:"user_id"`                     // '用户ID'
	MerchantID      int64           `json:"merchant_id" gorm:"merchant_id"`             // '商户ID'
	GameID          int64           `json:"game_id" gorm:"game_id"`                     // '游戏ID'
	RoundID         string          `json:"round_id" gorm:"round_id"`                   // '牌局ID'
	OrderNo         string          `json:"order_no" gorm:"order_no"`                   // '注单唯一ID'
	PlatformOrderNo string          `json:"platform_order_no" gorm:"platform_order_no"` // '游戏平台原始单号'
	CurrencyCode    string          `json:"currency_code" gorm:"currency_code"`         // '币种code'
	BetAmount       decimal.Decimal `json:"bet_amount" gorm:"bet_amount"`               // '投注金额（仅 record_type 为 1 或 3 使用）'
	SettleAmount    decimal.Decimal `json:"settle_amount" gorm:"settle_amount"`         // '结算金额（仅 record_type 为 2 或 4 使用）'
	SourceTime      int64           `json:"source_time" gorm:"source_time"`             // '投注或结算来源时间戳（毫秒）'
	Status          int64           `json:"status" gorm:"status"`                       // '状态：1-有效，2-取消，3-冲正等业务状态'
	gormx.Model
}

// TableName 表名称
func (*GameBetDetail) TableName() string {
	return TableNameGameBetDetail
}
