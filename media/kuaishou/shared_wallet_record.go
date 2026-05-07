package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// SharedWalletRecord 客户共享钱包交易明细查询
func (a *KuaishouAdapter) SharedWalletRecord(ctx context.Context, req *kuaishouModel.SharedWalletRecordReq) (resp *kuaishouModel.SharedWalletRecordResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.SharedWalletRecordResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/trade/shared/record", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
