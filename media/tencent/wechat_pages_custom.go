package tencent

import (
	"context"

	"github.com/cngamesdk/media-sdk/media/tencent/model"
	"github.com/cngamesdk/media-sdk/utils"
)

// WechatPagesCustomAddSelf 基于组件创建微信原生页
// https://developers.e.qq.com/v3.0/docs/api/wechat_pages_custom/add
func (a *TencentAdapter) WechatPagesCustomAddSelf(ctx context.Context, req *model.WechatPagesCustomAddReq) (
	resp *model.WechatPagesCustomAddResp, err error) {
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
	var result model.WechatPagesCustomAddResp
	if requestErr := a.RequestPostJson(ctx, headers, model.ApiUrl3+"/wechat_pages_custom/add?"+globalQuery, req, &result); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}
