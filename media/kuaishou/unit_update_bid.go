package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// UnitUpdateBid 修改广告组出价
func (a *KuaishouAdapter) UnitUpdateBid(ctx context.Context, req *kuaishouModel.UnitUpdateBidReq) (resp *kuaishouModel.UnitUpdateBidResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.UnitUpdateBidResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/v1/ad_unit/update/bid", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
