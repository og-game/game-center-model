package models

import "github.com/og-game/glib/stores/gormx"

const (
	TableNameGameTagGame = "game_tag_game"
)

// GameTagGame 游戏标签-游戏关联表
type GameTagGame struct {
	TagGameID uint `gorm:"column:tag_game_id;primaryKey;autoIncrement" json:"tag_game_id"`
	TagID     uint `gorm:"column:tag_id;not null;index:idx_tag_id" json:"tag_id"`    // 标签id
	GameID    uint `gorm:"column:game_id;not null;index:idx_game_id" json:"game_id"` // 游戏id
	gormx.Model
}

// TableName 指定表名
func (*GameTagGame) TableName() string {
	return TableNameGameTagGame
}
