package tencent

import (
	"context"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
	"github.com/cngamesdk/media-sdk/utils"
)

// AdgroupsGetSelf 获取广告
// https://developers.e.qq.com/v3.0/docs/api/adgroups/get
func (a *TencentAdapter) AdgroupsGetSelf(ctx context.Context, req *model.AdgroupsGetReq) (
	resp *model.AdgroupsGetResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	var result model.AdgroupsGetResp
	if requestErr := a.RequestGet(ctx, nil, model.ApiUrl3+"/adgroups/get", req, &result); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}

// AdgroupsAddSelf 创建广告
// https://developers.e.qq.com/v3.0/docs/api/adgroups/add
func (a *TencentAdapter) AdgroupsAddSelf(ctx context.Context, req *model.AdgroupsAddReq) (
	resp *model.AdgroupsAddResp, err error) {
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
	var result model.AdgroupsAddResp
	if requestErr := a.RequestPostJson(ctx, headers, model.ApiUrl3+"/adgroups/add?"+globalQuery, req, &result); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}

// AdgroupsDeleteSelf 删除广告
// https://developers.e.qq.com/v3.0/docs/api/adgroups/delete
func (a *TencentAdapter) AdgroupsDeleteSelf(ctx context.Context, req *model.AdgroupsDeleteReq) (
	resp *model.AdgroupsDeleteResp, err error) {
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
	var result model.AdgroupsDeleteResp
	if requestErr := a.RequestPostJson(ctx, headers, model.ApiUrl3+"/adgroups/delete?"+globalQuery, req, &result); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}

// AdgroupsUpdateSelf 更新广告
// https://developers.e.qq.com/v3.0/docs/api/adgroups/update
func (a *TencentAdapter) AdgroupsUpdateSelf(ctx context.Context, req *model.AdgroupsUpdateReq) (
	resp *model.AdgroupsUpdateResp, err error) {
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
	var result model.AdgroupsUpdateResp
	if requestErr := a.RequestPostJson(ctx, headers, model.ApiUrl3+"/adgroups/update?"+globalQuery, req, &result); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}
