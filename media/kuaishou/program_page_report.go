package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// ProgramPageReport 程序化落地页报表
func (a *KuaishouAdapter) ProgramPageReport(ctx context.Context, req *kuaishouModel.ProgramPageReportReq) (resp *kuaishouModel.ProgramPageReportResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.ProgramPageReportResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/v1/page/report/detail", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
