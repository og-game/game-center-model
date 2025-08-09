package models

import (
	"github.com/og-game/glib/stores/gormx"
	"github.com/shopspring/decimal"
)

const TableNameGameBetDetail = "game_bet_detail"

// GameBetDetail 投注/结算明细记录表
type GameBetDetail struct {
	DetailID           int64           `json:"detail_id" gorm:"detail_id;primaryKey"`            // 主键
	RecordType         int             `json:"record_type" gorm:"record_type"`                   // '记录类型：1-投注，2-结算，3-押金，4-返还押金 5-取消投注 6-结算撤单 7-重新派奖 8-调整金额'
	UserID             int64           `json:"user_id" gorm:"user_id"`                           // '用户ID'
	MerchantID         int64           `json:"merchant_id" gorm:"merchant_id"`                   // '商户ID'
	MerchantUserID     string          `json:"merchant_user_id" gorm:"merchant_user_id"`         // 商户用户id
	PlatformID         int64           `json:"platform_id" gorm:"platform_id"`                   // 厂商id
	GameID             int64           `json:"game_id" gorm:"game_id"`                           // '游戏ID'
	OriginalPlatformID int64           `json:"original_platform_id" gorm:"original_platform_id"` // 原始厂商id
	OriginalGameID     int64           `json:"original_game_id" gorm:"original_game_id"`         // 原始游戏id
	CategoryCode       string          `json:"category_code" gorm:"category_code"`               // 分类code
	RoundID            string          `json:"round_id" gorm:"round_id"`                         // '牌局ID'
	OrderID            string          `json:"order_id" gorm:"order_id"`                         // '注单唯一ID'
	PlatformOrderID    string          `json:"platform_order_id" gorm:"platform_order_id"`       // '游戏平台原始单号'
	CurrencyCode       string          `json:"currency_code" gorm:"currency_code"`               // '币种code'
	BetAmount          decimal.Decimal `json:"bet_amount" gorm:"bet_amount"`                     // '投注金额'
	SettleAmount       decimal.Decimal `json:"settle_amount" gorm:"settle_amount"`               // '结算金额'
	AdjustmentAmount   decimal.Decimal `json:"adjustment_amount" gorm:"adjustment_amount"`       // 调整金额
	DepositAmount      decimal.Decimal `json:"deposit_amount" gorm:"deposit_amount"`             // 押金金额
	SourceTime         int64           `json:"source_time" gorm:"source_time"`                   // '投注或结算来源时间戳（毫秒）'
	Status             int64           `json:"status" gorm:"status"`                             // '状态：1-有效，2-取消，3-冲正等业务状态'
	ClientIP           string          `json:"client_ip" gorm:"client_ip"`                       // '客户端IP'
	DeviceID           string          `json:"device_id" gorm:"device_id"`                       // '设备ID'
	DeviceOS           string          `json:"device_os" gorm:"device_os"`                       // '设备操作系统'
	gormx.Model
}

// TableName 表名称
func (*GameBetDetail) TableName() string {
	return TableNameGameBetDetail
}
