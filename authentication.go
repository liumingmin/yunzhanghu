package yunzhanghu

import "context"

/* 银行卡四要素鉴权请求（下发短信验证码） */

const VerifyRequestURI = "/authentication/verify-request"

type (
	ReqVerifyRequest struct {
		CardNo   string `json:"card_no"`
		IdCard   string `json:"id_card"`
		RealName string `json:"real_name"`
		Mobile   string `json:"mobile"`
	}

	retVerifyRequest struct {
		CommonResponse
		Data VerifyRequestData `json:"data"`
	}
	VerifyRequestData struct {
		Ref string `json:"ref"`
	}
)

func (y *Yunzhanghu) VerifyRequest(ctx context.Context, req *ReqVerifyRequest) (*VerifyRequestData, error) {
	apiName := "银行卡四要素鉴权请求（下发短信验证码）"
	resp := new(retVerifyRequest)

	json, err := y.postJSON(ctx, VerifyRequestURI, apiName, req)
	if err != nil {
		return nil, err
	}

	err = y.decodeWithError(json, resp, apiName)
	if err != nil {
		return nil, err
	}
	return &resp.Data, err
}

/* 银行卡四要素确认鉴权（上传短信验证码） */

const VerifyConfirmURI = ""

type (
	ReqVerifyConfirm struct {
		CardNo   string `json:"card_no"`
		IdCard   string `json:"id_card"`
		RealName string `json:"real_name"`
		Mobile   string `json:"mobile"`
		Captcha  string `json:"captcha"`
		Ref      string `json:"ref"`
	}

	retVerifyConfirm struct {
		CommonResponse
	}
)

func (y *Yunzhanghu) VerifyConfirm(ctx context.Context, req *ReqVerifyConfirm) error {
	apiName := "银行卡四要素确认鉴权（上传短信验证码）"
	resp := new(retVerifyConfirm)

	json, err := y.postJSON(ctx, VerifyConfirmURI, apiName, req)
	if err != nil {
		return err
	}

	return y.decodeWithError(json, resp, apiName)
}

/* 银行卡四要素验证 */

const VerifyBankCardFourFactorURI = "/authentication/verify-bankcard-four-factor"

type (
	ReqVerifyBankCardFourFactor struct {
		CardNo   string `json:"card_no"`
		IdCard   string `json:"id_card"`
		RealName string `json:"real_name"`
		Mobile   string `json:"mobile"`
	}

	retVerifyBankCardFourFactor struct {
		CommonResponse
	}
)

func (y *Yunzhanghu) VerifyBankCardFourFactor(ctx context.Context, req *ReqVerifyBankCardFourFactor) error {
	apiName := "银行卡四要素验证"
	resp := new(retVerifyBankCardFourFactor)

	json, err := y.postJSON(ctx, VerifyBankCardFourFactorURI, apiName, req)
	if err != nil {
		return err
	}

	return y.decodeWithError(json, resp, apiName)
}

/* 银行卡三要素验证 */

const VerifyBankCardThreeFactorURI = "/authentication/verify-bankcard-three-factor"

type (
	ReqVerifyBankCardThreeFactor struct {
		CardNo   string `json:"card_no"`
		IdCard   string `json:"id_card"`
		RealName string `json:"real_name"`
	}

	retVerifyBankCardThreeFactor struct {
		CommonResponse
		Data VerifyBankCardThreeFactorData `json:"data"`
	}
	VerifyBankCardThreeFactorData struct {
	}
)

func (y *Yunzhanghu) VerifyBankCardThreeFactor(ctx context.Context, req *ReqVerifyBankCardThreeFactor) error {
	apiName := "银行卡三要素验证"
	resp := new(retVerifyBankCardThreeFactor)

	json, err := y.postJSON(ctx, VerifyBankCardThreeFactorURI, apiName, req)
	if err != nil {
		return err
	}

	return y.decodeWithError(json, resp, apiName)
}

/* 身份证实名验证 */

const VerifyIdCardURI = "/authentication/verify-id"

type (
	ReqVerifyIdCard struct {
		IdCard   string `json:"id_card"`
		RealName string `json:"real_name"`
	}

	retVerifyIdCard struct {
		CommonResponse
	}
)

func (y *Yunzhanghu) VerifyIdCard(ctx context.Context, req *ReqVerifyIdCard) error {
	apiName := "身份证实名验证"
	resp := new(retVerifyIdCard)

	json, err := y.postJSON(ctx, VerifyIdCardURI, apiName, req)
	if err != nil {
		return err
	}

	return y.decodeWithError(json, resp, apiName)
}

/* 查看用户免验证名单是否存在 */

const UserWhiteCheckURI = "/api/payment/v1/user/white/check"

type (
	ReqUserWhiteCheck struct {
		RealName string `json:"real_name"`
		IdCard   string `json:"id_card"`
	}

	retUserWhiteCheck struct {
		CommonResponse
		Data UserWhiteCheckData `json:"data"`
	}
	UserWhiteCheckData struct {
		Ok string `json:"ok"`
	}
)

func (y *Yunzhanghu) UserWhiteCheck(ctx context.Context, req *ReqUserWhiteCheck) (*UserWhiteCheckData, error) {
	apiName := "查看用户免验证名单是否存在"
	resp := new(retUserWhiteCheck)

	json, err := y.postJSON(ctx, UserWhiteCheckURI, apiName, req)
	if err != nil {
		return nil, err
	}

	err = y.decodeWithError(json, resp, apiName)
	if err != nil {
		return nil, err
	}
	return &resp.Data, err
}

/* 银行卡信息查询接口 */

const BankCardQueryURI = ""

type (
	ReqBankCardQuery struct {
		CardNo   string `json:"card_no"`
		BankName string `json:"bank_name"`
	}

	retBankCardQuery struct {
		CommonResponse
		Data BankCardQueryData `json:"data"`
	}
	BankCardQueryData struct {
		BankCode  string `json:"bank_code"`
		BankName  string `json:"bank_name"`
		CardType  string `json:"card_type"`
		IsSupport bool   `json:"is_support"`
	}
)

func (y *Yunzhanghu) BankCardQuery(ctx context.Context, req ReqBankCardQuery) (*BankCardQueryData, error) {
	apiName := "银行卡信息查询接口"
	resp := new(retBankCardQuery)

	json, err := y.getJson(ctx, BankCardQueryURI, apiName, req)
	if err != nil {
		return nil, err
	}

	err = y.decodeWithError(json, resp, apiName)
	if err != nil {
		return nil, err
	}
	return &resp.Data, err
}
