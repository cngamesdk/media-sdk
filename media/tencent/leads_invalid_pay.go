package tencent

import (
	"context"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
)

// LeadsInvalidPayGetSelf 获取无效赔付明细
// https://developers.e.qq.com/v3.0/docs/api/leads_invalid_pay/get
func (a *TencentAdapter) LeadsInvalidPayGetSelf(ctx context.Context, req *model.LeadsInvalidPayGetReq) (
	resp *model.LeadsInvalidPayGetResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	var result model.LeadsInvalidPayGetResp
	if requestErr := a.RequestGet(ctx, nil, model.ApiUrl3+"/leads_invalid_pay/get", req, &result); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}
