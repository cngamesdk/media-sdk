package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// AgentAdvertiserModifyNameSelf 代理商-修改账户名称
func (a *KuaishouAdapter) AgentAdvertiserModifyNameSelf(ctx context.Context, req *kuaishouModel.AgentAdvertiserModifyNameReq) (resp *kuaishouModel.AgentAdvertiserModifyNameResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.AgentAdvertiserModifyNameResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/agent/v1/advertiser/modify/name", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
