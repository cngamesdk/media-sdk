package tencent

import (
	"context"
	"strconv"

	"github.com/cngamesdk/media-sdk/media/tencent/model"
	"github.com/cngamesdk/media-sdk/utils"
)

// ImageGetSelf 获取图片信息
// https://developers.e.qq.com/v3.0/docs/api/images/get
func (a *TencentAdapter) ImageGetSelf(ctx context.Context, req *model.ImageGetReq) (
	resp *model.ImageGetResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	var result model.ImageGetResp
	if requestErr := a.RequestGet(ctx, nil, model.ApiUrl3+"/images/get", req, &result); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}

// ImageUpdateSelf 修改图片信息
// https://developers.e.qq.com/v3.0/docs/api/images/update
func (a *TencentAdapter) ImageUpdateSelf(ctx context.Context, req *model.ImageUpdateReq) (
	resp *model.ImageUpdateResp, err error) {
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
	var result model.ImageUpdateResp
	if requestErr := a.RequestPostJson(ctx, headers, model.ApiUrl3+"/images/update?"+globalQuery, req, &result); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}

// ImageDeleteSelf 删除图片
// https://developers.e.qq.com/v3.0/docs/api/images/delete
func (a *TencentAdapter) ImageDeleteSelf(ctx context.Context, req *model.ImageDeleteReq) (
	resp *model.ImageDeleteResp, err error) {
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
	var result model.ImageDeleteResp
	if requestErr := a.RequestPostJson(ctx, headers, model.ApiUrl3+"/images/delete?"+globalQuery, req, &result); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}

// https://developers.e.qq.com/v3.0/docs/api/images/add
func (a *TencentAdapter) ImageAddSelf(ctx context.Context, req *model.ImageAddReq) (
	resp *model.ImageAddResp, err error) {
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
	if req.AccountID != 0 {
		fields["account_id"] = strconv.FormatInt(req.AccountID, 10)
	}
	if req.OrganizationID != 0 {
		fields["organization_id"] = strconv.FormatInt(req.OrganizationID, 10)
	}
	fields["upload_type"] = req.UploadType
	fields["signature"] = req.Signature
	if req.ImageUsage != "" {
		fields["image_usage"] = req.ImageUsage
	}
	if req.Description != "" {
		fields["description"] = req.Description
	}
	if req.ResizeWidth != 0 {
		fields["resize_width"] = strconv.Itoa(req.ResizeWidth)
	}
	if req.ResizeHeight != 0 {
		fields["resize_height"] = strconv.Itoa(req.ResizeHeight)
	}
	if req.ResizeFileSize != 0 {
		fields["resize_file_size"] = strconv.Itoa(req.ResizeFileSize)
	}

	apiURL := model.ApiUrl3 + "/images/add?" + globalQuery
	var result model.ImageAddResp

	if req.UploadType == model.ImageUploadTypeFile {
		if requestErr := a.RequestPostMultipart(
			ctx,
			apiURL,
			fields,
			"file",
			req.ImageFileName,
			req.ImageFile,
			&result,
		); requestErr != nil {
			err = requestErr
			return
		}
	} else {
		fields["bytes"] = req.Bytes
		if requestErr := a.RequestPostMultipartFields(ctx, apiURL, fields, &result); requestErr != nil {
			err = requestErr
			return
		}
	}
	resp = &result
	return
}
