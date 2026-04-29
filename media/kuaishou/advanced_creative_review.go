package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// AdvancedCreativeReview 获取程序化创意/智能创意审核信息
func (a *KuaishouAdapter) AdvancedCreativeReview(ctx context.Context, req *kuaishouModel.AdvancedCreativeReviewReq) (resp *kuaishouModel.AdvancedCreativeReviewResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.AdvancedCreativeReviewResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/v2/creative/advanced/program/review_detail", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
