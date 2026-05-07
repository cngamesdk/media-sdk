package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// SharedWalletBindOpt 客户共享钱包账户绑定/解绑
func (a *KuaishouAdapter) SharedWalletBindOpt(ctx context.Context, req *kuaishouModel.SharedWalletBindOptReq) (resp *kuaishouModel.SharedWalletBindOptResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.SharedWalletBindOptResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/trade/shared/bindopt", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
