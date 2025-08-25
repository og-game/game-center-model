package models

import "gorm.io/datatypes"

const (
	TableNameNotifyEventRecord = "notify_event_record"
)

// NotifyEventRecord 事件记录表
type NotifyEventRecord struct {
	Id             uint64         `json:"id" gorm:"id;primaryKey;autoIncrement"`
	EventId        string         `json:"event_id" gorm:"event_id;type:varchar(64);not null;uniqueIndex:uniq_event_id;comment:事件唯一ID"`
	CategoryCode   string         `json:"category_code" gorm:"category_code;type:varchar(64);not null;comment:通知大类"`
	EventType      string         `json:"event_type" gorm:"event_type;type:varchar(64);not null;index:idx_event_type;comment:事件类型"`
	MerchantId     uint64         `json:"merchant_id" gorm:"merchant_id;type:bigint unsigned;index:idx_merchant_category,priority:1;comment:目标商户ID"`
	SourceService  string         `json:"source_service" gorm:"source_service;type:varchar(64);comment:来源服务"`
	Priority       uint8          `json:"priority" gorm:"priority;type:tinyint unsigned;not null;default:2;comment:优先级 1=Low 2=Normal 3=High 4=Critical"`
	IdempotencyKey string         `json:"idempotency_key" gorm:"idempotency_key;type:varchar(128);comment:幂等键"`
	TraceId        string         `json:"trace_id" gorm:"trace_id;type:varchar(64);comment:链路追踪ID"`
	EventData      datatypes.JSON `json:"event_data" gorm:"event_data;type:json;not null;comment:事件数据"`
	Metadata       datatypes.JSON `json:"metadata" gorm:"metadata;type:json;comment:元数据"`
	EventTime      uint64         `json:"event_time" gorm:"event_time;type:bigint unsigned;not null;index:idx_event_time;comment:事件发生时间(毫秒时间戳)"`
	CreatedAt      uint           `json:"created_at" gorm:"created_at;type:int unsigned;not null;default:0"`
}

// TableName 表名称
func (*NotifyEventRecord) TableName() string {
	return TableNameNotifyEventRecord
}
