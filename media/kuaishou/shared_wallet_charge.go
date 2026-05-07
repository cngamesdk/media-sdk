package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// SharedWalletCharge 客户共享钱包账户充值/退款
func (a *KuaishouAdapter) SharedWalletCharge(ctx context.Context, req *kuaishouModel.SharedWalletChargeReq) (resp *kuaishouModel.SharedWalletChargeResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.SharedWalletChargeResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/trade/shared/transfer", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
