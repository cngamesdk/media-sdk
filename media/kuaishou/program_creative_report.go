package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// ProgramCreativeReport 程序化创意数据（包含程序化2.0+智能创意，不包括省心投物料）
func (a *KuaishouAdapter) ProgramCreativeReport(ctx context.Context, req *kuaishouModel.ProgramCreativeReportReq) (resp *kuaishouModel.ProgramCreativeReportResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.ProgramCreativeReportResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/v1/report/program_creative_report", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
