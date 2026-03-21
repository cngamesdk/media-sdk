package toutiao

import (
	"context"
	model2 "github.com/cngamesdk/media-sdk/media/toutiao/model"
)

// ProjectCreateSelf 创建事件资产
// https://open.oceanengine.com/labels/7/docs/1850398228888576?origin=left_nav
func (a *ToutiaoAdapter) EventManagerAssetsCreateSelf(ctx context.Context, req *model2.EventManagerAssetsCreateReq) (resp *model2.EventManagerAssetsCreateResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result model2.EventManagerAssetsCreateResp
	errRequest := a.RequestPostJson(ctx, headers, model2.BaseUrlApi+"/open_api/2/event_manager/assets/create/", req, &result)
	if errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
