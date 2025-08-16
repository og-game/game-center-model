package models

import "github.com/og-game/glib/stores/gormx"

const TableNameGamePlatformLanguage = "game_platform_language"

// GamePlatformLanguage 游戏厂商语言code映射表
type GamePlatformLanguage struct {
	GamePlatformLanguageId int64  `json:"game_platform_language_id" gorm:"game_platform_language_id"`
	PlatformId             int64  `json:"platform_id" gorm:"platform_id"`                       // 游戏厂商id
	LanguageCode           string `json:"language_code" gorm:"language_code"`                   // 语言code
	PlatformLanguageCode   string `json:"platform_language_code" gorm:"platform_language_code"` // 厂商语言code
	Status                 int    `json:"status" gorm:"status"`                                 // 状态 1 启用 2禁用
	gormx.Model
}

// TableName 表名称
func (*GamePlatformLanguage) TableName() string {
	return TableNameGamePlatformLanguage
}
