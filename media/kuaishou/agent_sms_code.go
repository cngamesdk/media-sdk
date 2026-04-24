package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// AgentSmsCodeSendSelf 代理商-创建广告主发送验证码
func (a *KuaishouAdapter) AgentSmsCodeSendSelf(ctx context.Context, req *kuaishouModel.AgentSmsCodeSendReq) (resp *kuaishouModel.AgentSmsCodeSendResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.AgentSmsCodeSendResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/agent/v1/sms_code/send", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
