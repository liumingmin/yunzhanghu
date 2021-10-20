package yunzhanghu

type Page struct {
	Offset int `json:"offset"` //偏移量，最小从 0 开始
	Length int `json:"length"` //每页最多返回条数，最大为 200
}

type BankCardOrderReq struct {
	OrderId   string  //订单ID
	RealName  string  //姓名
	IdCard    string  //身份证号
	CardNo    string  //银行卡号
	PhoneNo   string  //手机号
	Pay       float64 //付款金额
	PayRemark string  //备注
}

type AliPayOrderReq struct {
	OrderId   string  //订单ID
	RealName  string  //姓名
	IdCard    string  //身份证号
	CardNo    string  //银行卡号
	PhoneNo   string  //手机号
	Pay       float64 //付款金额
	PayRemark string  //备注
}
