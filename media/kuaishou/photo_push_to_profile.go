package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// PhotoPushToProfile 隐藏视频同步个人主页
func (a *KuaishouAdapter) PhotoPushToProfile(ctx context.Context, req *kuaishouModel.PhotoPushToProfileReq) (resp *kuaishouModel.PhotoPushToProfileResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.PhotoPushToProfileResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/v1/photo/pushToProfile", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
