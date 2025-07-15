package models

const TableNameUserGameLog = "user_game_log"

// UserGameLog undefined
type UserGameLog struct {
	GameLogId  int64  `json:"game_log_id" gorm:"game_log_id"`
	UserId     int64  `json:"user_id" gorm:"user_id"`         // 用户id
	MerchantId int64  `json:"merchant_id" gorm:"merchant_id"` // 商户id
	PlatformId int64  `json:"platform_id" gorm:"platform_id"` // 游戏平台id
	GameId     int64  `json:"game_id" gorm:"game_id"`         // 游戏id
	DeviceId   string `json:"device_id" gorm:"device_id"`     // 设备id
	DeviceOs   string `json:"device_os" gorm:"device_os"`     // 设备os
	DeviceIp   string `json:"device_ip" gorm:"device_ip"`     // 设备ip
	CreatedAt  int64  `json:"created_at" gorm:"created_at"`
	UpdatedAt  int64  `json:"updated_at" gorm:"updated_at"`
	DeletedAt  int64  `json:"deleted_at" gorm:"deleted_at"`
}

// TableName 表名称
func (*UserGameLog) TableName() string {
	return TableNameUserGameLog
}
