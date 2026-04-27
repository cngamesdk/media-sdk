package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// AgentFinanceCanDecreaseQuerySelf 广告主可转/退金额查询
func (a *KuaishouAdapter) AgentFinanceCanDecreaseQuerySelf(ctx context.Context, req *kuaishouModel.AgentFinanceCanDecreaseQueryReq) (resp *kuaishouModel.AgentFinanceCanDecreaseQueryResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.AgentFinanceCanDecreaseQueryResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/agent/v1/finance/account/can/decrease/query", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
