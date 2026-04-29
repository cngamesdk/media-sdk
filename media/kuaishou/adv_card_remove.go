package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// AdvCardRemove 删除高级创意
func (a *KuaishouAdapter) AdvCardRemove(ctx context.Context, req *kuaishouModel.AdvCardRemoveReq) (resp *kuaishouModel.AdvCardRemoveResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.AdvCardRemoveResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/v1/asset/adv_card/remove", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
