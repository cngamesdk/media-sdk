package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// AgentCopyAccountsSelf 代理商-可复制账户列表
func (a *KuaishouAdapter) AgentCopyAccountsSelf(ctx context.Context, req *kuaishouModel.AgentCopyAccountsReq) (resp *kuaishouModel.AgentCopyAccountsResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.AgentCopyAccountsResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/agent/v1/bind/copy/accounts", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
