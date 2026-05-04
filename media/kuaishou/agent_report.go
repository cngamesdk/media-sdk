package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// AgentReport 代理商数据（t-1 数据需要第二天 9 点以后获取）
func (a *KuaishouAdapter) AgentReport(ctx context.Context, req *kuaishouModel.AgentReportReq) (resp *kuaishouModel.AgentReportResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.AgentReportResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/v1/agent/report", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
