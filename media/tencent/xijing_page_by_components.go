package tencent

import (
	"context"

	"github.com/cngamesdk/media-sdk/media/tencent/model"
	"github.com/cngamesdk/media-sdk/utils"
)

// XijingPageByCompAddSelf 蹊径基于组件创建落地页
// https://developers.e.qq.com/v3.0/docs/api/xijing_page_by_components/add
func (a *TencentAdapter) XijingPageByCompAddSelf(ctx context.Context, req *model.XijingPageByCompAddReq) (
	resp *model.XijingPageByCompAddResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	globalQuery, globalQueryErr := utils.ConvertStructToQueryString(req.GlobalReq)
	if globalQueryErr != nil {
		err = globalQueryErr
		return
	}
	req.GlobalReq.Clear()
	headers := make(model.Headers)
	headers.Json()
	var result model.XijingPageByCompAddResp
	if requestErr := a.RequestPostJson(ctx, headers, model.ApiUrl3+"/xijing_page_by_components/add?"+globalQuery, req, &result); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}
