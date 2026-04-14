package tencent

import (
	"context"

	"github.com/cngamesdk/media-sdk/media/tencent/model"
)

// WechatChannelsAuthorizationGetSelf 视频号授权-获取视频号授权记录列表
// https://developers.e.qq.com/v3.0/docs/api/wechat_channels_authorization/get
func (a *TencentAdapter) WechatChannelsAuthorizationGetSelf(ctx context.Context, req *model.WechatChannelsAuthorizationGetReq) (
	resp *model.WechatChannelsAuthorizationGetResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	var result model.WechatChannelsAuthorizationGetResp
	if requestErr := a.RequestGet(ctx, nil, model.ApiUrl3+"/wechat_channels_authorization/get", req, &result); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}
