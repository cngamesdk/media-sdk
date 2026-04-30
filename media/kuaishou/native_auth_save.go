package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// NativeAuthSave 快手号授权
func (a *KuaishouAdapter) NativeAuthSave(ctx context.Context, req *kuaishouModel.NativeAuthSaveReq) (resp kuaishouModel.NativeAuthSaveResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.NativeAuthSaveResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/v1/native/auth/save", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = result
	return
}
