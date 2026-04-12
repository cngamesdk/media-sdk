package tencent

import (
	"context"

	"github.com/cngamesdk/media-sdk/media/tencent/model"
	"github.com/cngamesdk/media-sdk/utils"
)

// XijingPageAddSelf 蹊径基于模板创建落地页
// https://developers.e.qq.com/v3.0/docs/api/xijing_page/add
func (a *TencentAdapter) XijingPageAddSelf(ctx context.Context, req *model.XijingPageAddReq) (
	resp *model.XijingPageAddResp, err error) {
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
	var result model.XijingPageAddResp
	if requestErr := a.RequestPostJson(ctx, headers, model.ApiUrl3+"/xijing_page/add?"+globalQuery, req, &result); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}

// XijingPageUpdateSelf 蹊径送审落地页
// https://developers.e.qq.com/v3.0/docs/api/xijing_page/update
func (a *TencentAdapter) XijingPageUpdateSelf(ctx context.Context, req *model.XijingPageUpdateReq) (
	resp *model.XijingPageUpdateResp, err error) {
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
	var result model.XijingPageUpdateResp
	if requestErr := a.RequestPostJson(ctx, headers, model.ApiUrl3+"/xijing_page/update?"+globalQuery, req, &result); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}

// XijingPageDeleteSelf 蹊径删除落地页
// https://developers.e.qq.com/v3.0/docs/api/xijing_page/delete
func (a *TencentAdapter) XijingPageDeleteSelf(ctx context.Context, req *model.XijingPageDeleteReq) (
	resp *model.XijingPageDeleteResp, err error) {
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
	var result model.XijingPageDeleteResp
	if requestErr := a.RequestPostJson(ctx, headers, model.ApiUrl3+"/xijing_page/delete?"+globalQuery, req, &result); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}
