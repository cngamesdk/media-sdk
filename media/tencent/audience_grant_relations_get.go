package tencent

import (
	"context"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
)

// AudienceGrantRelationsGet 获取人群授权信息
// https://developers.e.qq.com/v3.0/docs/api/audience_grant_relations/get
func (a *TencentAdapter) AudienceGrantRelationsGet(ctx context.Context, req *model.AudienceGrantRelationsGetReq) (
	resp *model.AudienceGrantRelationsGetResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	var result model.AudienceGrantRelationsGetResp
	if requestErr := a.RequestGet(ctx, nil, model.ApiUrl3+"/audience_grant_relations/get", req, &result); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}
