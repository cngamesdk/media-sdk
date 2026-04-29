package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// WordInfoNegativeWordUpdate 更新否定词
func (a *KuaishouAdapter) WordInfoNegativeWordUpdate(ctx context.Context, req *kuaishouModel.WordInfoNegativeWordUpdateReq) (resp *kuaishouModel.WordInfoNegativeWordUpdateResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.WordInfoNegativeWordUpdateResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/v1/ad_unit/update/negative_word", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
