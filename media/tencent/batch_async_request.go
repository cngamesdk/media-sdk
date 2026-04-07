package tencent

import (
	"context"

	"github.com/cngamesdk/media-sdk/media/tencent/model"
	"github.com/cngamesdk/media-sdk/utils"
)

// BatchAsyncRequestSpecGetSelf 获取批量异步请求任务详情
// https://developers.e.qq.com/v3.0/docs/api/batch_async_request_specification/get
func (a *TencentAdapter) BatchAsyncRequestSpecGetSelf(ctx context.Context, req *model.BatchAsyncRequestSpecGetReq) (
	resp *model.BatchAsyncRequestSpecGetResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	var result model.BatchAsyncRequestSpecGetResp
	if requestErr := a.RequestGet(ctx, nil, model.ApiUrl3+"/batch_async_request_specification/get", req, &result); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}

// BatchAsyncRequestGetSelf 获取批量异步请求任务列表
// https://developers.e.qq.com/v3.0/docs/api/batch_async_requests/get
func (a *TencentAdapter) BatchAsyncRequestGetSelf(ctx context.Context, req *model.BatchAsyncRequestGetReq) (
	resp *model.BatchAsyncRequestGetResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	var result model.BatchAsyncRequestGetResp
	if requestErr := a.RequestGet(ctx, nil, model.ApiUrl3+"/batch_async_requests/get", req, &result); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}

// BatchAsyncRequestAddSelf 创建批量异步请求任务
// https://developers.e.qq.com/v3.0/docs/api/batch_async_requests/add
func (a *TencentAdapter) BatchAsyncRequestAddSelf(ctx context.Context, req *model.BatchAsyncRequestAddReq) (
	resp *model.BatchAsyncRequestAddResp, err error) {
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
	var result model.BatchAsyncRequestAddResp
	if requestErr := a.RequestPostJson(ctx, headers, model.ApiUrl3+"/batch_async_requests/add?"+globalQuery, req, &result); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}
