package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// CampaignList 查询广告计划
func (a *KuaishouAdapter) CampaignList(ctx context.Context, req *kuaishouModel.CampaignListReq) (resp *kuaishouModel.CampaignListResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.CampaignListResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/campaign/list", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
