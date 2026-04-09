package tencent

import (
	"context"

	"github.com/cngamesdk/media-sdk/media/tencent/model"
)

// ImageGetSelf 获取图片信息
// https://developers.e.qq.com/v3.0/docs/api/images/get
func (a *TencentAdapter) ImageGetSelf(ctx context.Context, req *model.ImageGetReq) (
	resp *model.ImageGetResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	var result model.ImageGetResp
	if requestErr := a.RequestGet(ctx, nil, model.ApiUrl3+"/images/get", req, &result); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}
