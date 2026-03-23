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

// EventAssetsDetailSelf 获取已创建资产详情（新）
// https://open.oceanengine.com/labels/7/docs/1800988620664954?origin=left_nav
func (a *ToutiaoAdapter) EventAssetsDetailSelf(ctx context.Context, req *model2.EventAssetsDetailReq) (resp *model2.EventAssetsDetailResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result model2.EventAssetsDetailResp
	errRequest := a.RequestGet(ctx, headers, model2.BaseUrlApi+"/open_api/2/tools/event/all_assets/detail/", req, &result)
	if errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}

// EventManagerAvailableEventsGetSelf 获取可创建事件列表
// https://open.oceanengine.com/labels/7/docs/1709793059412996?origin=left_nav
func (a *ToutiaoAdapter) EventManagerAvailableEventsGetSelf(ctx context.Context, req *model2.EventManagerAvailableEventsGetReq) (resp *model2.EventManagerAvailableEventsGetResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result model2.EventManagerAvailableEventsGetResp
	errRequest := a.RequestGet(ctx, headers, model2.BaseUrlApi+"/open_api/2/event_manager/available_events/get/", req, &result)
	if errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}

// EventManagerEventsCreateSelf 资产下创建事件
// https://open.oceanengine.com/labels/7/docs/1709792900524035?origin=left_nav
func (a *ToutiaoAdapter) EventManagerEventsCreateSelf(ctx context.Context, req *model2.EventManagerEventsCreateReq) (resp *model2.EventManagerEventsCreateResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result model2.EventManagerEventsCreateResp
	errRequest := a.RequestPostJson(ctx, headers, model2.BaseUrlAd+"/open_api/2/event_manager/events/create/", req, &result)
	if errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}

// EventManagerEventConfigsGetSelf 获取资产下已创建事件列表
// https://open.oceanengine.com/labels/7/docs/1709793086075972?origin=left_nav
func (a *ToutiaoAdapter) EventManagerEventConfigsGetSelf(ctx context.Context, req *model2.EventManagerEventConfigsGetReq) (resp *model2.EventManagerEventConfigsGetResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result model2.EventManagerEventConfigsGetResp
	errRequest := a.RequestGet(ctx, headers, model2.BaseUrlAd+"/open_api/2/event_manager/event_configs/get/", req, &result)
	if errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}

// EventManagerOptimizedGoalGetSelf 获取可用优化目标（巨量营销升级版）
// https://open.oceanengine.com/labels/7/docs/1740944984250381?origin=left_nav
func (a *ToutiaoAdapter) EventManagerOptimizedGoalGetSelf(ctx context.Context, req *model2.EventManagerOptimizedGoalGetReq) (resp *model2.EventManagerOptimizedGoalGetResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result model2.EventManagerOptimizedGoalGetResp
	errRequest := a.RequestGet(ctx, headers, model2.BaseUrlApi+"/open_api/v3.0/event_manager/optimized_goal/get_v2/", req, &result)
	if errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
