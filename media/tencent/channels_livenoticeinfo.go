package tencent

import (
	"context"

	"github.com/cngamesdk/media-sdk/media/tencent/model"
)

// ChannelsLivenoticeinfoGetSelf 视频号-获取视频号当前的预约直播信息
// https://developers.e.qq.com/v3.0/docs/api/channels_livenoticeinfo/get
func (a *TencentAdapter) ChannelsLivenoticeinfoGetSelf(ctx context.Context, req *model.ChannelsLivenoticeinfoGetReq) (
	resp *model.ChannelsLivenoticeinfoGetResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	var result model.ChannelsLivenoticeinfoGetResp
	if requestErr := a.RequestGet(ctx, nil, model.ApiUrl3+"/channels_livenoticeinfo/get", req, &result); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}
