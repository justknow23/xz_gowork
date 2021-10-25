package model

type Pre_insurance struct {
	Id                  int    `orm:"id" json:"id"`                                     // id
	Ver                 int    `orm:"ver" json:"ver"`                                   // 版本
	BookorderId         int    `orm:"bookorder_id" json:"bookorder_id"`                 // 订单id
	ParentOrderId       int    `orm:"parent_order_id" json:"parent_order_id"`           // 父订单id
	LuId                int    `orm:"lu_id" json:"lu_id"`                               // 房源ID
	LandlordId          int    `orm:"landlord_id" json:"landlord_id"`                   // 房东ID
	OperatorId          int    `orm:"operator_id" json:"operator_id"`                   // 预订人ID
	CheckinTime         string `orm:"checkin_time" json:"checkin_time"`                 // 入住时间
	CheckoutTime        string `orm:"checkout_time" json:"checkout_time"`               // 离店时间
	CheckRuntime        string `orm:"check_runtime" json:"check_runtime"`               // 时区运算结果运行时间
	TimeZone            string `orm:"time_zone" json:"time_zone"`                       // 时区
	BookorderChannel    string `orm:"bookorder_channel" json:"bookorder_channel"`       // 下单渠道
	BookorderProcessing string `orm:"bookorder_processing" json:"bookorder_processing"` // 当前处理状态：init/process/succeed/fail/cancel
	CreateTime          string `orm:"create_time" json:"create_time"`                   // 创建时间
	UpdateTime          string `orm:"update_time" json:"update_time"`                   // 更新时间
}

func (*Pre_insurance) TableName() string {
	return "pre_insurance"
}
