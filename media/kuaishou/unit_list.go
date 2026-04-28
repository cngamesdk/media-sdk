package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// UnitList 查询广告组
func (a *KuaishouAdapter) UnitList(ctx context.Context, req *kuaishouModel.UnitListReq) (resp *kuaishouModel.UnitListResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.UnitListResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/unit/list", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
