package models

import "github.com/shopspring/decimal"

const TableNameBadDebtRecord = "bad_debt_record"

// BadDebtRecord 坏账记录表
type BadDebtRecord struct {
	BadDebtId        int64           `json:"bad_debt_id" gorm:"bad_debt_id,primaryKey"`    // 主键ID
	MerchantId       int             `json:"merchant_id" gorm:"merchant_id"`               // 商户ID
	UserId           int64           `json:"user_id" gorm:"user_id"`                       // 用户ID
	PlatformId       int             `json:"platform_id" gorm:"platform_id"`               // 平台ID
	DebtAmount       decimal.Decimal `json:"debt_amount" gorm:"debt_amount"`               // 坏账金额（应扣除但无法扣除的金额）
	UserBalance      decimal.Decimal `json:"user_balance" gorm:"user_balance"`             // 用户当前余额
	DeficitAmount    decimal.Decimal `json:"deficit_amount" gorm:"deficit_amount"`         // 缺口金额（负余额的绝对值）
	CurrencyCode     string          `json:"currency_code" gorm:"currency_code"`           // 币种代码
	OriginalOrderNo  string          `json:"original_order_no" gorm:"original_order_no"`   // 原始业务订单号
	PlatformOrderId  string          `json:"platform_order_id" gorm:"platform_order_id"`   // 三方平台订单ID
	MerchantOrderId  string          `json:"merchant_order_id" gorm:"merchant_order_id"`   // 下游商户订单ID
	TransactionId    string          `json:"transaction_id" gorm:"transaction_id"`         // 中台交易流水号
	DebtType         int             `json:"debt_type" gorm:"debt_type"`                   // 坏账类型：1=用户余额不足,2=用户已提现,3=账户冻结,4=其他原因
	DebtReason       string          `json:"debt_reason" gorm:"debt_reason"`               // 坏账原因描述
	TriggerAction    string          `json:"trigger_action" gorm:"trigger_action"`         // 触发动作：GAME_DEBIT/CANCEL_BET/CHARGEBACK等
	Status           int             `json:"status" gorm:"status"`                         // 处理状态：1=待处理,2=部分回收,3=已回收,4=核销,5=争议中
	RecoveryAmount   decimal.Decimal `json:"recovery_amount" gorm:"recovery_amount"`       // 已回收金额
	WriteOffAmount   decimal.Decimal `json:"write_off_amount" gorm:"write_off_amount"`     // 核销金额
	RemainingAmount  decimal.Decimal `json:"remaining_amount" gorm:"remaining_amount"`     // 剩余欠款金额
	UserInfoSnapshot string          `json:"user_info_snapshot" gorm:"user_info_snapshot"` // 用户信息快照（JSON格式）
	Remark           string          `json:"remark" gorm:"remark"`                         // 备注信息
	ExtData          string          `json:"ext_data" gorm:"ext_data"`                     // 扩展数据（JSON格式）
	CreatedAt        int64           `json:"created_at" gorm:"created_at"`                 // 创建时间戳
	UpdatedAt        int64           `json:"updated_at" gorm:"updated_at"`                 // 更新时间戳
	DeletedAt        int64           `json:"deleted_at" gorm:"deleted_at"`                 // 删除时间戳（软删除）
}

// TableName 表名称
func (*BadDebtRecord) TableName() string {
	return TableNameBadDebtRecord
}
