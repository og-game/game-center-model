package models

import (
	"github.com/og-game/glib/stores/gormx"
	"github.com/shopspring/decimal"
	"gorm.io/datatypes"
)

const TableNameGame = "game"

// Game 游戏表
type Game struct {
	GameId              int64                       `json:"game_id" gorm:"game_id;primaryKey"`                  // 游戏id
	ThreeId             string                      `json:"three_id" gorm:"three_id"`                           // 三方游戏id
	PlatformId          int64                       `json:"platform_id" gorm:"platform_id"`                     // 平台id
	Countries           datatypes.JSONSlice[string] `json:"countries" gorm:"countries"`                         // 国家列表
	Languages           datatypes.JSONSlice[string] `json:"languages" gorm:"languages"`                         // 语言列表
	Currencies          datatypes.JSONSlice[string] `json:"currencies" gorm:"currencies"`                       // 货币列表
	Rate                decimal.Decimal             `json:"rate" gorm:"rate"`                                   // 厂商费率
	CategoryCode        string                      `json:"category_code" gorm:"category_code"`                 // 分类code
	GameName            string                      `json:"game_name" gorm:"game_name"`                         // 游戏名称
	Icon                string                      `json:"icon" gorm:"icon"`                                   // 图标
	ThumbCover          string                      `json:"thumb_cover" gorm:"thumb_cover"`                     // 缩略图
	HorizontalCover     string                      `json:"horizontal_cover" gorm:"horizontal_cover"`           // 垂直封面
	VerticalCover       string                      `json:"vertical_cover" gorm:"vertical_cover"`               // 水平封面
	ScreenType          int                         `json:"screen_type" gorm:"screen_type"`                     // 横竖屏 1都支持 2 竖屏 3横屏
	Weigh               int                         `json:"weigh" gorm:"weigh"`                                 // 排序值
	OpenMethod          int                         `json:"open_method" gorm:"open_method"`                     // 打开方式 1 iframe 2 新标签页
	Status              int                         `json:"status" gorm:"status"`                               // 状态 1启用 2禁用
	PlatformGameStatus  int                         `json:"platform_game_status" gorm:"platform_game_status"`   // 状态 1启用 2禁用
	IsTrialPlay         int                         `json:"is_trial_play" gorm:"is_trial_play"`                 // 是否允许试玩 1 是 2否
	IsDelayWithdraw     int                         `json:"is_delay_withdraw" gorm:"is_delay_withdraw"`         // 是否需要延迟提款
	IsMaterial          int                         `json:"is_material" gorm:"is_material"`                     // 是否支持素材 1 支持 2 不支持
	DelayWithdrawSecond int                         `json:"delay_withdraw_second" gorm:"delay_withdraw_second"` // 延迟提款秒数
	RebateRate          decimal.Decimal             `json:"rebate_rate" gorm:"rebate_rate"`                     // 返奖率
	Remark              string                      `json:"remark" gorm:"remark"`                               // 备注
	gormx.Model
}

// TableName 表名称
func (*Game) TableName() string {
	return TableNameGame
}
