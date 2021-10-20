package yunzhanghu

import "context"

/* 下载个税扣缴明细表 */

const TaxFileDownloadURI = ""

type (
	ReqTaxFileDownload struct {
		DealerId  string `json:"dealer_id"`
		EntId     string `json:"ent_id"`
		YearMonth string `json:"year_month"`
	}

	retTaxFileDownload struct {
		CommonResponse
		Data *TaxFileDownloadData `json:"data"`
	}

	TaxFileDownloadData struct {
		FileInfo []*FileInfo `json:"file_info"`
	}

	FileInfo struct {
		Name string `json:"name"`
		Url  string `json:"url"`
		Pwd  string `json:"pwd"`
	}
)

func (y *Yunzhanghu) TaxFileDownload(ctx context.Context, req *ReqTaxFileDownload) ([]*FileInfo, error) {
	apiName := "下载个税扣缴明细表"
	resp := new(retTaxFileDownload)

	json, err := y.postJSON(ctx, TaxFileDownloadURI, apiName, req)
	if err != nil {
		return nil, err
	}

	err = y.decodeWithError(json, resp, apiName)
	if err != nil || resp.Data == nil {
		return nil, err
	}
	return resp.Data.FileInfo, err
}

/* 查询纳税人是否为跨集团用户 */

const UserCrossURI = "/api/tax/v1/user/cross"

type (
	ReqUserCross struct {
		DealerId string `json:"dealer_id"`
		Year     string `json:"year"`
		IdCard   string `json:"id_card"`
		EntId    string `json:"ent_id"`
	}

	retUserCross struct {
		CommonResponse
		Data UserCrossData `json:"data"`
	}
	UserCrossData struct {
		IsCross bool `json:"is_cross"`
	}
)

func (y *Yunzhanghu) UserCross(ctx context.Context, req *ReqUserCross) (*UserCrossData, error) {
	apiName := "查询纳税人是否为跨集团用户"
	resp := new(retUserCross)

	json, err := y.postJSON(ctx, UserCrossURI, apiName, req)
	if err != nil {
		return nil, err
	}

	err = y.decodeWithError(json, resp, apiName)
	if err != nil {
		return nil, err
	}
	return &resp.Data, err
}
