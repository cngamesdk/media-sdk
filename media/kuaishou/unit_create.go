package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// UnitCreate 创建广告组
func (a *KuaishouAdapter) UnitCreate(ctx context.Context, req *kuaishouModel.UnitCreateReq) (resp *kuaishouModel.UnitCreateResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.UnitCreateResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/unit/create", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
