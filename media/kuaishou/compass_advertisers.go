package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// CompassAdvertisersSelf 获取罗盘绑定广告主列表
func (a *KuaishouAdapter) CompassAdvertisersSelf(ctx context.Context, req *kuaishouModel.CompassAdvertisersReq) (resp *kuaishouModel.CompassAdvertisersResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.CompassAdvertisersResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/uc/v1/advertisers", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
