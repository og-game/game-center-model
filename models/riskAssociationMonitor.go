package models

import "github.com/og-game/glib/stores/gormx"

const (
	TableNameRiskAssociationMonitor = "risk_association_monitor"
)

// RiskAssociationMonitor 刷子同关联监控表
type RiskAssociationMonitor struct {
	AssociationID     uint64 `gorm:"primaryKey;autoIncrement;column:association_id" json:"association_id"`
	RiskAssociationID string `gorm:"column:risk_association_id;type:varchar(255);not null" json:"risk_association_id"`
	MerchantID        int    `gorm:"column:merchant_id;not null;index:idx_merchant_id" json:"merchant_id"`
	MerchantName      string `gorm:"column:merchant_name;type:varchar(100)" json:"merchant_name"`
	CurrencyCode      string `gorm:"column:currency_code;type:varchar(10)" json:"currency_code"`
	AssociationReason string `gorm:"column:association_reason;type:varchar(255)" json:"association_reason"`
	AssociationInfo   string `gorm:"column:association_info;type:text" json:"association_info"`
	AssociationCount  int    `gorm:"column:association_count;default:0;index:idx_association_count" json:"association_count"`
	Status            int8   `gorm:"column:status;default:1" json:"status"` // 1-有效，2-无效

	// 时间字段
	gormx.Model
}

// TableName 指定表名
func (*RiskAssociationMonitor) TableName() string {
	return TableNameRiskAssociationMonitor
}
