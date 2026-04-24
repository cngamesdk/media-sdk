package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// AgentAccountTransferSelf 广告主转账
func (a *KuaishouAdapter) AgentAccountTransferSelf(ctx context.Context, req *kuaishouModel.AgentAccountTransferReq) (resp *kuaishouModel.AgentAccountTransferResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.AgentAccountTransferResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/agent/v1/finance/account/transfer", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
