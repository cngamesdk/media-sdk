package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// SharedWalletTransfer 客户共享钱包间转账
func (a *KuaishouAdapter) SharedWalletTransfer(ctx context.Context, req *kuaishouModel.SharedWalletTransferReq) (resp *kuaishouModel.SharedWalletTransferResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.SharedWalletTransferResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/trade/shared/wallet/transfer", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
