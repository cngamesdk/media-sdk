package tencent

import (
	"context"

	"github.com/cngamesdk/media-sdk/media/tencent/model"
	"github.com/cngamesdk/media-sdk/utils"
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

// WechatPagesDeleteSelf 删除微信落地页
// https://developers.e.qq.com/v3.0/docs/api/wechat_pages/delete
func (a *TencentAdapter) WechatPagesDeleteSelf(ctx context.Context, req *model.WechatPagesDeleteReq) (
	resp *model.WechatPagesDeleteResp, err error) {
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
	var result model.WechatPagesDeleteResp
	if requestErr := a.RequestPostJson(ctx, headers, model.ApiUrl3+"/wechat_pages/delete?"+globalQuery, req, &result); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}

// WechatPagesAddSelf 基于模板创建微信原生页
// https://developers.e.qq.com/v3.0/docs/api/wechat_pages/add
func (a *TencentAdapter) WechatPagesAddSelf(ctx context.Context, req *model.WechatPagesAddReq) (
	resp *model.WechatPagesAddResp, err error) {
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
	var result model.WechatPagesAddResp
	if requestErr := a.RequestPostJson(ctx, headers, model.ApiUrl3+"/wechat_pages/add?"+globalQuery, req, &result); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}
