package tencent

import (
	"context"

	"github.com/cngamesdk/media-sdk/media/tencent/model"
)

// ProfileGetSelf 获取朋友圈头像昵称跳转页
// https://developers.e.qq.com/v3.0/docs/api/profiles/get
func (a *TencentAdapter) ProfileGetSelf(ctx context.Context, req *model.ProfileGetReq) (
	resp *model.ProfileGetResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	var result model.ProfileGetResp
	if requestErr := a.RequestGet(ctx, nil, model.ApiUrl3+"/profiles/get", req, &result); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}
