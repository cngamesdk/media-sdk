package tencent

import (
	"context"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
	"github.com/cngamesdk/media-sdk/utils"
)

// EcommerceOrderGetSelf 获取订单数据
// https://developers.e.qq.com/v3.0/docs/api/ecommerce_order/get
func (a *TencentAdapter) EcommerceOrderGetSelf(ctx context.Context, req *model.EcommerceOrderGetReq) (
	resp *model.EcommerceOrderGetResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	var result model.EcommerceOrderGetResp
	if requestErr := a.RequestGet(ctx, nil, model.ApiUrl3+"/ecommerce_order/get", req, &result); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}

// EcommerceOrderUpdateSelf 更新订单状态
// https://developers.e.qq.com/v3.0/docs/api/ecommerce_order/update
func (a *TencentAdapter) EcommerceOrderUpdateSelf(ctx context.Context, req *model.EcommerceOrderUpdateReq) (
	resp *model.EcommerceOrderUpdateResp, err error) {
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
	var result model.EcommerceOrderUpdateResp
	if requestErr := a.RequestPostJson(ctx, headers, model.ApiUrl3+"/ecommerce_order/update?"+globalQuery, req, &result); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}
