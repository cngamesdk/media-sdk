package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// UnitMonitorUrlsUpdate 批量更新监测链接
func (a *KuaishouAdapter) UnitMonitorUrlsUpdate(ctx context.Context, req *kuaishouModel.UnitMonitorUrlsUpdateReq) (resp *kuaishouModel.UnitMonitorUrlsUpdateResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.UnitMonitorUrlsUpdateResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/v3/unit/batchUpdateMonitorUrls", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
