package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// AdvCardList 获取高级创意列表
func (a *KuaishouAdapter) AdvCardList(ctx context.Context, req *kuaishouModel.AdvCardListReq) (resp *kuaishouModel.AdvCardListResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.AdvCardListResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/v1/asset/adv_card/list", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
