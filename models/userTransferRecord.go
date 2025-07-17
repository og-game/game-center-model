package models

import (
	"github.com/og-game/glib/stores/gormx"
	"github.com/shopspring/decimal"
)

const TableNameUserTransferRecord = "user_transfer_record"

// UserTransferRecord 用户转账记录表
type UserTransferRecord struct {
	RecordId        int64           `json:"record_id" gorm:"record_id;primaryKey"`      // 主键ID
	TransactionId   string          `json:"transaction_id" gorm:"transaction_id"`       // 中台交易ID（业务唯一标识）
	UserId          int64           `json:"user_id" gorm:"user_id"`                     // 用户ID
	MerchantUserID  string          `json:"merchant_user_id" gorm:"merchant_user_id"`   // 商户用户id
	MerchantId      int64           `json:"merchant_id" gorm:"merchant_id"`             // 商户ID
	PlatformId      int64           `json:"platform_id" gorm:"platform_id"`             // 平台ID
	GameId          int64           `json:"game_id" gorm:"game_id"`                     // 游戏ID（可选）
	TransferType    int             `json:"transfer_type" gorm:"transfer_type"`         // 转账类型：1=转入，2=转出
	Amount          decimal.Decimal `json:"amount" gorm:"amount"`                       // 转账金额
	CurrencyCode    string          `json:"currency_code" gorm:"currency_code"`         // 币种代码
	FeeAmount       decimal.Decimal `json:"fee_amount" gorm:"fee_amount"`               // 手续费金额
	ActualAmount    decimal.Decimal `json:"actual_amount" gorm:"actual_amount"`         // 实际到账金额
	Status          int             `json:"status" gorm:"status"`                       // 转账状态：1=PENDING,2=PROCESSING,3=SUCCESS,4=FAILED
	PlatformType    int             `json:"platform_type" gorm:"platform_type"`         // 平台类型：1=单一钱包，2=转账钱包
	IsInternal      int             `json:"is_internal" gorm:"is_internal"`             // 是否内部转账：1=是，2=否（第三方）
	PlatformOrderId string          `json:"platform_order_id" gorm:"platform_order_id"` // 第三方订单号
	PlatformStatus  string          `json:"platform_status" gorm:"platform_status"`     // 第三方状态
	RetryCount      int64           `json:"retry_count" gorm:"retry_count"`             // 重试次数
	MaxRetryCount   int64           `json:"max_retry_count" gorm:"max_retry_count"`     // 最大重试次数
	ErrorCode       string          `json:"error_code" gorm:"error_code"`               // 错误代码
	ErrorMessage    string          `json:"error_message" gorm:"error_message"`         // 错误信息
	FailureReason   string          `json:"failure_reason" gorm:"failure_reason"`       // 失败原因
	MerchantOrderId string          `json:"merchant_order_id" gorm:"merchant_order_id"` // 下游业务订单号
	Remark          string          `json:"remark" gorm:"remark"`                       // 备注信息
	gormx.Model
}

// TableName 表名称
func (*UserTransferRecord) TableName() string {
	return TableNameUserTransferRecord
}

// UserTransferRecord 玩家转账记录表
//type UserTransferRecord struct {
//	RecordId             int64   `json:"record_id" gorm:"record_id;primaryKey"`                // 记录id
//	MerchantOrderNo      string  `json:"merchant_order_no" gorm:"merchant_order_no"`           // 商户订单号
//	OrderNo              string  `json:"order_no" gorm:"order_no"`                             // 中台订单号
//	PlatformOrderNo      string  `json:"platform_order_no" gorm:"platform_order_no"`           // 厂商订单号
//	MerchantId           int64   `json:"merchant_id" gorm:"merchant_id"`                       // 商户id
//	UserId               int64   `json:"user_id" gorm:"user_id"`                               // 玩家id
//	PlatformId           int64   `json:"platform_id" gorm:"platform_id"`                       // 游戏平台id
//	GameId               int64   `json:"game_id" gorm:"game_id"`                               // 游戏id
//	Amount               decimal.Decimal `json:"amount" gorm:"amount"`                                 // 操作金额
//	BalanceAfterTransfer decimal.Decimal `json:"balance_after_transfer" gorm:"balance_after_transfer"` // 转账后余额
//	TransferType         int     `json:"transfer_type" gorm:"transfer_type"`                   // 转账类型 1 转入 2转出
//	SuccessAt            int64   `json:"success_at" gorm:"success_at"`                         // 转账成功时间
//	gormx.Model
//}
