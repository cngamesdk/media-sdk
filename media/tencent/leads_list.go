package tencent

import (
	"context"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
)

// LeadsListGetSelf 获取线索列表
// https://developers.e.qq.com/v3.0/docs/api/leads_list/get
func (a *TencentAdapter) LeadsListGetSelf(ctx context.Context, req *model.LeadsListGetReq) (
	resp *model.LeadsListGetResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	var result model.LeadsListGetResp
	if requestErr := a.RequestGet(ctx, nil, model.ApiUrl3+"/leads_list/get", req, &result); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}
