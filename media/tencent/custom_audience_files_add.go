package tencent

import (
	"context"
	"strconv"

	"github.com/cngamesdk/media-sdk/media/tencent/model"
	"github.com/cngamesdk/media-sdk/utils"
)

// CustomAudienceFilesAdd 上传客户人群数据文件
// https://developers.e.qq.com/v3.0/docs/api/custom_audience_files/add
func (a *TencentAdapter) CustomAudienceFilesAdd(ctx context.Context, req *model.CustomAudienceFilesAddReq) (
	resp *model.CustomAudienceFilesAddResp, err error) {
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

	fields := make(map[string]string)
	fields["account_id"] = strconv.FormatInt(req.AccountID, 10)
	fields["audience_id"] = strconv.FormatInt(req.AudienceID, 10)
	fields["user_id_type"] = req.UserIDType
	fields["operation_type"] = req.OperationType
	if req.OpenAppID != "" {
		fields["open_app_id"] = req.OpenAppID
	}

	var result model.CustomAudienceFilesAddResp
	if requestErr := a.RequestPostMultipart(
		ctx,
		model.ApiUrl3+"/custom_audience_files/add?"+globalQuery,
		fields,
		"file",
		req.FileName,
		req.File,
		&result,
	); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}
