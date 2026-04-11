package tencent

import (
	"context"

	"github.com/cngamesdk/media-sdk/media/tencent/model"
)

// WechatPagesGetSelf 获取微信落地页列表
// https://developers.e.qq.com/v3.0/docs/api/wechat_pages/get
func (a *TencentAdapter) WechatPagesGetSelf(ctx context.Context, req *model.WechatPagesGetReq) (
	resp *model.WechatPagesGetResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	var result model.WechatPagesGetResp
	if requestErr := a.RequestGet(ctx, nil, model.ApiUrl3+"/wechat_pages/get", req, &result); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}
