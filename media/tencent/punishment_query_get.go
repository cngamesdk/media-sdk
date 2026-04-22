package tencent

import (
	"context"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
	"github.com/cngamesdk/media-sdk/utils"
)

// PunishmentQueryGet 获取违规处罚列表
// https://developers.e.qq.com/v3.0/docs/api/punishment_query/get
func (a *TencentAdapter) PunishmentQueryGet(ctx context.Context, req *model.PunishmentQueryGetReq) (
	resp *model.PunishmentQueryGetResp, err error) {
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
	var result model.PunishmentQueryGetResp
	if requestErr := a.RequestPostJson(ctx, headers, model.ApiUrl3+"/punishment_query/get?"+globalQuery, req, &result); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}
