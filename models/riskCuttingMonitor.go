package models

import (
	"github.com/og-game/glib/stores/gormx"
	"gorm.io/datatypes"
)

const (
	TableNameRiskCuttingMonitor = "risk_cutting_monitor"
)

// RiskCuttingMonitor 对赌初筛监控表
type RiskCuttingMonitor struct {
	CuttingID       uint64         `gorm:"primaryKey;autoIncrement;column:cutting_id" json:"cutting_id"`
	BatchID         string         `gorm:"column:batch_id;type:varchar(50);not null;uniqueIndex:uk_batch_evidence,priority:1" json:"batch_id"`
	EvidenceID      string         `gorm:"column:evidence_id;type:varchar(50);not null;uniqueIndex:uk_batch_evidence,priority:2" json:"evidence_id"`
	TriggerRuleCode string         `gorm:"column:trigger_rule_code;type:varchar(50);not null" json:"trigger_rule_code"`
	TriggerRuleName string         `gorm:"column:trigger_rule_name;type:varchar(50);not null" json:"trigger_rule_name"`
	UserID          int            `gorm:"column:user_id" json:"user_id"`
	UserName        string         `gorm:"column:user_name;type:varchar(100)" json:"user_name"`
	MerchantID      int            `gorm:"column:merchant_id" json:"merchant_id"`
	MerchantName    string         `gorm:"column:merchant_name;type:varchar(100)" json:"merchant_name"`
	CurrencyCode    string         `gorm:"column:currency_code;type:varchar(10)" json:"currency_code"`
	TriggerTime     uint           `gorm:"column:trigger_time;default:0" json:"trigger_time"`
	Metadata        datatypes.JSON `gorm:"column:metadata;type:json" json:"metadata"`
	Status          int8           `gorm:"column:status;default:2" json:"status"`                       // 1-已处理，2-待处理
	OperatorID      int64          `gorm:"column:operator_id" json:"operator_id"`                       // 操作人ID
	OperatorName    string         `gorm:"column:operator_name;type:varchar(100)" json:"operator_name"` // 操作人姓名
	OperatorIP      string         `gorm:"column:operator_ip;type:varchar(45)" json:"operator_ip"`      // 操作人IP
	OperatorTime    int            `gorm:"column:operator_time;not null" json:"operator_time"`          // 操作时间

	// 时间字段
	gormx.Model
}

// TableName 指定表名
func (*RiskCuttingMonitor) TableName() string {
	return TableNameRiskCuttingMonitor
}
