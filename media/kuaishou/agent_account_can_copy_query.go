package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// AgentAccountCanCopyQuerySelf 拉取账户可复制信息列表
func (a *KuaishouAdapter) AgentAccountCanCopyQuerySelf(ctx context.Context, req *kuaishouModel.AgentAccountCanCopyQueryReq) (resp kuaishouModel.AgentAccountCanCopyQueryResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/agent/v1/account/can/copy/query", req, &resp); errRequest != nil {
		err = errRequest
		return
	}
	return
}
