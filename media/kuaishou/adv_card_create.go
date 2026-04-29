package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// AdvCardCreate 创建高级创意
func (a *KuaishouAdapter) AdvCardCreate(ctx context.Context, req *kuaishouModel.AdvCardCreateReq) (resp *kuaishouModel.AdvCardCreateResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.AdvCardCreateResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/v1/asset/adv_card/create", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
