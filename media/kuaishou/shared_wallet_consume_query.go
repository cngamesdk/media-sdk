package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// SharedWalletConsumeQuery 客户共享钱包消耗明细查询
func (a *KuaishouAdapter) SharedWalletConsumeQuery(ctx context.Context, req *kuaishouModel.SharedWalletConsumeQueryReq) (resp *kuaishouModel.SharedWalletConsumeQueryResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.SharedWalletConsumeQueryResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/trade/shared/consumeQuery", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
