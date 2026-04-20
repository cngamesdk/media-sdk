package tencent

import (
	"context"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
)

// LeadsCallRecordGetSelf 获取通话结果
// https://developers.e.qq.com/v3.0/docs/api/leads_call_record/get
func (a *TencentAdapter) LeadsCallRecordGetSelf(ctx context.Context, req *model.LeadsCallRecordGetReq) (
	resp *model.LeadsCallRecordGetResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	var result model.LeadsCallRecordGetResp
	if requestErr := a.RequestGet(ctx, nil, model.ApiUrl3+"/leads_call_record/get", req, &result); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}
