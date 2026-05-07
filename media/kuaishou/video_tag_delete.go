package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// VideoTagDelete 删除视频标签
func (a *KuaishouAdapter) VideoTagDelete(ctx context.Context, req *kuaishouModel.VideoTagDeleteReq) (resp *kuaishouModel.VideoTagDeleteResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.VideoTagDeleteResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/v1/file/ad/video/tag/delete", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
