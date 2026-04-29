package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// CreativeReviewDetail 获取创意审核详情
func (a *KuaishouAdapter) CreativeReviewDetail(ctx context.Context, req *kuaishouModel.CreativeReviewDetailReq) (resp *kuaishouModel.CreativeReviewDetailResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.CreativeReviewDetailResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/creative/element/reviewDetails", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
