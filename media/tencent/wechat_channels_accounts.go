package tencent

import (
	"context"

	"github.com/cngamesdk/media-sdk/media/tencent/model"
)

// WechatChannelsAccountsGetSelf 视频号-获取视频号列表
// https://developers.e.qq.com/v3.0/docs/api/wechat_channels_accounts/get
func (a *TencentAdapter) WechatChannelsAccountsGetSelf(ctx context.Context, req *model.WechatChannelsAccountsGetReq) (
	resp *model.WechatChannelsAccountsGetResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	var result model.WechatChannelsAccountsGetResp
	if requestErr := a.RequestGet(ctx, nil, model.ApiUrl3+"/wechat_channels_accounts/get", req, &result); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}
