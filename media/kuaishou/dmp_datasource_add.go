package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// DmpDatasourceAdd 数据源新建
// https://ad.e.kuaishou.com/rest/openapi/gw/dmp/v1/datasource/add
func (a *KuaishouAdapter) DmpDatasourceAdd(ctx context.Context, req *kuaishouModel.DmpDatasourceAddReq) (resp *kuaishouModel.DmpDatasourceAddResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.DmpDatasourceAddResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dmp/v1/datasource/add", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
