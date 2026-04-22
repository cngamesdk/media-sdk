package tencent

import (
	"context"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
)

// PunishmentConfigGet 获取处罚系统配置
// https://developers.e.qq.com/v3.0/docs/api/punishment_config/get
func (a *TencentAdapter) PunishmentConfigGet(ctx context.Context, req *model.PunishmentConfigGetReq) (
	resp *model.PunishmentConfigGetResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	var result model.PunishmentConfigGetResp
	if requestErr := a.RequestGet(ctx, nil, model.ApiUrl3+"/punishment_config/get", req, &result); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}
