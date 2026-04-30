package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// NativeOpenAccount 开启原生扩量开关
func (a *KuaishouAdapter) NativeOpenAccount(ctx context.Context, req *kuaishouModel.NativeOpenAccountReq) (resp *kuaishouModel.NativeOpenAccountResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.NativeOpenAccountResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/native/openAccountNative", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
