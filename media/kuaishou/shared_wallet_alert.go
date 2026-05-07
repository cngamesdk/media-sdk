package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// SharedWalletAlert 客户共享钱包告警信息查询
func (a *KuaishouAdapter) SharedWalletAlert(ctx context.Context, req *kuaishouModel.SharedWalletAlertReq) (resp *kuaishouModel.SharedWalletAlertResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.SharedWalletAlertResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/trade/shared/alert", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
