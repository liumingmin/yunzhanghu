package yunzhanghu

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
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
		OrderId string  `json:"order_id"`
		Ref     string  `json:"ref"`
		Pay     float64 `json:"pay,string"`
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
		OrderId string  `json:"order_id"`
		Ref     string  `json:"ref"`
		Pay     float64 `json:"pay,string"`
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
