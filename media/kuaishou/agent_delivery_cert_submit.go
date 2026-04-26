package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// AgentDeliveryCertSubmitSelf 创建或更新投放资质
func (a *KuaishouAdapter) AgentDeliveryCertSubmitSelf(ctx context.Context, req *kuaishouModel.AgentDeliveryCertSubmitReq) (resp *kuaishouModel.AgentDeliveryCertSubmitResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.AgentDeliveryCertSubmitResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/v1/agent/advertiser/delivery_cert/submit", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
