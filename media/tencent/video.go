package tencent

import (
	"context"
	"strconv"

	"github.com/cngamesdk/media-sdk/media/tencent/model"
	"github.com/cngamesdk/media-sdk/utils"
)

// VideoGetSelf 获取视频文件
// https://developers.e.qq.com/v3.0/docs/api/videos/get
func (a *TencentAdapter) VideoGetSelf(ctx context.Context, req *model.VideoGetReq) (
	resp *model.VideoGetResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	var result model.VideoGetResp
	if requestErr := a.RequestGet(ctx, nil, model.ApiUrl3+"/videos/get", req, &result); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}

// VideoAddSelf 添加视频文件
// https://developers.e.qq.com/v3.0/docs/api/videos/add
func (a *TencentAdapter) VideoAddSelf(ctx context.Context, req *model.VideoAddReq) (
	resp *model.VideoAddResp, err error) {
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

	// 构建 multipart 表单字段
	fields := make(map[string]string)
	if req.AccountID != 0 {
		fields["account_id"] = strconv.FormatInt(req.AccountID, 10)
	}
	if req.OrganizationID != 0 {
		fields["organization_id"] = strconv.FormatInt(req.OrganizationID, 10)
	}
	fields["signature"] = req.Signature
	if req.Description != "" {
		fields["description"] = req.Description
	}
	if req.AdcreativeTemplateID != 0 {
		fields["adcreative_template_id"] = strconv.FormatInt(req.AdcreativeTemplateID, 10)
	}

	var result model.VideoAddResp
	if requestErr := a.RequestPostMultipart(
		ctx,
		model.ApiUrl3+"/videos/add?"+globalQuery,
		fields,
		"video_file",
		req.VideoFileName,
		req.VideoFile,
		&result,
	); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}

// VideoUpdateSelf 修改视频信息
// https://developers.e.qq.com/v3.0/docs/api/videos/update
func (a *TencentAdapter) VideoUpdateSelf(ctx context.Context, req *model.VideoUpdateReq) (
	resp *model.VideoUpdateResp, err error) {
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
	var result model.VideoUpdateResp
	if requestErr := a.RequestPostJson(ctx, headers, model.ApiUrl3+"/videos/update?"+globalQuery, req, &result); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}
