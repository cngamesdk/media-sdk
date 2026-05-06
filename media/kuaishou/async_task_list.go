package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// AsyncTaskList 获取任务状态
func (a *KuaishouAdapter) AsyncTaskList(ctx context.Context, req *kuaishouModel.AsyncTaskListReq) (resp *kuaishouModel.AsyncTaskListResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.AsyncTaskListResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/v1/async_task/list", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
