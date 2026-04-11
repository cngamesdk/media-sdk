package tencent

import (
	"context"

	"github.com/cngamesdk/media-sdk/media/tencent/model"
)

// XijingTemplateGetSelf 获取蹊径落地页模板
// https://developers.e.qq.com/v3.0/docs/api/xijing_template/get
func (a *TencentAdapter) XijingTemplateGetSelf(ctx context.Context, req *model.XijingTemplateGetReq) (
	resp *model.XijingTemplateGetResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	var result model.XijingTemplateGetResp
	if requestErr := a.RequestGet(ctx, nil, model.ApiUrl3+"/xijing_template/get", req, &result); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}
