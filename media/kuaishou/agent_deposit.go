package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// AgentDepositListSelf 代理商流水列表
func (a *KuaishouAdapter) AgentDepositListSelf(ctx context.Context, req *kuaishouModel.AgentDepositListReq) (resp *kuaishouModel.AgentDepositListResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.AgentDepositListResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/agent/v1/finance/deposit/agent", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
