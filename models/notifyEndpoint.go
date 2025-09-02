package models

import (
	"github.com/og-game/glib/stores/gormx"
	"gorm.io/datatypes"
)

const (
	TableNameNotifyEndpoint = "notify_endpoint"
)

// NotifyEndpoint 通知端点配置表
type NotifyEndpoint struct {
	EndpointId      uint64         `json:"endpoint_id" gorm:"endpoint_id;primaryKey;autoIncrement;comment:端点ID"`
	EndpointType    uint8          `json:"endpoint_type" gorm:"endpoint_type;type:tinyint unsigned;not null;default:1;comment:所有者类型 1=商户 2=系统内部"`
	MerchantId      uint64         `json:"merchant_id" gorm:"merchant_id;type:bigint unsigned;not null;default:0;comment:商户ID，系统内部时为0"`
	EndpointName    string         `json:"endpoint_name" gorm:"endpoint_name;type:varchar(64);not null;comment:端点名称"`
	Channel         uint8          `json:"channel" gorm:"channel;type:tinyint unsigned;not null;comment:推送通道 1=HTTP 2=WebSocket 3=gRPC 4=Robot"`
	Url             string         `json:"url" gorm:"url;type:varchar(512);not null;comment:推送地址"`
	AuthType        uint8          `json:"auth_type" gorm:"auth_type;type:tinyint unsigned;not null;default:1;comment:认证类型 1=None 2=Token 3=Signature 4=OAuth"`
	AuthCredentials datatypes.JSON `json:"auth_credentials" gorm:"auth_credentials;type:text;comment:认证凭据(加密存储)"`
	RobotType       uint8          `json:"robot_type" gorm:"robot_type;type:tinyint unsigned;not null;default:0;comment:机器人类型 0=None 1=Slack 2=Lark 3=Teams 4=Telegram"`
	RobotConfig     datatypes.JSON `json:"robot_config" gorm:"robot_config;type:json;comment:机器人特定配置"`
	TimeoutSeconds  uint           `json:"timeout_seconds" gorm:"timeout_seconds;type:int unsigned;not null;default:5;comment:超时时间(秒)"`
	MaxRetry        uint8          `json:"max_retry" gorm:"max_retry;type:tinyint unsigned;not null;default:7;comment:最大重试次数"`
	Enabled         uint8          `json:"enabled" gorm:"enabled;type:tinyint unsigned;not null;default:1;comment:是否启用 1=是 2=否"`
	gormx.Model
	UniqueKey string `json:"unique_key" gorm:"unique_key;type:varchar(128);->;uniqueIndex:uniq_endpoint"` // Generated column, read-only
}

// TableName 表名称
func (*NotifyEndpoint) TableName() string {
	return TableNameNotifyEndpoint
}
