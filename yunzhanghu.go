package yunzhanghu

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const (
	defaultApiAddr = "https://api-service.yunzhanghu.com"
)

var (
	random *rand.Rand
)

type Yunzhanghu struct {
	DesKey  string
	Appkey  string
	Dealer  string
	Broker  string
	ApiAddr string
	WxAppID string // 商户微信 AppID

	//白名单IP
	//120.55.214.118 、 116.62.0.220 、 118.31.31.71 、 120.55.214.49 、 49.4.23.21 、 117.78.48.61 、
	//39.98.185.242 、 39.98.236.201
	PayNotifyUrl      string
	PayNotifyCallback func(context.Context, *PaymentNotify) bool

	H5SignNotifyUrl      string
	H5SignNotifyCallback func(context.Context, *H5SignNotify) bool
	H5SignRedirectURL    string
}

func init() {
	random = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func NewDefault() *Yunzhanghu {
	return &Yunzhanghu{
		ApiAddr:         defaultApiAddr,
		PayNotifyUrl:    "http://127.0.0.1:8000/pay-notify",
		H5SignNotifyUrl: "http://127.0.0.1:8000/sign-notify",
		PayNotifyCallback: func(ctx context.Context, notify *PaymentNotify) bool {
			fmt.Fprintf(os.Stdout, "no handler callback: %#v\n", notify)
			return true
		},
		H5SignNotifyCallback: func(ctx context.Context, notify *H5SignNotify) bool {
			fmt.Fprintf(os.Stdout, "no handler callback: %#v\n", notify)
			return true
		},
	}
}
