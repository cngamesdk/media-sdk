package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// MerchantDetailReport 小店通转化数据报表
func (a *KuaishouAdapter) MerchantDetailReport(ctx context.Context, req *kuaishouModel.MerchantDetailReportReq) (resp *kuaishouModel.MerchantDetailReportResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.MerchantDetailReportResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/v1/merchant/report/detail_report", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
