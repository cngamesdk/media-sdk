package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// AppShareCanShareAccountList 获取可共享的账号列表
// https://ad.e.kuaishou.com/rest/openapi/gw/dsp/appcenter/app/canShare/accountList
func (a *KuaishouAdapter) AppShareCanShareAccountList(ctx context.Context, req *kuaishouModel.AppShareCanShareAccountListReq) (resp *kuaishouModel.AppShareCanShareAccountListResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.AppShareCanShareAccountListResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/appcenter/app/canShare/accountList", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
