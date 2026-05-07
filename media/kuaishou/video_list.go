package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// VideoList 获取视频信息list
func (a *KuaishouAdapter) VideoList(ctx context.Context, req *kuaishouModel.VideoListReq) (resp *kuaishouModel.VideoListResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.VideoListResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/v1/file/ad/video/list", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
