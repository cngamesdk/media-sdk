package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// AdvancedCreativeCreate 创建程序化创意
func (a *KuaishouAdapter) AdvancedCreativeCreate(ctx context.Context, req *kuaishouModel.AdvancedCreativeCreateReq) (resp *kuaishouModel.AdvancedCreativeCreateResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.AdvancedCreativeCreateResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/advanced_creative/create", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
