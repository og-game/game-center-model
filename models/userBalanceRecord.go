package models

import (
	"github.com/og-game/glib/stores/gormx"
	"github.com/shopspring/decimal"
)

const TableNameUserBalanceRecode = "user_balance_record"

// UserBalanceRecord 用户余额变动记录表（账变记录表）
type UserBalanceRecord struct {
	RecordId        int64           `json:"record_id" gorm:"record_id,primaryKey"`      // 主键ID
	MerchantId      int64           `json:"merchant_id" gorm:"merchant_id"`             // 商户ID
	MerchantUserID  string          `json:"merchant_user_id" gorm:"merchant_user_id"`   // 商户用户id
	UserId          int64           `json:"user_id" gorm:"user_id"`                     // 用户ID
	PlatformId      int64           `json:"platform_id" gorm:"platform_id"`             // 平台ID
	TransactionType int             `json:"transaction_type" gorm:"transaction_type"`   // 交易类型：1=转入游戏,2=从游戏转出,3=投注,4=结算派奖,5=押金,6=返还押金,7=取消投注,8=结算撤单,9=重新派奖,10=调整金额
	Amount          decimal.Decimal `json:"amount" gorm:"amount"`                       // 变动金额（正数表示增加，负数表示减少）
	BalanceBefore   decimal.Decimal `json:"balance_before" gorm:"balance_before"`       // 变动前余额
	BalanceAfter    decimal.Decimal `json:"balance_after" gorm:"balance_after"`         // 变动后余额
	CurrencyCode    string          `json:"currency_code" gorm:"currency_code"`         // 币种代码
	RelatedOrderId  string          `json:"related_order_id" gorm:"related_order_id"`   // 关联业务记录ID
	PlatformOrderId string          `json:"platform_order_id" gorm:"platform_order_id"` // 平台订单ID（第三方平台的订单号）
	MerchantOrderId string          `json:"merchant_order_id" gorm:"merchant_order_id"` // 下游商户订单ID
	TransactionId   string          `json:"transaction_id" gorm:"transaction_id"`       // 中台交易流水号
	Description     string          `json:"description" gorm:"description"`             // 交易描述
	Remark          string          `json:"remark" gorm:"remark"`                       // 备注信息
	ClientIp        string          `json:"client_ip" gorm:"client_ip"`                 // 客户端IP
	DeviceID        string          `json:"device_id" gorm:"device_id"`                 // '设备ID'
	DeviceOS        string          `json:"device_os" gorm:"device_os"`                 // '设备操作系统'
	ExtData         string          `json:"ext_data" gorm:"ext_data"`                   // 扩展数据（JSON格式）
	gormx.Model
}

// TableName 表名称
func (*UserBalanceRecord) TableName() string {
	return TableNameUserBalanceRecode
}
