package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// MaterialReport 素材报表
func (a *KuaishouAdapter) MaterialReport(ctx context.Context, req *kuaishouModel.MaterialReportReq) (resp *kuaishouModel.MaterialReportResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.MaterialReportResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/v1/report/material_report", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
