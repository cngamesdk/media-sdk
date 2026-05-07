package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// VideoAutoShareSwitch 查询账号共享视频库按钮是否开启
func (a *KuaishouAdapter) VideoAutoShareSwitch(ctx context.Context, req *kuaishouModel.VideoAutoShareSwitchReq) (resp *kuaishouModel.VideoAutoShareSwitchResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.VideoAutoShareSwitchResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/video/queryAutoShareSwitch", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
