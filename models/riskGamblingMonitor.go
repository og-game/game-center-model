package models

import (
	"github.com/og-game/glib/stores/gormx"
	"gorm.io/datatypes"
)

const (
	TableNameRiskGamblingMonitor = "risk_gambling_monitor"
)

// RiskGamblingMonitor 对赌监控表
type RiskGamblingMonitor struct {
	GamblingID       uint64         `gorm:"primaryKey;autoIncrement;column:gambling_id" json:"gambling_id"`
	BatchID          string         `gorm:"column:batch_id;type:varchar(50);not null;uniqueIndex:uk_batch_evidence,priority:1" json:"batch_id"`
	EvidenceID       string         `gorm:"column:evidence_id;type:varchar(50);not null;uniqueIndex:uk_batch_evidence,priority:2" json:"evidence_id"`
	RiskScore        int            `gorm:"column:risk_score;default:0" json:"risk_score"`
	TriggerRuleIDs   datatypes.JSON `gorm:"column:trigger_rule_ids;type:json" json:"trigger_rule_ids"`     // [1, 2]
	TriggerRuleNames datatypes.JSON `gorm:"column:trigger_rule_names;type:json" json:"trigger_rule_names"` // ["name1", "name2"]
	SameGameCount    int            `gorm:"column:same_game_count;default:0" json:"same_game_count"`
	ScoreCalcTime    int            `gorm:"column:score_calc_time" json:"score_calc_time"`

	// 玩家A信息
	UserACurrency     string `gorm:"column:user_a_currency;type:varchar(10)" json:"user_a_currency"`
	UserAMerchantID   int    `gorm:"column:user_a_merchant_id" json:"user_a_merchant_id"`
	UserAMerchantName string `gorm:"column:user_a_merchant_name;type:varchar(100)" json:"user_a_merchant_name"`
	UserAID           int    `gorm:"column:user_a_id" json:"user_a_id"`
	UserAName         string `gorm:"column:user_a_name;type:varchar(100)" json:"user_a_name"`

	// 玩家B信息
	UserBCurrency     string `gorm:"column:user_b_currency;type:varchar(10)" json:"user_b_currency"`
	UserBMerchantID   int    `gorm:"column:user_b_merchant_id" json:"user_b_merchant_id"`
	UserBMerchantName string `gorm:"column:user_b_merchant_name;type:varchar(100)" json:"user_b_merchant_name"`
	UserBID           int    `gorm:"column:user_b_id" json:"user_b_id"`
	UserBName         string `gorm:"column:user_b_name;type:varchar(100)" json:"user_b_name"`
	// 处理信息
	Status int8   `gorm:"column:status;default:2" json:"status"` // 1-已处理，2-待处理
	Remark string `gorm:"column:remark;type:text" json:"remark"`

	OperatorID   int64  `gorm:"column:operator_id" json:"operator_id"`                       // 操作人ID
	OperatorName string `gorm:"column:operator_name;type:varchar(100)" json:"operator_name"` // 操作人姓名
	OperatorIP   string `gorm:"column:operator_ip;type:varchar(45)" json:"operator_ip"`      // 操作人IP
	OperatorTime int    `gorm:"column:operator_time;not null" json:"operator_time"`          // 操作时间

	gormx.Model
}

// TableName 指定表名
func (*RiskGamblingMonitor) TableName() string {
	return TableNameRiskGamblingMonitor
}
