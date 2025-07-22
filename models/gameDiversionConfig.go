package models

import "gorm.io/datatypes"

const TableNameGameDiversionConfig = "game_diversion_config"

// GameDiversionConfig 游戏分流配置表
type GameDiversionConfig struct {
	ConfigId   int64                                      `json:"config_id" gorm:"config_id"`
	PlatformId int32                                      `json:"platform_id" gorm:"platform_id"` // 原游戏厂商id
	GameId     int32                                      `json:"game_id" gorm:"game_id"`         // 原游戏id
	TargetType int8                                       `json:"target_type" gorm:"target_type"` // 类型 1 根据商户 2 根据币种
	TargetId   string                                     `json:"target_id" gorm:"target_id"`     // 目标id 1 是商户ID 2 是币种 code
	Strategy   datatypes.JSONSlice[GameDiversionStrategy] `json:"strategy" gorm:"strategy"`       // 策略
	Status     int8                                       `json:"status" gorm:"status"`           // 状态 1 启用 2禁用
	CreatedAt  int64                                      `json:"created_at" gorm:"created_at"`
	UpdatedAt  int64                                      `json:"updated_at" gorm:"updated_at"`
	DeletedAt  int64                                      `json:"deleted_at" gorm:"deleted_at"`
}

type GameDiversionStrategy struct {
	GameID int64 `json:"game_id" gorm:"game_id"` // 游戏id
	Weight int64 `json:"weight" gorm:"weight"`   // 权重
}

// TableName 表名称
func (*GameDiversionConfig) TableName() string {
	return TableNameGameDiversionConfig
}
