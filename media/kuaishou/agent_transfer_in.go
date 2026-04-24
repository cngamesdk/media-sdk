package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// AgentTransferInSelf 广告主退钱给代理商
func (a *KuaishouAdapter) AgentTransferInSelf(ctx context.Context, req *kuaishouModel.AgentTransferInReq) (resp *kuaishouModel.AgentTransferInResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.AgentTransferInResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/agent/v2/finance/transfer/in", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
