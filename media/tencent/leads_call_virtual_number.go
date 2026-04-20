package tencent

import (
	"context"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
)

// LeadsCallVirtualNumberGetSelf 获取中间号
// https://developers.e.qq.com/v3.0/docs/api/leads_call_virtual_number/get
func (a *TencentAdapter) LeadsCallVirtualNumberGetSelf(ctx context.Context, req *model.LeadsCallVirtualNumberGetReq) (
	resp *model.LeadsCallVirtualNumberGetResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	var result model.LeadsCallVirtualNumberGetResp
	if requestErr := a.RequestGet(ctx, nil, model.ApiUrl3+"/leads_call_virtual_number/get", req, &result); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}
