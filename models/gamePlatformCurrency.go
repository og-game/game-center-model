package models

import (
	"github.com/og-game/glib/stores/gormx"
	"github.com/shopspring/decimal"
)

const TableNameGamePlatformCurrency = "game_platform_currency"

// GamePlatformCurrency 厂商币种比例表
type GamePlatformCurrency struct {
	GamePlatformCurrencyId int64           `json:"game_platform_currency_id" gorm:"game_platform_currency_id;primaryKey;autoIncrement"`
	PlatformId             int64           `json:"platform_id" gorm:"platform_id"`     // 游戏厂商id
	CurrencyCode           string          `json:"currency_code" gorm:"currency_code"` // 币种code
	Rate                   decimal.Decimal `json:"rate" gorm:"rate"`                   // 比例
	Status                 int             `json:"status" gorm:"status"`               // 状态 1 启用 2禁用
	gormx.Model
}

// TableName 表名称
func (*GamePlatformCurrency) TableName() string {
	return TableNameGamePlatformCurrency
}
