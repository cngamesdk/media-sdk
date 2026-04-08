package tencent

import (
	"context"

	"github.com/cngamesdk/media-sdk/media/tencent/model"
)

// VideoGetSelf 获取视频文件
// https://developers.e.qq.com/v3.0/docs/api/videos/get
func (a *TencentAdapter) VideoGetSelf(ctx context.Context, req *model.VideoGetReq) (
	resp *model.VideoGetResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	var result model.VideoGetResp
	if requestErr := a.RequestGet(ctx, nil, model.ApiUrl3+"/videos/get", req, &result); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}
