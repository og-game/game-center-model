package models

import "github.com/og-game/glib/stores/gormx"

const (
	TableNameNotifySubscribeEndpoint = "notify_subscribe_endpoint"
)

// NotifySubscribeEndpoint 订阅与端点关联表
type NotifySubscribeEndpoint struct {
	Id             uint64 `json:"id" gorm:"id;primaryKey;autoIncrement;comment:自增ID"`
	SubscriptionId uint   `json:"subscription_id" gorm:"subscription_id;type:int unsigned;not null;uniqueIndex:uniq_subscription_endpoint,priority:1;index:idx_subscription;comment:订阅ID"`
	EndpointId     uint64 `json:"endpoint_id" gorm:"endpoint_id;type:bigint unsigned;not null;uniqueIndex:uniq_subscription_endpoint,priority:2;index:idx_endpoint;comment:端点ID"`
	gormx.Model
}

// TableName 表名称
func (*NotifySubscribeEndpoint) TableName() string {
	return TableNameNotifySubscribeEndpoint
}
