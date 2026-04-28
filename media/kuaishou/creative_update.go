package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// CreativeUpdate 修改自定义创意
func (a *KuaishouAdapter) CreativeUpdate(ctx context.Context, req *kuaishouModel.CreativeUpdateReq) (resp *kuaishouModel.CreativeUpdateResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.CreativeUpdateResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/creative/update", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
