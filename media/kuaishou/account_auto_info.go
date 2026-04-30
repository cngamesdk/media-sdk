package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// AccountAutoInfo 查询账户智投(标准)配置信息
func (a *KuaishouAdapter) AccountAutoInfo(ctx context.Context, req *kuaishouModel.AccountAutoInfoReq) (resp *kuaishouModel.AccountAutoInfoResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.AccountAutoInfoResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/account/query/auto/info", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
