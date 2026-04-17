package tencent

import (
	"context"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
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
