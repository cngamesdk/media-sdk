package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// WordInfoFileQuery 获取关键词导出文件
func (a *KuaishouAdapter) WordInfoFileQuery(ctx context.Context, req *kuaishouModel.WordInfoFileQueryReq) (resp *kuaishouModel.WordInfoFileQueryResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.WordInfoFileQueryResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/search/keyword/file/query", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
