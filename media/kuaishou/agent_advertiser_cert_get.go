package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// AgentAdvertiserCertGetSelf 获取账号信息和开户资质
func (a *KuaishouAdapter) AgentAdvertiserCertGetSelf(ctx context.Context, req *kuaishouModel.AgentAdvertiserCertGetReq) (resp *kuaishouModel.AgentAdvertiserCertGetResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.AgentAdvertiserCertGetResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/v1/agent/advertiser/cert/get", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
