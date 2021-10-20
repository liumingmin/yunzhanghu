package yunzhanghu

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

const (
	SIGN_CERT_TYPE_IDCARD      = 0 //身份证
	SIGN_CERT_TYPE_HK_PASSPORT = 2 //港澳居民来往内地通行证
	SIGN_CERT_TYPE_PASSPORT    = 3 //护照
	SIGN_CERT_TYPE_TW_PASSPORT = 5 //台湾居民来往大陆通行证
)

const (
	SIGN_STATUS_NO        = 0 //未签约
	SIGN_STATUS_YES       = 1 //已签约
	SIGN_STATUS_CANCEL    = 2 //已解约
	SIGN_STATUS_NOT_FOUND = 9 //不存在签约关系
)

const (
	h5PreSignURI = "/api/sdk/v1/presign"
)

type (
	reqH5PreSign struct {
		DealerID        string `json:"dealer_id"`        // 商户代码
		BrokerID        string `json:"broker_id"`        // 综合服务主体 ID
		RealName        string `json:"real_name"`        // 身份证姓名
		IDCard          string `json:"id_card"`          // 身份证号码
		CertificateType int    `json:"certificate_type"` // 证件类型
	}

	H5PreSignData struct {
		Uid    string `json:"uid"`
		Token  string `json:"token"`
		Status int64  `json:"status"`
	}

	retH5PreSign struct {
		CommonResponse
		Data H5PreSignData `json:"data"`
	}
)

func (y *Yunzhanghu) H5PreSign(ctx context.Context, realName, idCard string, certType int) (*H5PreSignData, error) {
	var (
		input = &reqH5PreSign{
			DealerID:        y.Dealer,
			BrokerID:        y.Broker,
			RealName:        realName,
			IDCard:          idCard,
			CertificateType: certType,
		}
		output  = new(retH5PreSign)
		apiName = "H5预申请签约"
	)
	responseBytes, err := y.postForm(ctx, h5PreSignURI, apiName, input)
	if err != nil {
		return nil, err
	}
	if err = y.decodeWithError(responseBytes, output, apiName); err != nil {
		return nil, err
	}
	return &output.Data, nil
}

const (
	h5SignURI = "/api/sdk/v1/sign/h5"
)

type (
	reqH5Sign struct {
		Token       string `json:"token"`        // H5 签约 token
		Color       string `json:"color"`        // H5 页面主题颜色
		URL         string `json:"url"`          // 回调 url 地址
		RedirectURL string `json:"redirect_url"` // 跳转 url
	}
	H5SignData struct {
		URL string `json:"url"` //h5 签约页面 url
	}
	retH5Sign struct {
		CommonResponse
		Data H5SignData `json:"data"`
	}
)

func (y *Yunzhanghu) H5Sign(ctx context.Context, token, color string) (*H5SignData, error) {
	var (
		input = &reqH5Sign{
			Token:       token,
			Color:       color,
			URL:         y.H5SignNotifyUrl,
			RedirectURL: y.H5SignRedirectURL,
		}
		output  = new(retH5Sign)
		apiName = "H5签约"
	)
	responseBytes, err := y.getJson(ctx, h5SignURI, apiName, input)
	if err != nil {
		return nil, err
	}
	if err = y.decodeWithError(responseBytes, output, apiName); err != nil {
		return nil, err
	}
	return &output.Data, nil
}

const (
	h5SignStatusURI = "/api/sdk/v1/sign/user/status"
)

type (
	reqH5SignStatus struct {
		DealerId string `json:"dealer_id"` // 商户代码(必填)
		BrokerId string `json:"broker_id"` // 经纪公司(必填)
		RealName string `json:"real_name"` // 姓名(必填)
		IdCard   string `json:"id_card"`   // 身份证(必填)
	}
	H5SignStatusData struct {
		SignedAt string `json:"signed_at"` // 2020-07-05 15:15:15
		Status   int64  `json:"status"`    // 1
	}
	retH5SignStatus struct {
		CommonResponse
		Data H5SignStatusData `json:"data"`
	}
)

func (y *Yunzhanghu) H5SignStatus(ctx context.Context, realName, idCard string) (*H5SignStatusData, error) {
	var (
		input = &reqH5SignStatus{
			DealerId: y.Dealer,
			BrokerId: y.Broker,
			RealName: realName,
			IdCard:   idCard,
		}
		output  = new(retH5SignStatus)
		apiName = "获取用户签约状态"
	)
	responseBytes, err := y.getJson(ctx, h5SignStatusURI, apiName, input)
	if err != nil {
		return nil, err
	}
	if err = y.decodeWithError(responseBytes, output, apiName); err != nil {
		return nil, err
	}
	return &output.Data, nil
}

const h5SignReleaseURI = "/api/sdk/v1/sign/release"

type (
	reqSignRelease struct {
		DealerID        string `json:"dealer_id"`        // 商户代码
		BrokerID        string `json:"broker_id"`        // 综合服务主体 ID
		RealName        string `json:"real_name"`        // 身份证姓名
		IDCard          string `json:"id_card"`          // 身份证号码
		CertificateType int    `json:"certificate_type"` // 证件类型
	}
	SignReleaseData struct {
		status string `json:"status"`
	}
	retSignRelease struct {
		CommonResponse
		Data SignReleaseData `json:"data"`
	}
)

//H5 对接测试解约接口
func (y *Yunzhanghu) H5SignRelease(ctx context.Context, realName, idCard string, certType int) (*SignReleaseData, error) {
	var (
		input = &reqSignRelease{
			DealerID:        y.Dealer,
			BrokerID:        y.Broker,
			RealName:        realName,
			IDCard:          idCard,
			CertificateType: certType,
		}
		output  = new(retSignRelease)
		apiName = "H5测试解约接口"
	)
	responseBytes, err := y.postForm(ctx, h5SignReleaseURI, apiName, input)
	if err != nil {
		return nil, err
	}
	if err = y.decodeWithError(responseBytes, output, apiName); err != nil {
		return nil, err
	}
	return &output.Data, nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
type H5SignNotify struct {
	BrokerID string `json:"broker_id"` // yiyun
	DealerID string `json:"dealer_id"` // 01720374
	IDCard   string `json:"id_card"`   // 360232199009115318
	Phone    string `json:"phone"`     //
	RealName string `json:"real_name"` // 王小明
}

func (y *Yunzhanghu) H5SignCallbackHandler(ctx *gin.Context) {
	var resp *CallbackResponse
	if err := ctx.MustBindWith(&resp, binding.FormPost); err != nil || resp == nil || resp.Data == "" {
		ctx.String(http.StatusOK, RESP_STR_FAILD)
		return
	}

	var h5SignResp *H5SignNotify
	err := DecryptB64TriDesTo(resp.Data, []byte(y.DesKey), &h5SignResp)
	if err != nil {
		ctx.String(http.StatusOK, RESP_STR_FAILD)
		return
	}

	ok := y.H5SignNotifyCallback(ctx, h5SignResp)
	if !ok {
		ctx.String(http.StatusOK, RESP_STR_FAILD)
		return
	}

	ctx.String(http.StatusOK, RESP_STR_SUCCESS)
}
