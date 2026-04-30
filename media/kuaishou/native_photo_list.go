package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// NativePhotoList 获取达人原生视频列表
func (a *KuaishouAdapter) NativePhotoList(ctx context.Context, req *kuaishouModel.NativePhotoListReq) (resp *kuaishouModel.NativePhotoListResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.NativePhotoListResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/v1/native/photo/list", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
