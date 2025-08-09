package models

import (
	"github.com/og-game/glib/stores/gormx"
	"github.com/shopspring/decimal"
)

const TableNameGameBetSummary = "game_bet_summary"

// GameBetSummary 每轮投注汇总与通知状态记录
type GameBetSummary struct {
	SummaryID           int64           `json:"summary_id" grom:"summary_id,primaryKey"`            // 主键
	RoundID             string          `json:"round_id" grom:"round_id"`                           // 牌局ID
	UserID              int64           `json:"user_id" grom:"user_id"`                             // 用户ID
	MerchantID          int64           `json:"merchant_id" grom:"merchant_id"`                     // 商户ID
	MerchantUserID      string          `json:"merchant_user_id" gorm:"merchant_user_id"`           // 商户用户id
	PlatformID          int64           `json:"platform_id" gorm:"platform_id"`                     // 厂商id
	GameID              int64           `json:"game_id" grom:"game_id"`                             // 游戏ID
	OriginalPlatformID  int64           `json:"original_platform_id" gorm:"original_platform_id"`   // 原始厂商id
	OriginalGameID      int64           `json:"original_game_id" gorm:"original_game_id"`           // 原始游戏id
	CategoryCode        string          `json:"category_code" gorm:"category_code"`                 // 分类code
	CurrencyCode        string          `json:"currency_code" grom:"currency_code"`                 // 币种code
	BetCount            int64           `json:"bet_count" grom:"bet_count"`                         // 有效投注笔数（不含取消）
	BetAmount           decimal.Decimal `json:"bet_amount" grom:"bet_amount"`                       // 有效投注金额
	CancelBetCount      int64           `json:"cancel_bet_count" grom:"cancel_bet_count"`           // 取消投注笔数
	CancelBetAmount     decimal.Decimal `json:"cancel_bet_amount" grom:"cancel_bet_amount"`         // 取消投注金额
	SettleCount         int64           `json:"settle_count" grom:"settle_count"`                   // 有效结算笔数
	SettleAmount        decimal.Decimal `json:"settle_amount" grom:"settle_amount"`                 // 有效结算金额
	CancelSettleCount   int64           `json:"cancel_settle_count" grom:"cancel_settle_count"`     // 取消结算笔数
	CancelSettleAmount  decimal.Decimal `json:"cancel_settle_amount" grom:"cancel_settle_amount"`   // 取消结算金额
	AdjustmentCount     int64           `json:"adjustment_count" grom:"adjustment_count"`           // 调整次数
	AdjustmentAmount    decimal.Decimal `json:"adjustment_amount" gorm:"adjustment_amount"`         // 调整金额
	DepositCount        int64           `json:"deposit_count" gorm:"deposit_count"`                 // 押金次数
	DepositAmount       decimal.Decimal `json:"deposit_amount" gorm:"deposit_amount"`               // 押金金额
	ReturnDepositCount  int64           `json:"return_deposit_count" gorm:"return_deposit_count"`   // 返还押金次数
	ReturnDepositAmount decimal.Decimal `json:"return_deposit_amount" gorm:"return_deposit_amount"` // 返还押金金额
	BetAt               int64           `json:"bet_at" grom:"bet_at"`                               // 下注时间
	SettledAt           int64           `json:"settled_at" grom:"settled_at"`                       // 结算时间
	gormx.Model
}

func (*GameBetSummary) TableName() string {
	return TableNameGameBetSummary
}
