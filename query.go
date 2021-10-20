package yunzhanghu

import (
	"context"
)

/*查询日订单文件*/

const OrderDownloadURI = "/api/dataservice/v1/order/downloadurl"

type (
	ReqOrderDownload struct {
		OrderDate string `json:"order_date"` //格式：yyyy-MM-dd
	}
	retOrderDownload struct {
		CommonResponse
		Data OrderDownloadData `json:"data"`
	}
	OrderDownloadData struct {
		OrderDownloadUrl string `json:"order_download_url"` //订单文件下载地址
	}
)

func (y *Yunzhanghu) OrderDownload(ctx context.Context, date string) (*OrderDownloadData, error) {
	apiName := "查询日订单文件"
	req := ReqOrderDownload{OrderDate: date}
	resp := new(retOrderDownload)

	json, err := y.getJson(ctx, OrderDownloadURI, apiName, req)
	if err != nil {
		return nil, err
	}

	err = y.decodeWithError(json, resp, apiName)
	if err != nil {
		return nil, err
	}
	return &resp.Data, err
}

/*查询日流水文件
注意：每日 00:30 会自动生成日流水文件，建议每日 9:00 以后开始拉取流水
*/

const BillDownloadURI = "/api/dataservice/v2/bill/downloadurl"

type (
	ReqBillDownload struct {
		BillDate string `json:"bill_date"` //格式：yyyy-MM-dd
	}
	retBillDownload struct {
		CommonResponse
		Data BillDownloadData `json:"data"`
	}
	BillDownloadData struct {
		BillDownloadUrl string `json:"bill_download_url"` //日流水文件下载地址
	}
)

func (y *Yunzhanghu) BillDownload(ctx context.Context, date string) (*BillDownloadData, error) {
	apiName := "查询日流水文件"
	req := ReqBillDownload{
		BillDate: date,
	}
	resp := new(retBillDownload)

	json, err := y.getJson(ctx, BillDownloadURI, apiName, req)
	if err != nil {
		return nil, err
	}

	err = y.decodeWithError(json, resp, apiName)
	if err != nil {
		return nil, err
	}
	return &resp.Data, err
}

/* 查询商户充值记录 */

const RechargeRecordURI = "/api/dataservice/v2/recharge-record"

type (
	ReqRechargeRecord struct {
		BeginAt string `json:"begin_at"` //格式：yyyy-MM-dd
		EndAt   string `json:"end_at"`   //格式：yyyy-MM-dd
	}
	retRechargeRecord struct {
		CommonResponse
		Data RechargeRecordData `json:"data"`
	}
	RechargeRecordData struct {
		RechargeId        string `json:"recharge_id"`         // 充值记录 ID
		ActualAmount      string `json:"actual_amount"`       //单位：元，支持小数点后两位
		Amount            string `json:"amount"`              //单位：元，支持小数点后两位
		CreateAt          string `json:"create_at"`           //创建时间
		Remark            string `json:"remark"`              //备注
		RechargeChannel   string `json:"recharge_channel"`    //资金用途
		RechargeAccountNo string `json:"recharge_account_no"` //商户充值使用的付款银行账号
	}
)

func (y *Yunzhanghu) RechargeRecord(ctx context.Context, begin string, end string) (*RechargeRecordData, error) {
	apiName := "查询商户充值记录"
	req := ReqRechargeRecord{
		BeginAt: begin,
		EndAt:   end,
	}
	resp := new(retRechargeRecord)

	json, err := y.getJson(ctx, RechargeRecordURI, apiName, req)
	if err != nil {
		return nil, err
	}

	err = y.decodeWithError(json, resp, apiName)
	if err != nil {
		return nil, err
	}
	return &resp.Data, err
}

/* 查询日订单数据 */

const QueryOrderURI = "/api/dataservice/v1/orders"

type (
	ReqQueryOrder struct {
		OrderDate string `json:"order_date"`
		Channel   string `json:"channel"`
		DataType  string `json:"data_type"`
		Page
	}
	retQueryOrder struct {
		CommonResponse
		Data *QueryOrderData `json:"data"`
	}
	QueryOrderData struct {
		TotalNum int           `json:"total_num"`
		List     []*QueryOrder `json:"list"`
	}
	QueryOrder struct {
		BrokerId            string `json:"broker_id"`
		DealerId            string `json:"dealer_id"`
		OrderId             string `json:"order_id"`
		Ref                 string `json:"ref"`
		BatchId             string `json:"batch_id"`
		RealName            string `json:"real_name"`
		CardNo              string `json:"card_no"`
		BrokerAmount        string `json:"broker_amount"`
		BrokerFree          string `json:"broker_free"`
		Bill                string `json:"bill"`
		Status              string `json:"status"`
		StatusDetail        string `json:"status_detail"`
		StatusMessage       string `json:"status_message"`
		StatusDetailMessage string `json:"status_detail_message"`
		StatementId         string `json:"statement_id"`
		FeeStatmentId       string `json:"fee_statment_id"`
		BalStatmentId       string `json:"bal_statment_id"`
		Channel             string `json:"channel"`
		CreateAt            string `json:"create_at"`
		FinishedTime        string `json:"finished_time"`
	}
)

func (y *Yunzhanghu) QueryOrder(ctx context.Context, req ReqQueryOrder) ([]*QueryOrder, error) {
	apiName := "查询日订单数据"
	resp := new(retQueryOrder)

	json, err := y.getJson(ctx, QueryOrderURI, apiName, req)
	if err != nil {
		return nil, err
	}

	err = y.decodeWithError(json, resp, apiName)
	if err != nil || resp.Data == nil {
		return nil, err
	}

	return resp.Data.List, err
}

/* 查询日订单文件（打款和退款订单） */

const DownloadOrderDayURI = "/api/dataservice/v1/order/day/url"

type (
	ReqDownloadOrderDay struct {
		OrderDate string `json:"order_date"`
	}
	retDownloadOrderDay struct {
		CommonResponse
		Data DownloadOrderDayData `json:"data"`
	}
	DownloadOrderDayData struct {
		Url string `json:"url"`
	}
)

func (y *Yunzhanghu) DownloadOrderDay(ctx context.Context, date string) (*DownloadOrderDayData, error) {
	apiName := "查询日订单文件（打款和退款订单）"
	req := ReqDownloadOrderDay{OrderDate: date}
	resp := new(retDownloadOrderDay)

	json, err := y.getJson(ctx, DownloadOrderDayURI, apiName, req)
	if err != nil {
		return nil, err
	}

	err = y.decodeWithError(json, resp, apiName)
	if err != nil {
		return nil, err
	}
	return &resp.Data, err
}

/* 查询日流水数据 */

const QueryBillURI = "/api/dataservice/v1/bills"

type (
	ReqQueryBill struct {
		BillDate string `json:"bill_date"`
		DataType string `json:"data_type"` //如果为 encryption，则对返回的 data 进行加密
		Page
	}
	retQueryBill struct {
		CommonResponse
		Data *QueryBillData `json:"data"`
	}
	QueryBillData struct {
		TotalNum int          `json:"total_num"`
		List     []*QueryBill `json:"list"`
	}
	QueryBill struct {
		BrokerId          string `json:"broker_id"`
		DealerId          string `json:"dealer_id"`
		OrderId           string `json:"order_id"`
		Ref               string `json:"ref"`
		BrokerProductName string `json:"broker_product_name"`
		DealerProductName string `json:"dealer_product_name"`
		BizRef            string `json:"biz_ref"`
		AcctType          string `json:"acct_type"`
		Amount            string `json:"amount"`
		Balance           string `json:"balance"`
		BusinessCategory  string `json:"business_category"`
		BusinessType      string `json:"business_type"`
		ConsumptionType   string `json:"consumption_type"`
		CreateAt          string `json:"create_at"`
		Remark            string `json:"remark"`
	}
)

func (y *Yunzhanghu) QueryBill(ctx context.Context, req ReqQueryBill) ([]*QueryBill, error) {
	apiName := "查询日流水数据"
	resp := new(retQueryBill)

	json, err := y.getJson(ctx, QueryBillURI, apiName, req)
	if err != nil {
		return nil, err
	}

	err = y.decodeWithError(json, resp, apiName)
	if err != nil || resp.Data == nil {
		return nil, err
	}
	return resp.Data.List, err
}
