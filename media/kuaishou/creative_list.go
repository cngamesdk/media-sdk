package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// CreativeList 查询自定义创意
func (a *KuaishouAdapter) CreativeList(ctx context.Context, req *kuaishouModel.CreativeListReq) (resp *kuaishouModel.CreativeListResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.CreativeListResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/creative/list", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
