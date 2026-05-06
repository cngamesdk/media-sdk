package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// ExploreReport 探索工具关键报表数据
func (a *KuaishouAdapter) ExploreReport(ctx context.Context, req *kuaishouModel.ExploreReportReq) (resp *kuaishouModel.ExploreReportResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.ExploreReportResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/v1/unit/explore/report/detail", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
