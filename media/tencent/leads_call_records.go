package tencent

import (
	"context"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
)

// LeadsCallRecordsGetSelf 获取一个账号下的全部通话结果
// https://developers.e.qq.com/v3.0/docs/api/leads_call_records/get
func (a *TencentAdapter) LeadsCallRecordsGetSelf(ctx context.Context, req *model.LeadsCallRecordsGetReq) (
	resp *model.LeadsCallRecordsGetResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	var result model.LeadsCallRecordsGetResp
	if requestErr := a.RequestGet(ctx, nil, model.ApiUrl3+"/leads_call_records/get", req, &result); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}
