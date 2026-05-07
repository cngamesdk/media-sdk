package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// SharedWalletList 客户共享钱包列表
func (a *KuaishouAdapter) SharedWalletList(ctx context.Context, req *kuaishouModel.SharedWalletListReq) (resp *kuaishouModel.SharedWalletListResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.SharedWalletListResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/trade/shared/walletList", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
