package tencent

import (
	"context"

	"github.com/cngamesdk/media-sdk/media/tencent/model"
)

// ChannelsUserpageobjectsGetSelf 视频号-获取视频号动态列表
// https://developers.e.qq.com/v3.0/docs/api/channels_userpageobjects/get
func (a *TencentAdapter) ChannelsUserpageobjectsGetSelf(ctx context.Context, req *model.ChannelsUserpageobjectsGetReq) (
	resp *model.ChannelsUserpageobjectsGetResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	var result model.ChannelsUserpageobjectsGetResp
	if requestErr := a.RequestGet(ctx, nil, model.ApiUrl3+"/channels_userpageobjects/get", req, &result); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}
