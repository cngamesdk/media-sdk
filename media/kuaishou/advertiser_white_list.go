package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// AdvertiserWhiteList 获取创意分类标签白名单客户
func (a *KuaishouAdapter) AdvertiserWhiteList(ctx context.Context, req *kuaishouModel.AdvertiserWhiteListReq) (resp *kuaishouModel.AdvertiserWhiteListResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.AdvertiserWhiteListResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/v1/advertiser/white_list", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
