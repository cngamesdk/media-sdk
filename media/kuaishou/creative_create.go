package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// CreativeCreate 创建自定义创意
func (a *KuaishouAdapter) CreativeCreate(ctx context.Context, req *kuaishouModel.CreativeCreateReq) (resp *kuaishouModel.CreativeCreateResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.CreativeCreateResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/creative/create", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
