package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// AgentFetchAccountListSelf 批量拉取代理商下账户列表
func (a *KuaishouAdapter) AgentFetchAccountListSelf(ctx context.Context, req *kuaishouModel.AgentFetchAccountListReq) (resp *kuaishouModel.AgentFetchAccountListResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.AgentFetchAccountListResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/agent/v1/fetch/account/list", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
