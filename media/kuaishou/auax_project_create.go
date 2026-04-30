package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// AuaxProjectCreate 智投项目创建
func (a *KuaishouAdapter) AuaxProjectCreate(ctx context.Context, req *kuaishouModel.AuaxProjectCreateReq) (resp *kuaishouModel.AuaxProjectCreateResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.AuaxProjectCreateResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/v2/auax/project/create", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
