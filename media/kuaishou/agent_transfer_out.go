package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// AgentTransferOutSelf 代理商给广告主账户转账
func (a *KuaishouAdapter) AgentTransferOutSelf(ctx context.Context, req *kuaishouModel.AgentTransferOutReq) (resp *kuaishouModel.AgentTransferOutResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.AgentTransferOutResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/agent/v1/finance/transfer/out", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
