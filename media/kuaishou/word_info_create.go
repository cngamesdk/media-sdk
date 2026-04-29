package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// WordInfoCreate 创建关键词
func (a *KuaishouAdapter) WordInfoCreate(ctx context.Context, req *kuaishouModel.WordInfoCreateReq) (resp *kuaishouModel.WordInfoCreateResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.WordInfoCreateResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/v2/word_info/create", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
