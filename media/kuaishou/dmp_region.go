package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// DmpRegion 人群包数据
// https://ad.e.kuaishou.com/rest/openapi/gw/dsp/dmp/region
func (a *KuaishouAdapter) DmpRegion(ctx context.Context, req *kuaishouModel.DmpRegionReq) (resp *kuaishouModel.DmpRegionResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.DmpRegionResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/dmp/region", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
