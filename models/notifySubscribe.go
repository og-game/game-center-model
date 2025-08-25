package models

import (
	"github.com/og-game/glib/stores/gormx"
	"gorm.io/datatypes"
)

const (
	TableNameNotifySubscribe = "notify_subscribe"
)

// NotifySubscribe 通知订阅配置表
type NotifySubscribe struct {
	SubscriptionId   uint           `json:"subscription_id" gorm:"subscription_id;primaryKey;autoIncrement;comment:订阅ID"`
	SubscribeType    uint8          `json:"subscribe_type" gorm:"subscribe_type;type:tinyint unsigned;not null;default:1;comment:所有者类型 1=商户 2=系统内部"`
	MerchantId       uint64         `json:"merchant_id" gorm:"merchant_id;type:bigint unsigned;not null;default:0;index:idx_merchant_deleted,priority:1;comment:商户ID，系统内部时为0"`
	SubscriptionName string         `json:"subscription_name" gorm:"subscription_name;type:varchar(64);not null;default:'';comment:订阅名称"`
	Scope            uint8          `json:"scope" gorm:"scope;type:tinyint unsigned;not null;default:1;comment:订阅范围 1=自己的事件 2=全局事件 3=指定商户的事件"`
	TargetMerchants  datatypes.JSON `json:"target_merchants" gorm:"target_merchants;type:json;comment:目标商户列表(scope=3时使用)"`
	Mode             uint8          `json:"mode" gorm:"mode;type:tinyint unsigned;not null;default:3;comment:订阅模式 1=All 2=Category 3=Specific 4=Mixed"`
	Categories       datatypes.JSON `json:"categories" gorm:"categories;type:json;comment:订阅的大类列表"`
	EventTypes       datatypes.JSON `json:"event_types" gorm:"event_types;type:json;comment:订阅的具体事件类型列表"`
	FilterRules      datatypes.JSON `json:"filter_rules" gorm:"filter_rules;type:json;comment:过滤规则JSON"`
	Enabled          uint8          `json:"enabled" gorm:"enabled;type:tinyint unsigned;not null;default:1;index:idx_enabled,priority:1;comment:是否启用 1=是 2=否"`
	gormx.Model
}

// TableName 表名称
func (*NotifySubscribe) TableName() string {
	return TableNameNotifySubscribe
}
