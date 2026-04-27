package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// AgentAdvertiserDepositSelf 代理商-广告主流水列表
func (a *KuaishouAdapter) AgentAdvertiserDepositSelf(ctx context.Context, req *kuaishouModel.AgentAdvertiserDepositReq) (resp *kuaishouModel.AgentAdvertiserDepositResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.AgentAdvertiserDepositResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/agent/v1/finance/deposit/account", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
