package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// AgentAccountListSelf 代理商-账户列表
func (a *KuaishouAdapter) AgentAccountListSelf(ctx context.Context, req *kuaishouModel.AgentAccountListReq) (resp *kuaishouModel.AgentAccountListResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.AgentAccountListResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/agent/v1/account/list", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
