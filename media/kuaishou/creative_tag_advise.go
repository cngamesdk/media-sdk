package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// CreativeTagAdvise 创意标签填写建议
func (a *KuaishouAdapter) CreativeTagAdvise(ctx context.Context, req *kuaishouModel.CreativeTagAdviseReq) (resp *kuaishouModel.CreativeTagAdviseResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.CreativeTagAdviseResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/v1/creative/creative_tag/advise", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
