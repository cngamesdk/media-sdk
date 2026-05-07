package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// VideoCursorList 游标查询视频信息
func (a *KuaishouAdapter) VideoCursorList(ctx context.Context, req *kuaishouModel.VideoCursorListReq) (resp *kuaishouModel.VideoCursorListResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.VideoCursorListResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/video/listByCursor", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
