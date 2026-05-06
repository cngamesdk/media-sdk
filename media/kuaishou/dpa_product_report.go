package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// DpaProductReport 商品库报表
func (a *KuaishouAdapter) DpaProductReport(ctx context.Context, req *kuaishouModel.DpaProductReportReq) (resp *kuaishouModel.DpaProductReportResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.DpaProductReportResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/v1/dpa/product/report", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
