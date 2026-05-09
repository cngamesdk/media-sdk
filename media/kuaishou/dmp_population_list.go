package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// DmpPopulationList 人群列表查询
// https://ad.e.kuaishou.com/rest/openapi/v2/dmp/population/list
func (a *KuaishouAdapter) DmpPopulationList(ctx context.Context, req *kuaishouModel.DmpPopulationListReq) (resp *kuaishouModel.DmpPopulationListResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.DmpPopulationListResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/v2/dmp/population/list", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
