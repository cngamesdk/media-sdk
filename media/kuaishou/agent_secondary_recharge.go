package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// AgentSecondaryRechargeSelf 一代给二代充值
func (a *KuaishouAdapter) AgentSecondaryRechargeSelf(ctx context.Context, req *kuaishouModel.AgentSecondaryRechargeReq) (resp *kuaishouModel.AgentSecondaryRechargeResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.AgentSecondaryRechargeResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/agent/v1/finance/transfer/secondary/recharge", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
