package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// AutoProjectConfigUpdate 智投配置更新
func (a *KuaishouAdapter) AutoProjectConfigUpdate(ctx context.Context, req *kuaishouModel.AutoProjectConfigUpdateReq) (resp *kuaishouModel.AutoProjectConfigUpdateResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.AutoProjectConfigUpdateResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/simple/project/config/update", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
