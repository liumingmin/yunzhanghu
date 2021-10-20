package yunzhanghu

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

const (
	ORDER_CHANNEL_BANK   = "银行卡"
	ORDER_CHANNEL_ALIPAY = "支付宝"
	ORDER_CHANNEL_WEIXIN = "微信"

	WITHDRAW_PLATFORM_BANKPAY = "bankpay"
	WITHDRAW_PLATFORM_ALIPAY  = "alipay"
	WITHDRAW_PLATFORM_WXPAY   = "wxpay"
)

const (
	paymentOrderRealtimeURI = "/api/payment/v1/order-realtime"
)

type (
	reqOrderRealtime struct {
		OrderId   string  `json:"order_id"`             // 商户订单号，由商户保持唯⼀一性(必填)
		DealerId  string  `json:"dealer_id"`            // 商户代码(必填)
		BrokerId  string  `json:"broker_id"`            // 经纪公司(必填)
		RealName  string  `json:"real_name"`            // 银⾏行行开户姓名(必填)
		CardNo    string  `json:"card_no"`              // 银⾏行行开户卡号(必填)
		PhoneNo   string  `json:"phone_no"`             // ⽤用户或联系⼈人⼿手机号
		IdCard    string  `json:"id_card"`              // 银⾏行行开户身份证号
		Pay       float64 `json:"pay,string"`           // 打款⾦金金额(必填)
		PayRemark string  `json:"pay_remark,omitempty"` // 打款备注(选填，最⼤大20个字符，⼀个汉字占2个字符，不不允许特殊字符:' " & | @ % * ( ) - : # ¥)
		NotifyUrl string  `json:"notify_url,omitempty"` // 异步通知地址
	}

	OrderRealtimeData struct {
		OrderID string  `json:"order_id"`   // 商户订单号，原值返回
		Ref     string  `json:"ref"`        // 综合服务平台流水号，唯一
		Pay     float64 `json:"pay,string"` // 打款金额
	}

	retOrderRealtime struct {
		*CommonResponse
		Data OrderRealtimeData `json:"data"`
	}
)

func (y *Yunzhanghu) OrderRealTime(ctx context.Context, order BankCardOrderReq) (*OrderRealtimeData, error) {
	var (
		apiName = "银行卡实时下单"
		req     = &reqOrderRealtime{
			OrderId:   order.OrderId,
			DealerId:  y.Dealer,
			BrokerId:  y.Broker,
			RealName:  order.RealName,
			CardNo:    order.CardNo,
			PhoneNo:   order.PhoneNo,
			IdCard:    order.IdCard,
			Pay:       order.Pay,
			PayRemark: order.PayRemark,
			NotifyUrl: y.PayNotifyUrl,
		}
		ret = new(retOrderRealtime)
	)
	responseBytes, err := y.postForm(ctx, paymentOrderRealtimeURI, apiName, req)
	if err != nil {
		return nil, err
	}
	if err = y.decodeWithError(responseBytes, ret, apiName); err != nil {
		return nil, err
	}
	return &ret.Data, nil
}

const (
	paymentOrderAlipayURI = "/api/payment/v1/order-alipay"
)

type (
	reqOrderAlipay struct {
		OrderId   string  `json:"order_id"`             // 商户订单号，由商户保持唯⼀一性(必填)
		DealerId  string  `json:"dealer_id"`            // 商户代码(必填)
		BrokerId  string  `json:"broker_id"`            // 经纪公司(必填)
		RealName  string  `json:"real_name"`            // 姓名(必填)
		IdCard    string  `json:"id_card"`              // 身份证(必填)
		CardNo    string  `json:"card_no"`              // 收款⼈人⽀支付宝账户(必填)
		PhoneNo   string  `json:"phone_no"`             // ⽤用户或联系⼈人⼿手机号
		CheckName string  `json:"check_name"`           // 校验支付宝姓名，固定值：Check
		Pay       float64 `json:"pay,string"`           // 打款⾦金金额(单位为元, 必填)
		Notes     string  `json:"notes,omitempty"`      // 备注信息(选填)
		PayRemark string  `json:"pay_remark,omitempty"` // 打款备注(选填，最⼤大20个字符，⼀一个汉字占2个字符，不不允许特殊字符:' " & | @ % * ( ) - : // ¥)
		NotifyUrl string  `json:"notify_url,omitempty"` // 异步通知地址
	}

	OrderAlipayData struct {
		OrderID string  `json:"order_id"`   // 商户订单号，原值返回
		Ref     string  `json:"ref"`        // 综合服务平台流水号，唯一
		Pay     float64 `json:"pay,string"` // 打款金额
	}

	retOrderAlipay struct {
		CommonResponse
		Data OrderAlipayData `json:"data"`
	}
)

func (y *Yunzhanghu) OrderAlipay(ctx context.Context, order AliPayOrderReq) (*OrderAlipayData, error) {
	var (
		apiName = "支付宝实时下单"
		req     = &reqOrderAlipay{
			DealerId:  y.Dealer,
			BrokerId:  y.Broker,
			OrderId:   order.OrderId,
			RealName:  order.RealName,
			IdCard:    order.IdCard,
			CardNo:    order.CardNo,
			PhoneNo:   order.PhoneNo,
			CheckName: "Check",
			Pay:       order.Pay,
			PayRemark: order.PayRemark,
			NotifyUrl: y.PayNotifyUrl,
		}
		ret = new(retOrderAlipay)
	)
	responseBytes, err := y.postForm(ctx, paymentOrderAlipayURI, apiName, req)
	if err != nil {
		return nil, err
	}
	if err = y.decodeWithError(responseBytes, ret, apiName); err != nil {
		return nil, err
	}
	return &ret.Data, nil
}

const (
	orderWxpayURI = "/api/payment/v1/order-wxpay"
)

type (
	reqOrderWxpay struct {
		OrderID   string  `json:"order_id"`   // 商户订单号，由商户保持唯⼀一性(必填)
		DealerId  string  `json:"dealer_id"`  // 商户代码(必填)
		BrokerId  string  `json:"broker_id"`  // 经纪公司(必填)
		RealName  string  `json:"real_name"`  // 姓名(必填)
		IDCard    string  `json:"id_card"`    // 身份证(必填)
		Openid    string  `json:"openid"`     // wx2319u9jk231ad21
		PhoneNo   string  `json:"phone_no"`   // 用户或联系⼈人⼿手机号
		WxAppID   string  `json:"wx_app_id"`  // 商户微信 AppID
		WxpayMode string  `json:"wxpay_mode"` // 微信打款模式
		Pay       float64 `json:"pay,string"` // 金额
		PayRemark string  `json:"pay_remark"` // 打款备注
	}
	OrderWxpayData struct {
		OrderID string  `json:"order_id"`   // 商户订单号，原值返回
		Ref     string  `json:"ref"`        // 综合服务平台流水号，唯一
		Pay     float64 `json:"pay,string"` // 打款金额
	}
	retOrderWxpay struct {
		CommonResponse
		Data OrderWxpayData `json:"data"`
	}
)

func (y *Yunzhanghu) OrderWxpay(ctx context.Context, order *WxPayOrderReq) (*OrderWxpayData, error) {
	var (
		input = &reqOrderWxpay{
			OrderID:   order.OrderID,
			DealerId:  y.Dealer,
			BrokerId:  y.Broker,
			RealName:  order.RealName,
			IDCard:    order.IDCard,
			Openid:    order.Openid,
			PhoneNo:   order.PhoneNo,
			WxAppID:   y.WxAppID,
			WxpayMode: "transfer",
			Pay:       order.Pay,
			PayRemark: order.PayRemark,
		}
		output  = new(retOrderWxpay)
		apiName = "微信实时下单"
	)
	responseBytes, err := y.getJson(ctx, orderWxpayURI, apiName, input)
	if err != nil {
		return nil, err
	}
	if err = y.decodeWithError(responseBytes, output, apiName); err != nil {
		return nil, err
	}
	return &output.Data, nil
}

const (
	queryRealtimeOrderURI = "/api/payment/v1/query-realtime-order"
)

type (
	reqQueryRealtimeOrder struct {
		OrderId  string `json:"order_id"`  // 商户订单号
		Channel  string `json:"channel"`   // 银⾏行行卡，⽀支付宝，微信(不不填时为银⾏行行卡订单查询)(选填)
		DataType string `json:"data_type"` // 如果为encryption，则对返回的data进⾏行行加密(选填)
	}

	encryOrderData struct {
		EncryData string `json:"encry_data"` //  当且仅当data_type为encryption时，返回且仅返回该加密数据字段
	}

	retQueryRealtimeOrder struct {
		CommonResponse
		Data encryOrderData `json:"data"`
	}
)

func (y *Yunzhanghu) QueryRealtimeOrder(ctx context.Context, orderId, channel string) (*Order, error) {
	var (
		apiName = "查询⼀个订单"
		resp    = new(retQueryRealtimeOrder)
		req     = &reqQueryRealtimeOrder{
			OrderId:  orderId,
			Channel:  channel,
			DataType: "encryption",
		}
	)

	respnseBytes, err := y.getJson(ctx, queryRealtimeOrderURI, apiName, req)
	if err != nil {
		return nil, err
	}
	if err = y.decodeWithError(respnseBytes, resp, apiName); err != nil {
		return nil, err
	}

	var order *Order
	err = DecryptB64TriDesTo(resp.Data.EncryData, []byte(y.DesKey), &order)
	if err != nil {
		return nil, err
	}

	return order, nil
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type PaymentNotify struct {
	Data       Order  `json:"data"`
	NotifyID   string `json:"notify_id"`   // 14732279660721952
	NotifyTime string `json:"notify_time"` // 2020-05-25 11:49:34
}

func (y *Yunzhanghu) NotifyCallbackHandler(ctx *gin.Context) {
	var resp *CallbackResponse
	if err := ctx.MustBindWith(&resp, binding.FormPost); err != nil || resp == nil || resp.Data == "" {
		ctx.String(http.StatusOK, RESP_STR_FAILD)
		return
	}

	var content *PaymentNotify
	err := DecryptB64TriDesTo(resp.Data, []byte(y.DesKey), &content)
	if err != nil {
		ctx.String(http.StatusOK, RESP_STR_FAILD)
		return
	}

	ok := y.PayNotifyCallback(ctx, content)
	if !ok {
		ctx.String(http.StatusOK, RESP_STR_FAILD)
		return
	}

	ctx.String(http.StatusOK, RESP_STR_SUCCESS)
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

const (
	queryAccountsURI = "/api/payment/v1/query-accounts"
)

type (
	reqQueryAccounts struct {
		DealerID string `json:"dealer_id"` // 商户ID
	}
	AccountBalanceData struct {
		BrokerID         string  `json:"broker_id"`                 // 综合服务主体 ID
		AcctBalance      float64 `json:"acct_balance,string"`       // 余额账户余额
		AlipayBalance    float64 `json:"alipay_balance,string"`     // 支付宝余额
		BankCardBalance  float64 `json:"bank_card_balance,string"`  // 银行卡余额
		WxpayBalance     float64 `json:"wxpay_balance,string"`      // 微信余额
		IsAlipay         bool    `json:"is_alipay"`                 // 是否开通支付宝通道
		IsBankCard       bool    `json:"is_bank_card"`              // 是否开通银行卡通道
		IsWxpay          bool    `json:"is_wxpay"`                  // 是否开通微信通道
		RebateFeeBalance float64 `json:"rebate_fee_balance,string"` // 服务费返点余额
		TotalBalance     float64 `json:"total_balance,string"`      // 总余额
	}
	retQueryAccounts struct {
		CommonResponse
		Data struct {
			DealerInfos []*AccountBalanceData `json:"dealer_infos"`
		} `json:"data"`
	}
)

func (y *Yunzhanghu) QueryAccounts(ctx context.Context) ([]*AccountBalanceData, error) {
	var (
		input = &reqQueryAccounts{
			DealerID: y.Dealer,
		}
		output  = new(retQueryAccounts)
		apiName = "查询商户余额"
	)
	responseBytes, err := y.getJson(ctx, queryAccountsURI, apiName, input)
	if err != nil {
		return nil, err
	}
	if err = y.decodeWithError(responseBytes, output, apiName); err != nil {
		return nil, err
	}
	return output.Data.DealerInfos, nil
}

const (
	receiptFileURI = "/api/payment/v1/receipt/file"
)

type (
	reqReceiptFile struct {
		OrderID string `json:"order_id"`
		Ref     string `json:"ref"`
	}
	ReceiptFileData struct {
		ExpireTime string `json:"expire_time"` // 精确到秒 2020-09-05 18:36:37
		FileName   string `json:"file_name"`   // 电子回单名称 w_140_receipt_20180726.png
		URL        string `json:"url"`         // 电子回单的下载链接 https://xxxx
	}
	retReceiptFile struct {
		CommonResponse
		Data ReceiptFileData `json:"data"`
	}
)

func (y *Yunzhanghu) ReceiptFile(ctx context.Context, orderId, ref string) (*ReceiptFileData, error) {
	var (
		input = &reqReceiptFile{
			OrderID: orderId,
			Ref:     ref,
		}
		output  = new(retReceiptFile)
		apiName = "查询电子回单"
	)
	responseBytes, err := y.getJson(ctx, receiptFileURI, apiName, input)
	if err != nil {
		return nil, err
	}
	if err = y.decodeWithError(responseBytes, output, apiName); err != nil {
		return nil, err
	}
	return &output.Data, nil
}

const (
	cancelOrderURI = "/api/payment/v1/order/fail"
)

type (
	reqCancelOrder struct {
		Channel string `json:"channel"`  // 银行卡
		OrderID string `json:"order_id"` // 202009010016562012987
		Ref     string `json:"ref"`      // 176826728298982
	}
	CancelOrderData struct {
		Ok bool `json:"ok,string"` // true
	}
	retCancelOrder struct {
		CommonResponse
		Data CancelOrderData `json:"data"`
	}
)

func (y *Yunzhanghu) CancelOrder(ctx context.Context, orderId, ref, channel string) (*CancelOrderData, error) {
	var (
		input = &reqCancelOrder{
			Channel: channel,
			OrderID: orderId,
			Ref:     ref,
		}
		output  = new(retCancelOrder)
		apiName = "取消待打款订单"
	)
	responseBytes, err := y.postForm(ctx, cancelOrderURI, apiName, input)
	if err != nil {
		return nil, err
	}
	if err = y.decodeWithError(responseBytes, output, apiName); err != nil {
		return nil, err
	}
	return &output.Data, nil
}

const (
	queryVaAccountURI = "/api/payment/v1/va-account"
)

type (
	reqQueryVaAccount struct {
		BrokerID string `json:"broker_id"` // yiyun73
		DealerID string `json:"dealer_id"` // 12345678
	}
	VaAccountData struct {
		AcctName       string `json:"acct_name"`        // 云账户
		AcctNo         string `json:"acct_no"`          // 2772510300399876543210
		BankName       string `json:"bank_name"`        // 中国银行
		DealerAcctName string `json:"dealer_acct_name"` // 企业测试名称
	}
	retQueryVaAccount struct {
		CommonResponse
		Data VaAccountData `json:"data"`
	}
)

func (y *Yunzhanghu) QueryVaAccount(ctx context.Context) (*VaAccountData, error) {
	var (
		input = &reqQueryVaAccount{
			BrokerID: y.Broker,
			DealerID: y.Dealer,
		}
		output  = new(retQueryVaAccount)
		apiName = "查询商户VA账户信息"
	)
	responseBytes, err := y.getJson(ctx, queryVaAccountURI, apiName, input)
	if err != nil {
		return nil, err
	}
	if err = y.decodeWithError(responseBytes, output, apiName); err != nil {
		return nil, err
	}
	return &output.Data, nil
}
