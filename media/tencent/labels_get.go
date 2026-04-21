package tencent

import (
	"context"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
)

// LabelsGet 标签广场标签获取
// https://developers.e.qq.com/v3.0/docs/api/labels/get
func (a *TencentAdapter) LabelsGet(ctx context.Context, req *model.LabelsGetReq) (
	resp *model.LabelsGetResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	var result model.LabelsGetResp
	if requestErr := a.RequestGet(ctx, nil, model.ApiUrl3+"/labels/get", req, &result); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}
