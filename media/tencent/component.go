package tencent

import (
	"context"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
	"github.com/cngamesdk/media-sdk/utils"
)

// ComponentsGetSelf 获取创意组件
// https://developers.e.qq.com/v3.0/docs/api/components/get
func (a *TencentAdapter) ComponentsGetSelf(ctx context.Context, req *model.ComponentsGetReq) (
	resp *model.ComponentsGetResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	var result model.ComponentsGetResp
	if requestErr := a.RequestGet(ctx, nil, model.ApiUrl3+"/components/get", req, &result); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}

// ComponentsAddSelf 创建创意组件
// https://developers.e.qq.com/v3.0/docs/api/components/add
func (a *TencentAdapter) ComponentsAddSelf(ctx context.Context, req *model.ComponentsAddReq) (
	resp *model.ComponentsAddResp, err error) {
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
	var result model.ComponentsAddResp
	if requestErr := a.RequestPostJson(ctx, headers, model.ApiUrl3+"/components/add?"+globalQuery, req, &result); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}

// ComponentsDeleteSelf 删除创意组件
// https://developers.e.qq.com/v3.0/docs/api/components/delete
func (a *TencentAdapter) ComponentsDeleteSelf(ctx context.Context, req *model.ComponentsDeleteReq) (
	resp *model.ComponentsDeleteResp, err error) {
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
	var result model.ComponentsDeleteResp
	if requestErr := a.RequestPostJson(ctx, headers, model.ApiUrl3+"/components/delete?"+globalQuery, req, &result); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}

// ComponentSharingUpdateSelf 修改创意组件共享
// https://developers.e.qq.com/v3.0/docs/api/component_sharing/update
func (a *TencentAdapter) ComponentSharingUpdateSelf(ctx context.Context, req *model.ComponentSharingUpdateReq) (
	resp *model.ComponentSharingUpdateResp, err error) {
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
	var result model.ComponentSharingUpdateResp
	if requestErr := a.RequestPostJson(ctx, headers, model.ApiUrl3+"/component_sharing/update?"+globalQuery, req, &result); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}

// ComponentSharingGetSelf 查询创意组件共享信息
// https://developers.e.qq.com/v3.0/docs/api/component_sharing/get
func (a *TencentAdapter) ComponentSharingGetSelf(ctx context.Context, req *model.ComponentSharingGetReq) (
	resp *model.ComponentSharingGetResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	var result model.ComponentSharingGetResp
	if requestErr := a.RequestGet(ctx, nil, model.ApiUrl3+"/component_sharing/get", req, &result); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}

// ComponentDetailGetSelf 获取创意组件详情
// https://developers.e.qq.com/v3.0/docs/api/component_detail/get
func (a *TencentAdapter) ComponentDetailGetSelf(ctx context.Context, req *model.ComponentDetailGetReq) (
	resp *model.ComponentDetailGetResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	var result model.ComponentDetailGetResp
	if requestErr := a.RequestGet(ctx, nil, model.ApiUrl3+"/component_detail/get", req, &result); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}
