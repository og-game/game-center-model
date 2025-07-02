package models

import "github.com/og-game/glib/stores/gormx"

const TableNameGamePlatform = "game_platform"

// GamePlatform 游戏平台表
type GamePlatform struct {
	PlatformId    int64  `json:"platform_id" gorm:"platform_id"`       // 平台id
	PlatformName  string `json:"platform_name" gorm:"platform_name"`   // 平台名称
	PlatformCode  string `json:"platform_code" gorm:"platform_code"`   // 平台code
	RequestDomain string `json:"request_domain" gorm:"request_domain"` // 请求接口域名
	WalletType    int    `json:"wallet_type" gorm:"wallet_type"`       // 钱包类型 1 单一  2转账
	Weigh         int    `json:"weigh" gorm:"weigh"`                   // 排序值
	Status        int    `json:"status" gorm:"status"`                 // 状态 1 启用 2禁用
	gormx.Model
}

// TableName 表名称
func (*GamePlatform) TableName() string {
	return TableNameGamePlatform
}
