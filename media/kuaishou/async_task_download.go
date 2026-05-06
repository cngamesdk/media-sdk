package kuaishou

import (
	"context"
	"fmt"
	"net/http"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// AsyncTaskDownload 数据下载（GET请求，返回CSV文件内容）
func (a *KuaishouAdapter) AsyncTaskDownload(ctx context.Context, req *kuaishouModel.AsyncTaskDownloadReq) (resp *kuaishouModel.AsyncTaskDownloadResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	downloadUrl := fmt.Sprintf("%s/rest/openapi/v1/async_task/download?advertiser_id=%d&task_id=%d",
		kuaishouModel.AdUrl, req.AdvertiserId, req.TaskId)
	data, errRequest := a.Client.Request(ctx, http.MethodGet, downloadUrl, nil, headers)
	if errRequest != nil {
		err = errRequest
		return
	}
	resp = &kuaishouModel.AsyncTaskDownloadResp{
		FileData: data,
	}
	return
}
