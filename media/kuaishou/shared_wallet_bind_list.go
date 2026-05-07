package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// SharedWalletBindList 客户共享钱包绑定账户查询
func (a *KuaishouAdapter) SharedWalletBindList(ctx context.Context, req *kuaishouModel.SharedWalletBindListReq) (resp *kuaishouModel.SharedWalletBindListResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.SharedWalletBindListResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/trade/shared/bindList", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
