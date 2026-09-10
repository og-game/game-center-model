package models

import "github.com/og-game/glib/stores/gormx"

const TableNameMerchantMaterialGameRel = "merchant_material_game_rel"

// MerchantMaterialGameRel 商户游戏素材游戏关系表
type MerchantMaterialGameRel struct {
	MerchantMaterialGameId int64 `json:"merchant_material_game_id" gorm:"merchant_material_game_id"` // 商户游戏素材关系id
	MaterialId             int64 `json:"material_id" gorm:"material_id"`                             // 素材模板ID
	MerchantId             int64 `json:"merchant_id" gorm:"merchant_id"`                             // 商户id
	GameId                 int64 `json:"game_id" gorm:"game_id"`                                     // 游戏id
	gormx.Model
}

// TableName 表名称
func (*MerchantMaterialGameRel) TableName() string {
	return TableNameMerchantMaterialGameRel
}
