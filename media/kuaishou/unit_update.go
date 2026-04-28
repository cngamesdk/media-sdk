package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// UnitUpdate 修改广告组
func (a *KuaishouAdapter) UnitUpdate(ctx context.Context, req *kuaishouModel.UnitUpdateReq) (resp *kuaishouModel.UnitUpdateResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.UnitUpdateResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/unit/update", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
