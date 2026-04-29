package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// WordInfoUpdateStatus 修改关键词投放状态
func (a *KuaishouAdapter) WordInfoUpdateStatus(ctx context.Context, req *kuaishouModel.WordInfoUpdateStatusReq) (resp *kuaishouModel.WordInfoUpdateStatusResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.WordInfoUpdateStatusResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/v2/word_info/update/status", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
