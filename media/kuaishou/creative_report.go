package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// CreativeReport 广告创意数据（自定义，不包括省心投物料）
func (a *KuaishouAdapter) CreativeReport(ctx context.Context, req *kuaishouModel.CreativeReportReq) (resp *kuaishouModel.CreativeReportResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.CreativeReportResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/v1/report/creative_report", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
