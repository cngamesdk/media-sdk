package toutiao

import (
	"context"
	model2 "github.com/cngamesdk/media-sdk/media/toutiao/model"
)

// EventManagerAssetsCreateSelf 创建事件资产
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

// EventAssetsListSelf 获取账户下资产列表（新）
// https://open.oceanengine.com/labels/7/docs/1800985709803914?origin=left_nav
func (a *ToutiaoAdapter) EventAssetsListSelf(ctx context.Context, req *model2.EventAssetsListReq) (resp *model2.EventAssetsListResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result model2.EventAssetsListResp
	errRequest := a.RequestGet(ctx, headers, model2.BaseUrlApi+"/open_api/2/tools/event/all_assets/list/", req, &result)
	if errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
