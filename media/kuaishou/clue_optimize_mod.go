package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// ClueOptimizeSwitchModSelf 修改线索优选开关状态
func (a *KuaishouAdapter) ClueOptimizeSwitchModSelf(ctx context.Context, req *kuaishouModel.ClueOptimizeSwitchModReq) (resp *kuaishouModel.ClueOptimizeSwitchModResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.ClueOptimizeSwitchModResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/clueOptimize/ocpxSwitch/mod", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
