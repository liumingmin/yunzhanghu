package yunzhanghu

import "context"

/* 查询可开票额度 */

const InvoiceAmountURI = "/api/invoice/v2/invoice-amount"

type (
	ReqInvoiceAmount struct {
		BrokerId string `json:"broker_id"`
		DealerId string `json:"dealer_id"`
	}
	retInvoiceAmount struct {
		CommonResponse
		Data InvoiceAmountData `json:"data"`
	}

	InvoiceAmountData struct {
		Amount            string               `json:"amount"`
		BankNameAccount   []*BankNameAccount   `json:"bank_name_account"`
		GoodsServicesName []*GoodsServicesName `json:"goods_services_name"`
	}

	BankNameAccount struct {
		Item    string `json:"item"`
		Default bool   `json:"default"`
	}

	GoodsServicesName struct {
		Item    string `json:"item"`
		Default bool   `json:"default"`
	}
)

func (y *Yunzhanghu) InvoiceAmount(ctx context.Context, req *ReqInvoiceAmount) (*InvoiceAmountData, error) {
	apiName := "查询可开票额度"
	resp := new(retInvoiceAmount)

	json, err := y.postJSON(ctx, InvoiceAmountURI, apiName, req)
	if err != nil {
		return nil, err
	}

	err = y.decodeWithError(json, resp, apiName)
	if err != nil {
		return nil, err
	}
	return &resp.Data, err
}

/* 开票申请 */

const InvoiceApplyURI = "/api/invoice/v2/apply"

type (
	ReqInvoiceApply struct {
		InvoiceApplyId  string `json:"invoice_apply_id"`
		BrokerId        string `json:"broker_id"`
		DealerId        string `json:"dealer_id"`
		Amount          string `json:"amount"`
		InvoiceType     string `json:"invoice_type"`
		BankNameAccount string `json:"bank_name_account"`
		Remark          string `json:"remark"`
	}

	retInvoiceApply struct {
		CommonResponse
		Data InvoiceApplyData `json:"data"`
	}
	InvoiceApplyData struct {
		ApplicationId string `json:"application_id"`
		Count         int    `json:"count"`
	}
)

func (y *Yunzhanghu) InvoiceApply(ctx context.Context, req *ReqInvoiceApply) (*InvoiceApplyData, error) {
	apiName := "/api/invoice/v2/apply"
	resp := new(retInvoiceApply)

	json, err := y.postJSON(ctx, InvoiceApplyURI, apiName, req)
	if err != nil {
		return nil, err
	}

	err = y.decodeWithError(json, resp, apiName)
	if err != nil {
		return nil, err
	}
	return &resp.Data, err
}

/* 查询开票申请状态 */

const InvoiceStatusURI = "/api/invoice/v2/invoice/invoice-status"

type (
	ReqInvoiceStatus struct {
		InvoiceApplyId string `json:"invoice_apply_id"`
		ApplicationId  string `json:"application_id"`
	}

	retInvoiceStatus struct {
		CommonResponse
		Data InvoiceStatusData `json:"data"`
	}
	InvoiceStatusData struct {
		Status             string   `json:"status"`
		Count              int      `json:"count"`
		PriceTaxAmount     string   `json:"price_tax_amount"`
		PriceAmount        string   `json:"price_amount"`
		TaxAmount          string   `json:"tax_amount"`
		InvoiceType        string   `json:"invoice_type"`
		CustomerName       string   `json:"customer_name"`
		CustomerTaxNum     string   `json:"customer_tax_num"`
		CustomerAddressTel string   `json:"customer_address_tel"`
		BankNameAccount    string   `json:"bank_name_account"`
		GoodsServiceName   string   `json:"goods_service_name"`
		Remark             string   `json:"remark"`
		PostType           string   `json:"post_type"`
		WaybillNumber      []string `json:"waybill_number"`
	}
)

func (y *Yunzhanghu) InvoiceStatus(ctx context.Context, req *ReqInvoiceStatus) (*InvoiceStatusData, error) {
	apiName := "查询开票申请状态"
	resp := new(retInvoiceStatus)

	json, err := y.postJSON(ctx, InvoiceStatusURI, apiName, req)
	if err != nil {
		return nil, err
	}

	err = y.decodeWithError(json, resp, apiName)
	if err != nil {
		return nil, err
	}
	return &resp.Data, err
}

/* 下载发票 PDF */

const InvoicePDFDownloadURI = "/api/invoice/v2/invoice/invoice-pdf"

type (
	ReqInvoicePDFDownload struct {
		InvoiceApplyId string `json:"invoice_apply_id"`
		Application    string `json:"application"`
	}

	retInvoicePDFDownload struct {
		CommonResponse
		Data InvoicePDFDownloadData `json:"data"`
	}
	InvoicePDFDownloadData struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}
)

func (y *Yunzhanghu) InvoicePDFDownload(ctx context.Context, req *ReqInvoicePDFDownload) (*InvoicePDFDownloadData, error) {
	apiName := "下载发票 PDF"
	resp := new(retInvoicePDFDownload)

	json, err := y.postJSON(ctx, InvoicePDFDownloadURI, apiName, req)
	if err != nil {
		return nil, err
	}

	err = y.decodeWithError(json, resp, apiName)
	if err != nil {
		return nil, err
	}
	return &resp.Data, err
}

/* 发送发票扫描件压缩包下载链接邮件 */

const InvoiceReminderEmailURI = "/api/invoice/v2/invoice/reminder/email"

type (
	ReqInvoiceReminderEmail struct {
		InvoiceApplyId string `json:"invoice_apply_id"`
		ApplicationId  string `json:"application_id"`
	}

	retInvoiceReminderEmail struct {
		CommonResponse
	}
)

func (y *Yunzhanghu) InvoiceReminderEmail(ctx context.Context, req *ReqInvoiceReminderEmail) error {
	apiName := "发送发票扫描件压缩包下载链接邮件"
	resp := new(retInvoiceReminderEmail)

	json, err := y.postJSON(ctx, InvoiceReminderEmailURI, apiName, req)
	if err != nil {
		return err
	}

	return y.decodeWithError(json, resp, apiName)
}

/* 查询商户已开具发票金额和待开具发票金额 */

const InvoiceStatURI = ""

type (
	ReqInvoiceStat struct {
		BrokerId string `json:"broker_id"`
		DealerId string `json:"dealer_id"`
		Year     int    `json:"year"`
	}

	retInvoiceStat struct {
		CommonResponse
		Data InvoiceStatData `json:"data"`
	}
	InvoiceStatData struct {
		BrokerId    string `json:"broker_id"`
		DealerId    string `json:"dealer_id"`
		Invoiced    string `json:"invoiced"`
		NotInvoiced string `json:"not_invoiced"`
	}
)

func (y *Yunzhanghu) InvoiceStat(ctx context.Context, req ReqInvoiceStat) (*InvoiceStatData, error) {
	apiName := "查询商户已开具发票金额和待开具发票金额"
	resp := new(retInvoiceStat)

	json, err := y.getJson(ctx, InvoiceStatURI, apiName, req)
	if err != nil {
		return nil, err
	}

	err = y.decodeWithError(json, resp, apiName)
	if err != nil {
		return nil, err
	}
	return &resp.Data, err
}
