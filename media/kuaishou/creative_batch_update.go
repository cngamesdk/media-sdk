package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// CreativeBatchUpdate 批量修改自定义创意
func (a *KuaishouAdapter) CreativeBatchUpdate(ctx context.Context, req *kuaishouModel.CreativeBatchUpdateReq) (resp *kuaishouModel.CreativeBatchUpdateResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.CreativeBatchUpdateResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/creative/batch_update", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
