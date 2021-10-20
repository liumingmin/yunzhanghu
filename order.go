package yunzhanghu

type Order struct {
	OrderId             string            `json:"order_id"`                 //  商户订单号，由商户保持唯⼀一性
	Pay                 float64           `json:"pay,string"`               //  金额
	BrokerId            string            `json:"broker_id"`                //  经纪公司
	DealerId            string            `json:"dealer_id"`                //  商户代码
	RealName            string            `json:"real_name"`                //  姓名
	CardNo              string            `json:"card_no"`                  //  卡号
	IdCard              string            `json:"id_card"`                  //  身份证号
	PhoneNo             string            `json:"phone_no"`                 //  手机号
	Status              OrderStatus       `json:"status"`                   //  订单状态码，详⻅见:附录1订单状态码
	StatusDetail        OrderStatusDetail `json:"status_detail"`            //  订单详细状态码，详⻅见:附录2订单详细状态码
	StatusMessage       string            `json:"status_message"`           //  状态码说明，详⻅见:附录1订单状态码
	StatusDetailMessage string            `json:"status_detail_message"`    //  状态详细状态码说明，详⻅见:附录2订单详细状态码
	BrokerAmount        string            `json:"broker_amount"`            //  综合服务主体打款金额
	Ref                 string            `json:"ref"`                      //  综合服务平台流水号，唯一
	BrokerBankBill      string            `json:"broker_bank_bill"`         //  综合服务平台打款交易流水号
	WithdrawPlatform    string            `json:"withdraw_platform"`        //  bankpay：银行卡 alipay：支付宝 wxpay：微信
	BrokerFee           float64           `json:"broker_fee,string"`        //  综合服务主体服务费
	BrokerRealFee       float64           `json:"broker_real_fee,string"`   //  余额账户支出服务费
	BrokerDeductFee     float64           `json:"broker_deduct_fee,string"` //  抵扣账户支出服务费
	UserFee             float64           `json:"user_fee,string"`          //  用户服务费
	BankName            string            `json:"bank_name"`                //  银行名称
	PayRemark           string            `json:"pay_remark"`               //  打款备注(选填，最⼤大20个字符，⼀一个汉字占2个字符，不不允许特殊字符:' " & | @ % * ( ) - : # ¥)
	CreatedAt           Time              `json:"created_at"`               //  订单接收时间
	FinishedTime        Time              `json:"finished_time"`            //  订单处理理时间
}

func (o Order) GetOrderStatus() string {
	return o.Status.Message()
}

func (o Order) GetOrderStatusDetail() string {
	return o.StatusDetail.Message()
}
