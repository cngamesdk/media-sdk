package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// CreativePreview 创意体验
func (a *KuaishouAdapter) CreativePreview(ctx context.Context, req *kuaishouModel.CreativePreviewReq) (resp *kuaishouModel.CreativePreviewResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.CreativePreviewResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/v1/creative/preview", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
