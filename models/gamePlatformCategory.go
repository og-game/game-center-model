package models

import "github.com/og-game/glib/stores/gormx"

const TableNameGamePlatformCategory = "game_platform_category"

// GamePlatformCategory 厂商游戏类型code映射表
type GamePlatformCategory struct {
	GamePlatformCategoryId int64  `json:"game_platform_category_id" gorm:"game_platform_category_id"`
	PlatformId             int64  `json:"platform_id" gorm:"platform_id"`                       // 游戏厂商id
	CategoryCode           string `json:"category_code" gorm:"category_code"`                   // 类型code
	PlatformCategoryCode   string `json:"platform_category_code" gorm:"platform_category_code"` // 厂商类型code
	Status                 int    `json:"status" gorm:"status"`                                 // 状态 1 启用 2禁用
	gormx.Model
}

// TableName 表名称
func (*GamePlatformCategory) TableName() string {
	return TableNameGamePlatformCategory
}
