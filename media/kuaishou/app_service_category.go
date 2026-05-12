package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// AppServiceCategory 获取APP服务类目详情
// https://ad.e.kuaishou.com/rest/openapi/gw/dsp/appcenter/app/service/category
func (a *KuaishouAdapter) AppServiceCategory(ctx context.Context, req *kuaishouModel.AppServiceCategoryReq) (resp *kuaishouModel.AppServiceCategoryResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.AppServiceCategoryResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/appcenter/app/service/category", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
