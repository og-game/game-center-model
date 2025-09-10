package models

import (
	"github.com/og-game/glib/stores/gormx"
)

const (
	TableNameRiskCuttingMonitor = "risk_cutting_monitor"
)

// RiskCuttingMonitor 对赌初筛监控表
type RiskCuttingMonitor struct {
	CuttingID       uint64 `gorm:"primaryKey;autoIncrement;column:cutting_id" json:"cutting_id"`
	BatchID         string `gorm:"column:batch_id;type:varchar(50);not null;uniqueIndex:uk_batch_evidence,priority:1" json:"batch_id"`
	EvidenceID      string `gorm:"column:evidence_id;type:varchar(50);not null;uniqueIndex:uk_batch_evidence,priority:2" json:"evidence_id"`
	TriggerRuleCode string `gorm:"column:trigger_rule_code;type:varchar(50);not null" json:"trigger_rule_code"`
	TriggerRuleName string `gorm:"column:trigger_rule_name;type:varchar(50);not null" json:"trigger_rule_name"`
	UserID          int    `gorm:"column:user_id" json:"user_id"`
	UserName        string `gorm:"column:user_name;type:varchar(100)" json:"user_name"`
	MerchantID      int    `gorm:"column:merchant_id" json:"merchant_id"`
	MerchantName    string `gorm:"column:merchant_name;type:varchar(100)" json:"merchant_name"`
	CurrencyCode    string `gorm:"column:currency_code;type:varchar(10)" json:"currency_code"`
	TriggerTime     uint   `gorm:"column:trigger_time;default:0" json:"trigger_time"`
	Status          int8   `gorm:"column:status;default:2" json:"status"` // 1-已处理，2-待处理
	ProcessTime     uint   `gorm:"column:process_time;default:0" json:"process_time"`
	ProcessUser     string `gorm:"column:process_user;type:varchar(50)" json:"process_user"`

	// 时间字段
	gormx.Model
}

// TableName 指定表名
func (*RiskCuttingMonitor) TableName() string {
	return TableNameRiskCuttingMonitor
}
