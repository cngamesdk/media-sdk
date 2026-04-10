package tencent

import (
	"context"

	"github.com/cngamesdk/media-sdk/media/tencent/model"
	"github.com/cngamesdk/media-sdk/utils"
)

// ProfileGetSelf 获取朋友圈头像昵称跳转页
// https://developers.e.qq.com/v3.0/docs/api/profiles/get
func (a *TencentAdapter) ProfileGetSelf(ctx context.Context, req *model.ProfileGetReq) (
	resp *model.ProfileGetResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	var result model.ProfileGetResp
	if requestErr := a.RequestGet(ctx, nil, model.ApiUrl3+"/profiles/get", req, &result); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}

// ProfileDeleteSelf 删除朋友圈头像昵称跳转页
// https://developers.e.qq.com/v3.0/docs/api/profiles/delete
// 注意：只能删除 profile_type 为 PROFILE_TYPE_DEFINITION 的跳转页
func (a *TencentAdapter) ProfileDeleteSelf(ctx context.Context, req *model.ProfileDeleteReq) (
	resp *model.ProfileDeleteResp, err error) {
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
	var result model.ProfileDeleteResp
	if requestErr := a.RequestPostJson(ctx, headers, model.ApiUrl3+"/profiles/delete?"+globalQuery, req, &result); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}
