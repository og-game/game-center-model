package models

import "github.com/og-game/glib/stores/gormx"

const TableNameUserGameLog = "user_game_log"

// UserGameLog undefined
type UserGameLog struct {
	GameLogId  int64  `json:"game_log_id" gorm:"game_log_id,primaryKey"`
	UserId     int64  `json:"user_id" gorm:"user_id"`         // 用户id
	MerchantId int64  `json:"merchant_id" gorm:"merchant_id"` // 商户id
	PlatformId int64  `json:"platform_id" gorm:"platform_id"` // 游戏平台id
	GameId     int64  `json:"game_id" gorm:"game_id"`         // 游戏id
	DeviceId   string `json:"device_id" gorm:"device_id"`     // 设备id
	DeviceOs   string `json:"device_os" gorm:"device_os"`     // 设备os
	DeviceIp   string `json:"device_ip" gorm:"device_ip"`     // 设备ip
	gormx.Model
}

// TableName 表名称
func (*UserGameLog) TableName() string {
	return TableNameUserGameLog
}
