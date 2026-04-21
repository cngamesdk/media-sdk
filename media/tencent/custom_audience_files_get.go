package tencent

import (
	"context"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
)

// CustomAudienceFilesGet 获取客户人群数据文件
// https://developers.e.qq.com/v3.0/docs/api/custom_audience_files/get
func (a *TencentAdapter) CustomAudienceFilesGet(ctx context.Context, req *model.CustomAudienceFilesGetReq) (
	resp *model.CustomAudienceFilesGetResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	var result model.CustomAudienceFilesGetResp
	if requestErr := a.RequestGet(ctx, nil, model.ApiUrl3+"/custom_audience_files/get", req, &result); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}
