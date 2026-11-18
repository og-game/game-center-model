package models

import "github.com/og-game/glib/stores/gormx"

const TableNameBrandMaterialGameRel = "brand_material_game_rel"

// BrandMaterialGameRel 品牌游戏素材游戏关系表
type BrandMaterialGameRel struct {
	BrandMaterialGameId int64 `json:"brand_material_game_id" gorm:"brand_material_game_id"` // 商户游戏素材关系id
	MaterialId          int64 `json:"material_id" gorm:"material_id"`                       // 素材模板ID
	BrandId             int64 `json:"brand_id" gorm:"brand_id"`                             // 品牌id
	GameId              int64 `json:"game_id" gorm:"game_id"`                               // 游戏id
	gormx.Model
}

// TableName 表名称
func (*BrandMaterialGameRel) TableName() string {
	return TableNameBrandMaterialGameRel
}
