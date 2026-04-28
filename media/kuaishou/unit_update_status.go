package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// UnitUpdateStatus 修改广告组状态
func (a *KuaishouAdapter) UnitUpdateStatus(ctx context.Context, req *kuaishouModel.UnitUpdateStatusReq) (resp *kuaishouModel.UnitUpdateStatusResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.UnitUpdateStatusResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/v1/ad_unit/update/status", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
