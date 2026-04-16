package tencent

import (
	"context"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
	"github.com/cngamesdk/media-sdk/utils"
)

// AsyncReportsAddSelf 创建异步报表任务
// https://developers.e.qq.com/v3.0/docs/api/async_reports/add
func (a *TencentAdapter) AsyncReportsAddSelf(ctx context.Context, req *model.AsyncReportsAddReq) (
	resp *model.AsyncReportsAddResp, err error) {
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
	var result model.AsyncReportsAddResp
	if requestErr := a.RequestPostJson(ctx, headers, model.ApiUrl3+"/async_reports/add?"+globalQuery, req, &result); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}

// AsyncReportsGetSelf 获取异步报表任务
// https://developers.e.qq.com/v3.0/docs/api/async_reports/get
func (a *TencentAdapter) AsyncReportsGetSelf(ctx context.Context, req *model.AsyncReportsGetReq) (
	resp *model.AsyncReportsGetResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	var result model.AsyncReportsGetResp
	if requestErr := a.RequestGet(ctx, nil, model.ApiUrl3+"/async_reports/get", req, &result); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}

// AsyncReportFilesGetSelf 获取文件接口（异步报表文件下载）
// https://developers.e.qq.com/v3.0/docs/api/async_report_files/get
// 注意：该接口返回重定向到文件下载地址，响应为文件二进制内容，非JSON
func (a *TencentAdapter) AsyncReportFilesGetSelf(ctx context.Context, req *model.AsyncReportFilesGetReq) (
	resp *model.AsyncReportFilesGetResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	query, queryErr := utils.ConvertStructToQueryString(req)
	if queryErr != nil {
		err = queryErr
		return
	}
	url := model.DlUrl3 + "/async_report_files/get?" + query
	data, requestErr := a.Client.Get(ctx, url, nil)
	if requestErr != nil {
		err = requestErr
		return
	}
	resp = &model.AsyncReportFilesGetResp{
		FileData: data,
	}
	return
}
