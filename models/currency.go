package models

import (
	"github.com/og-game/glib/stores/gormx"
)

const TableNameCurrency = "currency"

// Currency 货币表
type Currency struct {
	CurrencyId   int64  `json:"currency_id" gorm:"currency_id;primaryKey"` // 币种ID
	CurrencyCode string `json:"currency_code" gorm:"currency_code"`        // 币种
	Flag         string `json:"flag" gorm:"flag"`                          // 旗帜
	FlagIcon     string `json:"flag_icon" gorm:"flag_icon"`                // 旗帜图标
	Symbol       string `json:"symbol" gorm:"symbol"`                      // 币种符号
	Status       int    `json:"status" gorm:"status"`                      // 状态 1 启用 2 停用
	Weigh        int    `json:"weigh" gorm:"weigh"`                        // 排序
	gormx.Model
}

// TableName 表名称
func (*Currency) TableName() string {
	return TableNameCurrency
}
