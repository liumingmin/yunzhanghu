package yunzhanghu

type Page struct {
	Offset int `json:"offset"` //偏移量，最小从 0 开始
	Length int `json:"length"` //每页最多返回条数，最大为 200
}

type BankCardOrderReq struct {
	OrderId   string
	RealName  string
	IdCard    string
	CardNo    string
	PhoneNo   string
	Pay       float64
	PayRemark string
}

type AliPayOrderReq struct {
	OrderId   string
	RealName  string
	IdCard    string
	CardNo    string
	PhoneNo   string
	Pay       float64
	PayRemark string
}
