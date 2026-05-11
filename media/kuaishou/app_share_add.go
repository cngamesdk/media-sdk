package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// AppShareAdd 添加应用共享
// https://ad.e.kuaishou.com/rest/openapi/gw/dsp/appcenter/share/add
func (a *KuaishouAdapter) AppShareAdd(ctx context.Context, req *kuaishouModel.AppShareAddReq) (resp *kuaishouModel.AppShareAddResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.AppShareAddResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/appcenter/share/add", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
