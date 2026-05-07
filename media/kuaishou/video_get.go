package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// VideoGet 获取视频信息
func (a *KuaishouAdapter) VideoGet(ctx context.Context, req *kuaishouModel.VideoGetReq) (resp *kuaishouModel.VideoGetResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.VideoGetResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/v1/file/ad/video/get", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
