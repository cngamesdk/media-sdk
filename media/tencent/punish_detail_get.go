package tencent

import (
	"context"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
	"github.com/cngamesdk/media-sdk/utils"
)

// PunishDetailGet 获取计量处罚明细
// https://developers.e.qq.com/v3.0/docs/api/punish_detail/get
func (a *TencentAdapter) PunishDetailGet(ctx context.Context, req *model.PunishDetailGetReq) (
	resp *model.PunishDetailGetResp, err error) {
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
	var result model.PunishDetailGetResp
	if requestErr := a.RequestPostJson(ctx, headers, model.ApiUrl3+"/punish_detail/get?"+globalQuery, req, &result); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}
