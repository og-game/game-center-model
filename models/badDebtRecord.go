package models

import (
	"github.com/og-game/glib/stores/gormx"
	"github.com/shopspring/decimal"
)

const TableNameBadDebtRecord = "bad_debt_record"

// BadDebtRecord 坏账记录表
type BadDebtRecord struct {
	BadDebtID          int64           `json:"bad_debt_id" gorm:"column:bad_debt_id;primaryKey;autoIncrement"`                             // 主键ID
	MerchantID         int64           `json:"merchant_id" gorm:"column:merchant_id"`                                                      // 商户ID
	MerchantUserID     string          `json:"merchant_user_id" gorm:"column:merchant_user_id;type:varchar(64);not null;default:''"`       // 商户用户ID
	CurrencyCode       string          `json:"currency_code" gorm:"column:currency_code;type:varchar(10);not null;default:''"`             // 币种代码
	UserID             int64           `json:"user_id" gorm:"column:user_id"`                                                              // 用户ID
	PlatformID         int64           `json:"platform_id" gorm:"column:platform_id"`                                                      // 平台ID[分流之后的真实平台ID]
	GameID             int64           `json:"game_id" gorm:"column:game_id"`                                                              // 游戏ID[分流之后的真实游戏ID]
	OriginalPlatformID int64           `json:"original_platform_id" gorm:"column:original_platform_id"`                                    // 原始平台ID[分流之前的用户点击的平台ID]
	OriginalGameID     int64           `json:"original_game_id" gorm:"column:original_game_id"`                                            // 原始游戏ID[分流之前的用户点击的游戏ID]
	DebtAmount         decimal.Decimal `json:"debt_amount" gorm:"column:debt_amount;type:decimal(20,6);not null;default:0.000000"`         // 坏账金额（应扣除的原始金额）
	BalanceBefore      decimal.Decimal `json:"balance_before" gorm:"column:balance_before;type:decimal(20,6);not null;default:0.000000"`   // 执行前余额（这里记录有负值，实际上平台系统的用户余额是无负数的）
	DeficitAmount      decimal.Decimal `json:"deficit_amount" gorm:"column:deficit_amount;type:decimal(20,6);not null;default:0.000000"`   // 缺口金额（应扣除的原始金额-执行前余额[如果执行前余额为负，则减去0]）
	BalanceAfter       decimal.Decimal `json:"balance_after" gorm:"column:balance_after;type:decimal(20,6);not null;default:0.000000"`     // 执行后余额（这里记录有负值，实际上中台系统的用户余额是无负数的）
	OriginalOrderID    string          `json:"original_order_id" gorm:"column:original_order_no;type:varchar(128);not null;default:''"`    // 原始业务订单号
	PlatformOrderID    string          `json:"platform_order_id" gorm:"column:platform_order_id;type:varchar(128);default:''"`             // 三方平台订单ID
	MerchantOrderID    string          `json:"merchant_order_id" gorm:"column:merchant_order_id;type:varchar(128);default:''"`             // 下游商户订单ID
	TransactionID      string          `json:"transaction_id" gorm:"column:transaction_id;type:varchar(128);not null;default:''"`          // 中台交易流水号
	DebtReason         string          `json:"debt_reason" gorm:"column:debt_reason;type:varchar(512);not null;default:''"`                // 坏账原因描述
	TriggerAction      int8            `json:"trigger_action" gorm:"column:trigger_action;type:tinyint;not null;default:0"`                // 触发动作：1=结算撤单,2=重新结算,3=调整扣款
	Status             int8            `json:"status" gorm:"column:status;type:tinyint;not null;default:1"`                                // 处理状态：1=待处理,2=部分回收,3=全部回收
	RecoveryAmount     decimal.Decimal `json:"recovery_amount" gorm:"column:recovery_amount;type:decimal(20,6);not null;default:0.000000"` // 已回收金额
	Remark             string          `json:"remark" gorm:"column:remark;type:varchar(1024);default:''"`                                  // 备注信息
	ExtData            string          `json:"ext_data" gorm:"column:ext_data;type:text"`                                                  // 扩展数据（JSON格式）-可能有用户的当时的快照信息等等
	NotifyCount        int8            `json:"notify_count" gorm:"column:notify_count;type:tinyint;not null;default:0"`                    // 通知次数
	NotifyStatus       int8            `json:"notify_status" gorm:"column:notify_status;type:tinyint;not null;default:1"`                  // 通知状态：1=等待通知,2=通知中,3=通知成功,4=通知失败
	NotifyResponse     string          `json:"notify_response" gorm:"column:notify_response;type:text"`                                    // 通知响应
	NextNotifyTime     int64           `json:"next_notify_time" gorm:"column:next_notify_time;type:int unsigned;not null;default:0"`       // 下次通知时间
	LastNotifyTime     int64           `json:"last_notify_time" gorm:"column:last_notify_time;type:int unsigned;not null;default:0"`       // 最后通知时间
	gormx.Model
}

// TableName 表名称
func (*BadDebtRecord) TableName() string {
	return TableNameBadDebtRecord
}
