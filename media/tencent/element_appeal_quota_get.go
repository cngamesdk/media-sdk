package tencent

import (
	"context"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
)

// ElementAppealQuotaGet 获取元素申诉复审配额
// https://developers.e.qq.com/v3.0/docs/api/element_appeal_quota/get
func (a *TencentAdapter) ElementAppealQuotaGet(ctx context.Context, req *model.ElementAppealQuotaGetReq) (
	resp *model.ElementAppealQuotaGetResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	var result model.ElementAppealQuotaGetResp
	if requestErr := a.RequestGet(ctx, nil, model.ApiUrl3+"/element_appeal_quota/get", req, &result); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}
