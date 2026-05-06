package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// KeywordReport 关键词报表
func (a *KuaishouAdapter) KeywordReport(ctx context.Context, req *kuaishouModel.KeywordReportReq) (resp *kuaishouModel.KeywordReportResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.KeywordReportResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/v1/report/word_info_report", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
