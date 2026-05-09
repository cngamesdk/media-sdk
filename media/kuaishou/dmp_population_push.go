package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// DmpPopulationPush 人群包跨账户推送
// https://ad.e.kuaishou.com/rest/openapi/v1/dmp/population/accounts/push
func (a *KuaishouAdapter) DmpPopulationPush(ctx context.Context, req *kuaishouModel.DmpPopulationPushReq) (resp *kuaishouModel.DmpPopulationPushResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.DmpPopulationPushResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/v1/dmp/population/accounts/push", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
