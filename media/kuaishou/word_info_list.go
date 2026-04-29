package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// WordInfoList 获取关键词列表
func (a *KuaishouAdapter) WordInfoList(ctx context.Context, req *kuaishouModel.WordInfoListReq) (resp *kuaishouModel.WordInfoListResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.WordInfoListResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/v2/word_info/list", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
