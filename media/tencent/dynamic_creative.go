package tencent

import (
	"context"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
	"github.com/cngamesdk/media-sdk/utils"
)

// DynamicCreativesGetSelf 获取创意
// https://developers.e.qq.com/v3.0/docs/api/dynamic_creatives/get
func (a *TencentAdapter) DynamicCreativesGetSelf(ctx context.Context, req *model.DynamicCreativesGetReq) (
	resp *model.DynamicCreativesGetResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	var result model.DynamicCreativesGetResp
	if requestErr := a.RequestGet(ctx, nil, model.ApiUrl3+"/dynamic_creatives/get", req, &result); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}

// DynamicCreativesAddSelf 创建创意
// https://developers.e.qq.com/v3.0/docs/api/dynamic_creatives/add
func (a *TencentAdapter) DynamicCreativesAddSelf(ctx context.Context, req *model.DynamicCreativesAddReq) (
	resp *model.DynamicCreativesAddResp, err error) {
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
	var result model.DynamicCreativesAddResp
	if requestErr := a.RequestPostJson(ctx, headers, model.ApiUrl3+"/dynamic_creatives/add?"+globalQuery, req, &result); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}
