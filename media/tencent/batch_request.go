package tencent

import (
	"context"

	"github.com/cngamesdk/media-sdk/media/tencent/model"
	"github.com/cngamesdk/media-sdk/utils"
)

// BatchRequestAddSelf 创建批量请求
// https://developers.e.qq.com/v3.0/docs/api/batch_requests/add
func (a *TencentAdapter) BatchRequestAddSelf(ctx context.Context, req *model.BatchRequestAddReq) (
	resp *model.BatchRequestAddResp, err error) {
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
	var result model.BatchRequestAddResp
	if requestErr := a.RequestPostJson(ctx, headers, model.ApiUrl3+"/batch_requests/add?"+globalQuery, req, &result); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}
