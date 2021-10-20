package yunzhanghu

import (
	"context"
	"testing"
	
)

func TestYunzhanghu_UserCardCheck(t *testing.T) {
	type args struct {
		ctx      context.Context
		realName string
		idCard   string
	}

	tests := []struct {
		name    string
		client  *Yunzhanghu
		args    args
		wantErr bool
	}{
		{
			name:   "t1",
			client: &Yunzhanghu{},
			args: args{
				ctx: context.Background(),
				// TODO
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		//t.Run(tt.name, func(t *testing.T) {
		if err := tt.client.VerifyIdCard(tt.args.ctx, &ReqVerifyIdCard{
			IdCard:   tt.args.idCard,
			RealName: tt.args.realName,
		} , ); err != nil {
			t.Errorf("UserCardCheck() error = %v, wantErr %v", err, tt.wantErr)
		}
		//})
	}
}

func TestYunzhanghu_VerifyBankcardThreeFactor(t *testing.T) {
	type args struct {
		ctx      context.Context
		cardNo   string
		idCard   string
		realName string
	}
	tests := []struct {
		name    string
		client  *Yunzhanghu
		args    args
		wantErr bool
	}{
		{
			name:   "ok",
			client: &Yunzhanghu{},
			args: args{
				ctx: context.Background(),
				// TODO
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.client.VerifyBankCardThreeFactor(tt.args.ctx, &ReqVerifyBankCardThreeFactor{
				CardNo:   tt.args.cardNo,
				IdCard:   tt.args.idCard,
				RealName: tt.args.realName,
			}); (err != nil) != tt.wantErr {
				t.Errorf("VerifyBankcardThreeFactor() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestYunzhanghu_VerifyBankcardFourFactor(t *testing.T) {
	type args struct {
		ctx      context.Context
		cardNo   string
		idCard   string
		realName string
		mobile   string
	}
	tests := []struct {
		name    string
		client  *Yunzhanghu
		args    args
		wantErr bool
	}{
		{
			name:   "ok",
			client: &Yunzhanghu{},
			args: args{
				ctx: context.Background(),
				// TODO
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.client.VerifyBankCardFourFactor(tt.args.ctx, &ReqVerifyBankCardFourFactor{
				CardNo:   tt.args.cardNo,
				IdCard:   tt.args.idCard,
				RealName: tt.args.realName,
				Mobile:   tt.args.mobile,
			}); (err != nil) != tt.wantErr {
				t.Errorf("VerifyBankcardFourFactor() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
