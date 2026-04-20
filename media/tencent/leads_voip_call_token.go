package tencent

import (
	"context"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
)

// LeadsVoipCallTokenGetSelf 获取网络电话 token
// https://developers.e.qq.com/v3.0/docs/api/leads_voip_call_token/get
func (a *TencentAdapter) LeadsVoipCallTokenGetSelf(ctx context.Context, req *model.LeadsVoipCallTokenGetReq) (
	resp *model.LeadsVoipCallTokenGetResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	var result model.LeadsVoipCallTokenGetResp
	if requestErr := a.RequestGet(ctx, nil, model.ApiUrl3+"/leads_voip_call_token/get", req, &result); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}
