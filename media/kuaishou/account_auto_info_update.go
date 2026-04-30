package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// AccountAutoInfoUpdate 更新账户智投(标准)配置信息
func (a *KuaishouAdapter) AccountAutoInfoUpdate(ctx context.Context, req *kuaishouModel.AccountAutoInfoUpdateReq) (resp *kuaishouModel.AccountAutoInfoUpdateResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.AccountAutoInfoUpdateResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/account/mod/auto/info", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
