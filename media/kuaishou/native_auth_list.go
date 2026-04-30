package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// NativeAuthList 获取快手号授权列表
func (a *KuaishouAdapter) NativeAuthList(ctx context.Context, req *kuaishouModel.NativeAuthListReq) (resp *kuaishouModel.NativeAuthListResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.NativeAuthListResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/v1/native/auth/list", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
