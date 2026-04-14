package tencent

import (
	"context"

	"github.com/cngamesdk/media-sdk/media/tencent/model"
	"github.com/cngamesdk/media-sdk/utils"
)

// WechatChannelsAuthorizationAddSelf 视频号授权-创建视频号授权
// https://developers.e.qq.com/v3.0/docs/api/wechat_channels_authorization/add
func (a *TencentAdapter) WechatChannelsAuthorizationAddSelf(ctx context.Context, req *model.WechatChannelsAuthorizationAddReq) (
	resp *model.WechatChannelsAuthorizationAddResp, err error) {
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
	var result model.WechatChannelsAuthorizationAddResp
	if requestErr := a.RequestPostJson(ctx, headers, model.ApiUrl3+"/wechat_channels_authorization/add?"+globalQuery, req, &result); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}
