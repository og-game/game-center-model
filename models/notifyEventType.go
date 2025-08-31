package models

import "github.com/og-game/glib/stores/gormx"

const (
	TableNameNotifyEventType = "notify_event_type"
)

// NotifyEventType 事件类型配置表
type NotifyEventType struct {
	EventTypeId      uint64 `json:"event_type_id" gorm:"event_type_id;primaryKey;autoIncrement;comment:事件类型ID"`
	CategoryCode     string `json:"category_code" gorm:"category_code;type:varchar(64);not null;uniqueIndex:uniq_category_event,priority:1;index:idx_category_enabled,priority:1;comment:所属大类代码"`
	EventType        string `json:"event_type" gorm:"event_type;type:varchar(64);not null;uniqueIndex:uniq_category_event,priority:2;comment:事件类型代码（唯一标识）----保持和proto文件一致"`
	Name             string `json:"name" gorm:"name;type:varchar(64);not null;comment:事件名称"`
	Description      string `json:"description" gorm:"description;type:varchar(255);comment:事件描述"`
	Priority         uint8  `json:"priority" gorm:"priority;type:tinyint unsigned;not null;default:2;comment:优先级 1=低 2=普通 3=高 4=紧急"`
	AllowBatch       uint8  `json:"allow_batch" gorm:"allow_batch;type:tinyint unsigned;not null;default:1;comment:是否允许批量推送 1=是 2=否"`
	DefaultBatchSize uint   `json:"default_batch_size" gorm:"default_batch_size;type:int unsigned;not null;default:100;comment:默认批量大小"`
	Enabled          uint8  `json:"enabled" gorm:"enabled;type:tinyint unsigned;not null;default:1;index:idx_category_enabled,priority:2;comment:是否启用 1=是 2=否"`
	gormx.Model
}

// TableName 表名称
func (*NotifyEventType) TableName() string {
	return TableNameNotifyEventType
}
