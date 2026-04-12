package tencent

import (
	"context"

	"github.com/cngamesdk/media-sdk/media/tencent/model"
)

// XijingComplexTemplateGetSelf 获取蹊径落地页互动模板配置
// https://developers.e.qq.com/v3.0/docs/api/xijing_complex_template/get
func (a *TencentAdapter) XijingComplexTemplateGetSelf(ctx context.Context, req *model.XijingComplexTemplateGetReq) (
	resp *model.XijingComplexTemplateGetResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	var result model.XijingComplexTemplateGetResp
	if requestErr := a.RequestGet(ctx, nil, model.ApiUrl3+"/xijing_complex_template/get", req, &result); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}
