package tencent

import (
	"context"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
)

// OrganizationAccountRelationGetSelf 查询组织下广告账户信息
// https://developers.e.qq.com/v3.0/docs/api/organization_account_relation/get
func (a *TencentAdapter) OrganizationAccountRelationGetSelf(ctx context.Context, req *model.OrganizationAccountRelationGetReq) (
	resp *model.OrganizationAccountRelationGetResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	var result model.OrganizationAccountRelationGetResp
	if requestErr := a.RequestGet(ctx, nil, model.ApiUrl3+"/organization_account_relation/get", req, &result); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}
