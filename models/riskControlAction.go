package models

import (
	"github.com/og-game/glib/stores/gormx"
	"gorm.io/datatypes"
)

const (
	TableNameRiskControlAction = "risk_control_action"
)

// RiskControlAction 风控处理记录表
type RiskControlAction struct {
	ActionID          uint64         `gorm:"column:action_id;primaryKey;autoIncrement" json:"action_id"`
	GamblingID        int            `gorm:"column:gambling_id;not null" json:"gambling_id"`                       // 处理的监控ID
	MerchantID        uint64         `gorm:"column:merchant_id;not null" json:"merchant_id"`                       // 商户ID
	MerchantName      string         `gorm:"column:merchant_name;type:varchar(100);not null" json:"merchant_name"` // 商户名称
	UserID            int64          `gorm:"column:user_id;not null" json:"user_id"`                               // 玩家的ID
	UserName          string         `gorm:"column:user_name;type:varchar(100)" json:"user_name"`                  // 玩家的名称
	ActionType        int8           `gorm:"column:action_type;not null" json:"action_type"`                       // 1-禁止所有游戏, 2-禁止部分游戏
	RestrictionReason string         `gorm:"column:restriction_reason;type:text" json:"restriction_reason"`        // 限制原因（可存储多个，用逗号分隔）
	BannedGameTags    datatypes.JSON `gorm:"column:banned_game_tags;type:json" json:"banned_game_tags"`            // 被禁止的游戏标签列表
	StartTime         int            `gorm:"column:start_time;default:0" json:"start_time"`                        // 处理开始时间
	EndTime           uint           `gorm:"column:end_time;default:0" json:"end_time"`                            // 处理结束时间（如果有期限）
	OperatorID        int64          `gorm:"column:operator_id" json:"operator_id"`                                // 操作人ID
	OperatorName      string         `gorm:"column:operator_name;type:varchar(100)" json:"operator_name"`          // 操作人姓名
	OperatorIP        string         `gorm:"column:operator_ip;type:varchar(45)" json:"operator_ip"`               // 操作人IP
	OperatorTime      int            `gorm:"column:operator_time;not null" json:"operator_time"`                   // 操作时间
	Status            int8           `gorm:"column:status;not null;default:1" json:"status"`                       // 1-生效中, 2-已过期, 3-已取消
	Remark            string         `gorm:"column:remark;type:text" json:"remark"`                                // 备注信息
	// 时间字段
	gormx.Model
}

// TableName 指定表名
func (*RiskControlAction) TableName() string {
	return TableNameRiskControlAction
}
