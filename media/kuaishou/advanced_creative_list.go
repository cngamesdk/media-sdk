package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// AdvancedCreativeList 查询程序化创意
func (a *KuaishouAdapter) AdvancedCreativeList(ctx context.Context, req *kuaishouModel.AdvancedCreativeListReq) (resp *kuaishouModel.AdvancedCreativeListResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.AdvancedCreativeListResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/advanced_creative/list", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
