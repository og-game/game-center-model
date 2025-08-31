package models

import "gorm.io/datatypes"

const (
	TableNameNotifyPushTask = "notify_push_task"
)

// NotifyPushTask 推送任务表
type NotifyPushTask struct {
	TaskId         uint64         `json:"task_id" gorm:"task_id;primaryKey;autoIncrement"`
	NotificationId string         `json:"notification_id" gorm:"notification_id;type:varchar(255);not null;comment:推送唯一ID"`
	MerchantId     uint64         `json:"merchant_id" gorm:"merchant_id;type:bigint unsigned;not null;index:idx_merchant_status,priority:1;comment:商户ID"`
	SubscriptionId uint           `json:"subscription_id" gorm:"subscription_id;type:int unsigned;index:idx_subscription;comment:订阅ID"`
	EndpointId     uint64         `json:"endpoint_id" gorm:"endpoint_id;type:bigint unsigned;not null;index:idx_endpoint;comment:端点ID"`
	Channel        uint8          `json:"channel" gorm:"channel;type:tinyint unsigned;not null;comment:推送通道 1=HTTP 2=WebSocket 3=gRPC 4=Robot"`
	EventIds       datatypes.JSON `json:"event_ids" gorm:"event_ids;type:json;not null;comment:事件ID集合"`
	IsBatch        uint8          `json:"is_batch" gorm:"is_batch;type:tinyint unsigned;not null;default:2;comment:是否批量推送 1=是 2=否"`
	BatchSize      uint           `json:"batch_size" gorm:"batch_size;type:int unsigned;not null;default:0;comment:批量推送数量"`
	Payload        datatypes.JSON `json:"payload" gorm:"payload;type:json;not null;comment:推送载荷"`
	Status         uint8          `json:"status" gorm:"status;type:tinyint unsigned;not null;default:1;index:idx_merchant_status,priority:2;index:idx_retry,priority:1;comment:状态 1=待推送 2=成功 3=失败 4=重试中"`
	RetryCount     uint8          `json:"retry_count" gorm:"retry_count;type:tinyint unsigned;not null;default:0;comment:重试次数"`
	ErrorMessage   string         `json:"error_message" gorm:"error_message;type:text;comment:错误信息"`
	HttpStatus     uint           `json:"http_status" gorm:"http_status;type:int unsigned;comment:HTTP状态码"`
	Response       string         `json:"response" gorm:"response;type:text;comment:响应内容"`
	LatencyMs      uint           `json:"latency_ms" gorm:"latency_ms;type:int unsigned;comment:延迟(毫秒)"`
	NextRetryAt    uint           `json:"next_retry_at" gorm:"next_retry_at;type:int unsigned;index:idx_retry,priority:2;comment:下次重试时间"`
	CreatedAt      uint           `json:"created_at" gorm:"created_at;type:int unsigned;not null;default:0;index:idx_created"`
	ExecutedAt     uint           `json:"executed_at" gorm:"executed_at;type:int unsigned;comment:执行时间"`
	CompletedAt    uint           `json:"completed_at" gorm:"completed_at;type:int unsigned;comment:完成时间"`
}

// TableName 表名称
func (*NotifyPushTask) TableName() string {
	return TableNameNotifyPushTask
}
