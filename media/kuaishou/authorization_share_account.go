package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// AuthorizationShareAccount 获取已共享授权记录
func (a *KuaishouAdapter) AuthorizationShareAccount(ctx context.Context, req *kuaishouModel.AuthorizationShareAccountReq) (resp kuaishouModel.AuthorizationShareAccountResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.AuthorizationShareAccountResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/authorization/shareAccount/authorizeInfo", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = result
	return
}
