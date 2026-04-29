package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// WordInfoUpdateMatchType 修改关键词匹配方式
func (a *KuaishouAdapter) WordInfoUpdateMatchType(ctx context.Context, req *kuaishouModel.WordInfoUpdateMatchTypeReq) (resp *kuaishouModel.WordInfoUpdateMatchTypeResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.WordInfoUpdateMatchTypeResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/v2/word_info/update/match_type", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
