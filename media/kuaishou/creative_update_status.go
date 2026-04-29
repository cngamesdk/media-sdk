package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// CreativeUpdateStatus 修改创意状态
func (a *KuaishouAdapter) CreativeUpdateStatus(ctx context.Context, req *kuaishouModel.CreativeUpdateStatusReq) (resp *kuaishouModel.CreativeUpdateStatusResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.CreativeUpdateStatusResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/v1/creative/update/status", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
