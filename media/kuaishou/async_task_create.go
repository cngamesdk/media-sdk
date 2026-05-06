package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// AsyncTaskCreate 创建历史数据查询任务
func (a *KuaishouAdapter) AsyncTaskCreate(ctx context.Context, req *kuaishouModel.AsyncTaskCreateReq) (resp *kuaishouModel.AsyncTaskCreateResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.AsyncTaskCreateResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/v1/async_task/create", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
