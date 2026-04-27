package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// AgentPayTokenSelf 代理商-获取新的交易号
func (a *KuaishouAdapter) AgentPayTokenSelf(ctx context.Context, req *kuaishouModel.AgentPayTokenReq) (resp *kuaishouModel.AgentPayTokenResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.AgentPayTokenResp
	if errRequest := a.RequestPostForm(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/agent/v1/finance/pay/payToken", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
