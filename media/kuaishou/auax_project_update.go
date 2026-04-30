package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// AuaxProjectUpdate 智投项目更新
func (a *KuaishouAdapter) AuaxProjectUpdate(ctx context.Context, req *kuaishouModel.AuaxProjectUpdateReq) (resp *kuaishouModel.AuaxProjectCreateResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.AuaxProjectCreateResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/v2/auax/project/update", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
