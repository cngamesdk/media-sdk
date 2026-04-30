package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// NativeUserList 获取原生快手号列表
func (a *KuaishouAdapter) NativeUserList(ctx context.Context, req *kuaishouModel.NativeUserListReq) (resp *kuaishouModel.NativeUserListResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.NativeUserListResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/v1/native/user/list", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
