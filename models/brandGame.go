package models

import (
	"github.com/og-game/glib/stores/gormx"
)

const TableNameBrandGame = "brand_game"

// BrandGame 品牌游戏表
type BrandGame struct {
	BrandGameMaterialId int64  `json:"brand_game_material_id" gorm:"brand_game_material_id;primaryKey"` // 品牌游戏素材ID
	BrandId             int64  `json:"brand_id" gorm:"brand_id"`                                        // 品牌ID
	GameId              int64  `json:"game_id" gorm:"game_id"`                                          // 游戏ID
	Status              int    `json:"status" gorm:"status"`                                            // 状态 1 启用 2 禁用
	BrandGameName       string `json:"brand_game_name" gorm:"brand_game_name"`                          // 品牌游戏名称
	HorizontalIcon      string `json:"horizontal_icon" gorm:"horizontal_icon"`                          // 横版图标
	VerticalIcon        string `json:"vertical_icon" gorm:"vertical_icon"`                              // 竖版图标
	PublishPack         int    `json:"publish_pack" gorm:"publish_pack"`                                // 发布状态 1 未发布 2 已发布
	gormx.Model
}

// TableName 表名称
func (*BrandGame) TableName() string {
	return TableNameBrandGame
}
