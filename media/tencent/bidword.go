package tencent

import (
	"context"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
	"github.com/cngamesdk/media-sdk/utils"
)

// BidwordAddSelf 创建关键词
// https://developers.e.qq.com/v3.0/docs/api/bidword/add
func (a *TencentAdapter) BidwordAddSelf(ctx context.Context, req *model.BidwordAddReq) (
	resp *model.BidwordAddResp, err error) {
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
	var result model.BidwordAddResp
	if requestErr := a.RequestPostJson(ctx, headers, model.ApiUrl3+"/bidword/add?"+globalQuery, req, &result); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}

// BidwordUpdateSelf 更新关键词
// https://developers.e.qq.com/v3.0/docs/api/bidword/update
func (a *TencentAdapter) BidwordUpdateSelf(ctx context.Context, req *model.BidwordUpdateReq) (
	resp *model.BidwordUpdateResp, err error) {
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
	var result model.BidwordUpdateResp
	if requestErr := a.RequestPostJson(ctx, headers, model.ApiUrl3+"/bidword/update?"+globalQuery, req, &result); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}

// BidwordDeleteSelf 删除关键词
// https://developers.e.qq.com/v3.0/docs/api/bidword/delete
func (a *TencentAdapter) BidwordDeleteSelf(ctx context.Context, req *model.BidwordDeleteReq) (
	resp *model.BidwordDeleteResp, err error) {
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
	var result model.BidwordDeleteResp
	if requestErr := a.RequestPostJson(ctx, headers, model.ApiUrl3+"/bidword/delete?"+globalQuery, req, &result); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}

// BidwordGetSelf 查询关键词
// https://developers.e.qq.com/v3.0/docs/api/bidword/get
func (a *TencentAdapter) BidwordGetSelf(ctx context.Context, req *model.BidwordGetReq) (
	resp *model.BidwordGetResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	var result model.BidwordGetResp
	if requestErr := a.RequestGet(ctx, nil, model.ApiUrl3+"/bidword/get", req, &result); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}
