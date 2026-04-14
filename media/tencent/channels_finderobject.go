package tencent

import (
	"context"

	"github.com/cngamesdk/media-sdk/media/tencent/model"
)

// ChannelsFinderobjectGetSelf 视频号-获取视频号动态详情
// https://developers.e.qq.com/v3.0/docs/api/channels_finderobject/get
func (a *TencentAdapter) ChannelsFinderobjectGetSelf(ctx context.Context, req *model.ChannelsFinderobjectGetReq) (
	resp *model.ChannelsFinderobjectGetResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	var result model.ChannelsFinderobjectGetResp
	if requestErr := a.RequestGet(ctx, nil, model.ApiUrl3+"/channels_finderobject/get", req, &result); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}
