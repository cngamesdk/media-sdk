package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// VideoBatchDelete 批量删除视频
func (a *KuaishouAdapter) VideoBatchDelete(ctx context.Context, req *kuaishouModel.VideoBatchDeleteReq) (resp *kuaishouModel.VideoBatchDeleteResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.VideoBatchDeleteResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/file/ad/video/delete", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
