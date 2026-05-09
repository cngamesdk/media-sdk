package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// DmpPopulationCreate 人群包创建(新)
// https://ad.e.kuaishou.com/rest/openapi/gw/dmp/v2/dmp/population/upload
func (a *KuaishouAdapter) DmpPopulationCreate(ctx context.Context, req *kuaishouModel.DmpPopulationCreateReq) (resp *kuaishouModel.DmpPopulationCreateResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.DmpPopulationCreateResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dmp/v2/dmp/population/upload", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
