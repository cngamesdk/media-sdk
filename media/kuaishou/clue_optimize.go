package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// ClueOptimizeSwitchSelf 获取线索优选开关状态
func (a *KuaishouAdapter) ClueOptimizeSwitchSelf(ctx context.Context, req *kuaishouModel.ClueOptimizeSwitchReq) (resp *kuaishouModel.ClueOptimizeSwitchResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.ClueOptimizeSwitchResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/clueOptimize/ocpxSwitch", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
