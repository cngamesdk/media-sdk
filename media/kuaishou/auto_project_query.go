package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// AutoProjectQuery 智投项目查询
func (a *KuaishouAdapter) AutoProjectQuery(ctx context.Context, req *kuaishouModel.AutoProjectQueryReq) (resp *kuaishouModel.AutoProjectQueryResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.AutoProjectQueryResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/simple/project/query", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
