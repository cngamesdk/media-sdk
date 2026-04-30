package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// AuaxSeriesList 智投短剧查询
func (a *KuaishouAdapter) AuaxSeriesList(ctx context.Context, req *kuaishouModel.AuaxSeriesListReq) (resp *kuaishouModel.AuaxSeriesListResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.AuaxSeriesListResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/v2/auax/series/list", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
