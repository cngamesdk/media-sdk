package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// AdvCardShare 推送高级创意
func (a *KuaishouAdapter) AdvCardShare(ctx context.Context, req *kuaishouModel.AdvCardShareReq) (resp *kuaishouModel.AdvCardShareResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.AdvCardShareResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/v1/adv_card/share", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
