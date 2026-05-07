package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// ImagePush 图片推送
func (a *KuaishouAdapter) ImagePush(ctx context.Context, req *kuaishouModel.ImagePushReq) (resp *kuaishouModel.ImagePushResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.ImagePushResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/v1/pic/share", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
