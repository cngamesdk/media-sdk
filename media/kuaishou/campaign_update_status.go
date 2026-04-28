package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// CampaignUpdateStatus 修改广告计划状态
func (a *KuaishouAdapter) CampaignUpdateStatus(ctx context.Context, req *kuaishouModel.CampaignUpdateStatusReq) (resp *kuaishouModel.CampaignUpdateStatusResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.CampaignUpdateStatusResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/v1/campaign/update/status", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
