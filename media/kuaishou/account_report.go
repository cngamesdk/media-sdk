package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// AccountReport 广告主数据（实时）
func (a *KuaishouAdapter) AccountReport(ctx context.Context, req *kuaishouModel.AccountReportReq) (resp *kuaishouModel.AccountReportResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.AccountReportResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/v1/report/account_report", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
