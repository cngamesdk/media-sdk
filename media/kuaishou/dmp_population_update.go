package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// DmpPopulationUpdate 人群包更新(新)
// https://ad.e.kuaishou.com/rest/openapi/gw/dmp/v2/dmp/population/update
func (a *KuaishouAdapter) DmpPopulationUpdate(ctx context.Context, req *kuaishouModel.DmpPopulationUpdateReq) (resp *kuaishouModel.DmpPopulationUpdateResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.DmpPopulationUpdateResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dmp/v2/dmp/population/update", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
