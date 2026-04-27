package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// AgentCheckAndCopySelf 代理商-复制账户
func (a *KuaishouAdapter) AgentCheckAndCopySelf(ctx context.Context, req *kuaishouModel.AgentCheckAndCopyReq) (resp kuaishouModel.AgentCheckAndCopyResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.AgentCheckAndCopyResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/agent/v1/bind/copy/checkAndCopy", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = result
	return
}
