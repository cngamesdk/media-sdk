package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// DmpDatasourceUpdate 数据源更新
// https://ad.e.kuaishou.com/rest/openapi/gw/dmp/v1/datasource/update
func (a *KuaishouAdapter) DmpDatasourceUpdate(ctx context.Context, req *kuaishouModel.DmpDatasourceUpdateReq) (resp *kuaishouModel.DmpDatasourceUpdateResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.DmpDatasourceUpdateResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dmp/v1/datasource/update", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
