package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// AdvancedCreativeUpdate 修改程序化创意
func (a *KuaishouAdapter) AdvancedCreativeUpdate(ctx context.Context, req *kuaishouModel.AdvancedCreativeUpdateReq) (resp *kuaishouModel.AdvancedCreativeUpdateResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.AdvancedCreativeUpdateResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/advanced_creative/update", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
