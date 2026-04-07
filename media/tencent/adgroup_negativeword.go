package tencent

import (
	"context"

	"github.com/cngamesdk/media-sdk/media/tencent/model"
	"github.com/cngamesdk/media-sdk/utils"
)

// AdgroupNegativewordAddSelf 新增广告否定词
// https://developers.e.qq.com/v3.0/docs/api/adgroup_negativewords/add
func (a *TencentAdapter) AdgroupNegativewordAddSelf(ctx context.Context, req *model.AdgroupNegativewordAddReq) (
	resp *model.AdgroupNegativewordAddResp, err error) {
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
	var result model.AdgroupNegativewordAddResp
	if requestErr := a.RequestPostJson(ctx, headers, model.ApiUrl3+"/adgroup_negativewords/add?"+globalQuery, req, &result); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}

// AdgroupNegativewordUpdateSelf 更新广告否定词
// https://developers.e.qq.com/v3.0/docs/api/adgroup_negativewords/update
func (a *TencentAdapter) AdgroupNegativewordUpdateSelf(ctx context.Context, req *model.AdgroupNegativewordUpdateReq) (
	resp *model.AdgroupNegativewordUpdateResp, err error) {
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
	var result model.AdgroupNegativewordUpdateResp
	if requestErr := a.RequestPostJson(ctx, headers, model.ApiUrl3+"/adgroup_negativewords/update?"+globalQuery, req, &result); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}

// AdgroupNegativewordGetSelf 查询广告否定词
// https://developers.e.qq.com/v3.0/docs/api/adgroup_negativewords/get
func (a *TencentAdapter) AdgroupNegativewordGetSelf(ctx context.Context, req *model.AdgroupNegativewordGetReq) (
	resp *model.AdgroupNegativewordGetResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	var result model.AdgroupNegativewordGetResp
	if requestErr := a.RequestGet(ctx, nil, model.ApiUrl3+"/adgroup_negativewords/get", req, &result); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}
