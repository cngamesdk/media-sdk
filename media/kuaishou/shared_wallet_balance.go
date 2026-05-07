package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// SharedWalletBalance 客户共享钱包余额查询
func (a *KuaishouAdapter) SharedWalletBalance(ctx context.Context, req *kuaishouModel.SharedWalletBalanceReq) (resp *kuaishouModel.SharedWalletBalanceResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.SharedWalletBalanceResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/trade/shared/balance", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
