package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// UnitReport 广告组数据（不包括省心投物料）
func (a *KuaishouAdapter) UnitReport(ctx context.Context, req *kuaishouModel.UnitReportReq) (resp *kuaishouModel.UnitReportResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.UnitReportResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/v1/report/unit_report", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
