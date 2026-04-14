package tencent

import (
	"context"

	"github.com/cngamesdk/media-sdk/media/tencent/model"
	"github.com/cngamesdk/media-sdk/utils"
)

// WechatChannelsAuthorizationUpdateSelf 视频号授权-更新视频号授权
// https://developers.e.qq.com/v3.0/docs/api/wechat_channels_authorization/update
func (a *TencentAdapter) WechatChannelsAuthorizationUpdateSelf(ctx context.Context, req *model.WechatChannelsAuthorizationUpdateReq) (
	resp *model.WechatChannelsAuthorizationUpdateResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	globalQuery, globalQueryErr := utils.ConvertStructToQueryString(req.GlobalReq)
	if globalQueryErr != nil {
		err = globalQueryErr
		return
	}
	req.GlobalReq.Clear()
	headers := make(model.Headers)
	headers.Json()
	var result model.WechatChannelsAuthorizationUpdateResp
	if requestErr := a.RequestPostJson(ctx, headers, model.ApiUrl3+"/wechat_channels_authorization/update?"+globalQuery, req, &result); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}
