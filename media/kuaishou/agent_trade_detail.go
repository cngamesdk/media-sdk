package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// AgentTradeDetailSelf 获取交易状态信息
func (a *KuaishouAdapter) AgentTradeDetailSelf(ctx context.Context, req *kuaishouModel.AgentTradeDetailReq) (resp *kuaishouModel.AgentTradeDetailResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.AgentTradeDetailResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/agent/v1/finance/trade/detail", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
