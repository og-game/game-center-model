package models

import (
	"github.com/og-game/glib/stores/gormx"
	"github.com/shopspring/decimal"
)

const TableNameGameBetSummary = "game_bet_summary"

// GameBetSummary 每轮投注汇总与通知状态记录
type GameBetSummary struct {
	SummaryID          int64           `json:"summary_id" grom:"summary_id,primaryKey"`          // 主键
	RoundID            string          `json:"round_id" grom:"round_id"`                         // 牌局ID
	UserID             int64           `json:"user_id" grom:"user_id"`                           // 用户ID
	MerchantID         int64           `json:"merchant_id" grom:"merchant_id"`                   // 商户ID
	GameID             int64           `json:"game_id" grom:"game_id"`                           // 游戏ID
	CurrencyCode       string          `json:"currency_code" grom:"currency_code"`               // 币种code
	TotalBetCount      int64           `json:"total_bet_count" grom:"total_bet_count"`           // 有效投注笔数（不含取消）
	TotalBetAmount     decimal.Decimal `json:"total_bet_amount" grom:"total_bet_amount"`         // 有效投注金额
	CancelBetCount     int64           `json:"cancel_bet_count" grom:"cancel_bet_count"`         // 取消投注笔数
	CancelBetAmount    decimal.Decimal `json:"cancel_bet_amount" grom:"cancel_bet_amount"`       // 取消投注金额
	TotalSettleCount   int64           `json:"total_settle_count" grom:"total_settle_count"`     // 有效结算笔数
	TotalSettleAmount  decimal.Decimal `json:"total_settle_amount" grom:"total_settle_amount"`   // 有效结算金额
	CancelSettleCount  int64           `json:"cancel_settle_count" grom:"cancel_settle_count"`   // 取消结算笔数
	CancelSettleAmount decimal.Decimal `json:"cancel_settle_amount" grom:"cancel_settle_amount"` // 取消结算金额
	SettleCompleted    int64           `json:"settle_completed" grom:"settle_completed"`         // 是否全部结算完成：0-否，1-是
	Notified           int64           `json:"notified" grom:"notified"`                         // 是否已通知下游：0-未通知，1-成功，2-失败
	NotifyAttempts     int64           `json:"notify_attempts" grom:"notify_attempts"`           // 通知尝试次数
	NotifyType         int64           `json:"notify_type" grom:"notify_type"`                   // 通知类型：1-正常结算，2-结算取消，3-投注取消等
	LastNotifyTime     int64           `json:"last_notify_time" grom:"last_notify_time"`         // 最后通知时间
	gormx.Model
}

func (*GameBetSummary) TableName() string {
	return TableNameGameBetSummary
}
