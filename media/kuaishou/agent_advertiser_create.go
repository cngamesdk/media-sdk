package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// AgentAdvertiserCreateSelf 代理商-创建广告主
func (a *KuaishouAdapter) AgentAdvertiserCreateSelf(ctx context.Context, req *kuaishouModel.AgentAdvertiserCreateReq) (resp *kuaishouModel.AgentAdvertiserCreateResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.AgentAdvertiserCreateResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/agent/v1/advertiser/create", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
