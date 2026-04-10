package tencent

import (
	"context"

	"github.com/cngamesdk/media-sdk/media/tencent/model"
	"github.com/cngamesdk/media-sdk/utils"
)

// ExtendPackageAddSelf 创建应用分包
// https://developers.e.qq.com/v3.0/docs/api/extend_package/add
func (a *TencentAdapter) ExtendPackageAddSelf(ctx context.Context, req *model.ExtendPackageAddReq) (
	resp *model.ExtendPackageAddResp, err error) {
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
	var result model.ExtendPackageAddResp
	if requestErr := a.RequestPostJson(ctx, headers, model.ApiUrl3+"/extend_package/add?"+globalQuery, req, &result); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}

// ExtendPackageUpdateSelf 更新应用子包版本
// https://developers.e.qq.com/v3.0/docs/api/extend_package/update
func (a *TencentAdapter) ExtendPackageUpdateSelf(ctx context.Context, req *model.ExtendPackageUpdateReq) (
	resp *model.ExtendPackageUpdateResp, err error) {
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
	var result model.ExtendPackageUpdateResp
	if requestErr := a.RequestPostJson(ctx, headers, model.ApiUrl3+"/extend_package/update?"+globalQuery, req, &result); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}
