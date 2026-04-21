package tencent

import (
	"context"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
	"github.com/cngamesdk/media-sdk/utils"
)

// AudienceGrantRelationsAdd 添加人群授权
// https://developers.e.qq.com/v3.0/docs/api/audience_grant_relations/add
func (a *TencentAdapter) AudienceGrantRelationsAdd(ctx context.Context, req *model.AudienceGrantRelationsAddReq) (
	resp *model.AudienceGrantRelationsAddResp, err error) {
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
	var result model.AudienceGrantRelationsAddResp
	if requestErr := a.RequestPostJson(ctx, headers, model.ApiUrl3+"/audience_grant_relations/add?"+globalQuery, req, &result); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}
