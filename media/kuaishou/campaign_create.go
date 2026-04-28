package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// CampaignCreate 创建广告计划
func (a *KuaishouAdapter) CampaignCreate(ctx context.Context, req *kuaishouModel.CampaignCreateReq) (resp *kuaishouModel.CampaignCreateResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.CampaignCreateResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/campaign/create", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
