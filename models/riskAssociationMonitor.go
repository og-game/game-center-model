package models

import (
	"github.com/og-game/glib/stores/gormx"
	"gorm.io/datatypes"
)

const (
	TableNameRiskAssociationMonitor = "risk_association_monitor"
)

// RiskAssociationMonitor 刷子同关联监控表
type RiskAssociationMonitor struct {
	AssociationID   uint64         `gorm:"primaryKey;autoIncrement;column:association_id" json:"association_id"`
	BatchID         string         `gorm:"column:batch_id;type:varchar(50);not null;uniqueIndex:uk_batch_evidence,priority:1" json:"batch_id"`
	EvidenceID      string         `gorm:"column:evidence_id;type:varchar(50);not null;uniqueIndex:uk_batch_evidence,priority:2" json:"evidence_id"`
	UserID          int            `gorm:"column:user_id" json:"user_id"`
	UserName        string         `gorm:"column:user_name;type:varchar(100)" json:"user_name"`
	MerchantID      int            `gorm:"column:merchant_id;not null;index:idx_merchant_id" json:"merchant_id"`
	MerchantName    string         `gorm:"column:merchant_name;type:varchar(100)" json:"merchant_name"`
	CurrencyCode    string         `gorm:"column:currency_code;type:varchar(10)" json:"currency_code"`
	AssociationType string         `gorm:"column:association_type" json:"association_type"`
	AssociationInfo string         `gorm:"column:association_info;type:varchar(512)" json:"association_info"`
	Metadata        datatypes.JSON `gorm:"column:metadata;type:json" json:"metadata"`
	Status          int8           `gorm:"column:status;default:1" json:"status"` // 1-有效，2-无效
	// 时间字段
	gormx.Model
}

// TableName 指定表名
func (*RiskAssociationMonitor) TableName() string {
	return TableNameRiskAssociationMonitor
}
