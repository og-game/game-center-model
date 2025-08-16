package models

import (
	"github.com/og-game/glib/stores/gormx"
	"gorm.io/datatypes"
)

const TableNameMQFailedMessage = "mq_failed_message"

type MQFailedMessage struct {
	MQFailedID    int64          `gorm:"column:mq_failed_id;primaryKey;autoIncrement;comment:主键ID"`
	Topic         string         `gorm:"column:topic;type:varchar(128);not null;comment:MQ主题"`
	Tag           string         `gorm:"column:tag;type:varchar(128);comment:消息标签"`
	MessageGroup  string         `gorm:"column:message_group;type:varchar(128);comment:顺序消息分组key"`
	MessageKeys   datatypes.JSON `gorm:"column:message_keys;type:json;comment:消息keys，JSON数组格式"`
	MessageBody   string         `gorm:"column:message_body;type:longtext;not null;comment:消息体内容"`
	MessageType   string         `gorm:"column:message_type;type:varchar(50);not null;comment:消息类型：balance_change|bet_detail|等等"`
	Properties    datatypes.JSON `gorm:"column:properties;type:json;comment:消息属性，JSON格式"`
	RetryCount    int            `gorm:"column:retry_count;not null;default:0;comment:重试次数"`
	MaxRetryCount int            `gorm:"column:max_retry_count;not null;default:3;comment:最大重试次数"`
	Status        int8           `gorm:"column:status;not null;default:1;comment:状态：1-待重试，2-成功，3-失败，4-手动取消"`
	ErrorMsg      string         `gorm:"column:error_msg;type:text;comment:错误信息"`
	SuccessTime   int64          `gorm:"column:success_time;comment:成功发送时间戳"`
	gormx.Model
}

// TableName overrides the default table name
func (*MQFailedMessage) TableName() string {
	return TableNameMQFailedMessage
}
