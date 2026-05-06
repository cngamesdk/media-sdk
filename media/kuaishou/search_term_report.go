package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// SearchTermReport 搜索词报表
func (a *KuaishouAdapter) SearchTermReport(ctx context.Context, req *kuaishouModel.SearchTermReportReq) (resp *kuaishouModel.SearchTermReportResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.SearchTermReportResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/v1/report/query_word_report", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
