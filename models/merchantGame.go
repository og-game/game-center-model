package models

import (
	"github.com/og-game/glib/stores/gormx"
	"github.com/shopspring/decimal"
	"gorm.io/datatypes"
)

const TableNameMerchantGame = "merchant_game"

// MerchantGame undefined
type MerchantGame struct {
	MerchantGameId int64                       `json:"merchant_game_id" gorm:"merchant_game_id"`
	MerchantId     int64                       `json:"merchant_id" gorm:"merchant_id"` // 商户id
	GameId         int64                       `json:"game_id" gorm:"game_id"`         // 游戏id
	Images         datatypes.JSONSlice[string] `json:"images" gorm:"images"`           // 图片列表
	Rate           decimal.Decimal             `json:"rate" gorm:"rate"`               // 费率
	Status         int                         `json:"status" gorm:"status"`           // 状态 1 启用 2禁用
	gormx.Model
}

// TableName 表名称
func (*MerchantGame) TableName() string {
	return TableNameMerchantGame
}
