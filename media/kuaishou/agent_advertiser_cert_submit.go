package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// AgentAdvertiserCertSubmitSelf 创建或更新账户信息和开户资质
func (a *KuaishouAdapter) AgentAdvertiserCertSubmitSelf(ctx context.Context, req *kuaishouModel.AgentAdvertiserCertSubmitReq) (resp *kuaishouModel.AgentAdvertiserCertSubmitResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.AgentAdvertiserCertSubmitResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/v1/agent/advertiser/cert/submit", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
