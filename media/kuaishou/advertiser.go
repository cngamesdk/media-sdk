package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// AdvertiserInfoSelf 获取广告主资质信息
func (a *KuaishouAdapter) AdvertiserInfoSelf(ctx context.Context, req *kuaishouModel.AdvertiserInfoReq) (resp *kuaishouModel.AdvertiserInfoResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.AdvertiserInfoResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/v1/advertiser/info", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
