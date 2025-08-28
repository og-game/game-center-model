package models

import "github.com/og-game/glib/stores/gormx"

const (
	TableNameNotifyCategory = "notify_category"
)

// NotifyCategory 大类配置表
type NotifyCategory struct {
	CategoryId       uint64 `json:"category_id" gorm:"category_id;primaryKey;autoIncrement"`
	CategoryCode     string `json:"category_code" gorm:"category_code;uniqueIndex:uniq_category_code;type:varchar(64);not null;comment:大类代码（唯一标识）---保持和proto文件一致"`
	Name             string `json:"name" gorm:"name;type:varchar(64);not null;comment:大类名称"`
	Description      string `json:"description" gorm:"description;type:varchar(255);comment:大类描述"`
	AllowBatch       uint8  `json:"allow_batch" gorm:"allow_batch;type:tinyint unsigned;not null;default:1;comment:是否允许批量推送 1=是 2=否"`
	DefaultBatchSize uint   `json:"default_batch_size" gorm:"default_batch_size;type:int unsigned;not null;default:10;comment:默认批量大小"`
	gormx.Model
}

// TableName 表名称
func (*NotifyCategory) TableName() string {
	return TableNameNotifyCategory
}
