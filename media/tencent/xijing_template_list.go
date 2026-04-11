package tencent

import (
	"context"

	"github.com/cngamesdk/media-sdk/media/tencent/model"
)

// XijingTemplateListGetSelf 获取蹊径落地页模板列表
// https://developers.e.qq.com/v3.0/docs/api/xijing_template_list/get
func (a *TencentAdapter) XijingTemplateListGetSelf(ctx context.Context, req *model.XijingTemplateListGetReq) (
	resp *model.XijingTemplateListGetResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	var result model.XijingTemplateListGetResp
	if requestErr := a.RequestGet(ctx, nil, model.ApiUrl3+"/xijing_template_list/get", req, &result); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}
