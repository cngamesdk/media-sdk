package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// VideoBatchUpdate 批量更新视频
func (a *KuaishouAdapter) VideoBatchUpdate(ctx context.Context, req *kuaishouModel.VideoBatchUpdateReq) (resp *kuaishouModel.VideoBatchUpdateResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.VideoBatchUpdateResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/v1/file/ad/video/update", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
