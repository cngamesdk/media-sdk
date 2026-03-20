package toutiao

import (
	"context"
	model2 "github.com/cngamesdk/media-sdk/media/toutiao/model"
)

// DataConversionClickMonitorLinkSelf 获取点击监测链接
// https://open.oceanengine.com/labels/7/docs/1696710655781900?origin=left_nav
func (a *ToutiaoAdapter) DataConversionClickMonitorLinkSelf(ctx context.Context) model2.CustomMacros {
	_ = ctx
	return model2.ClickMacros
}

// DataConversionShowMonitorLinkSelf 获取点击监测链接
// https://open.oceanengine.com/labels/7/docs/1696710655781900?origin=left_nav
func (a *ToutiaoAdapter) DataConversionShowMonitorLinkSelf(ctx context.Context) model2.CustomMacros {
	_ = ctx
	return model2.ShowMacros
}

// DataConversionEventReportSelf 数据转化事件上报
// https://open.oceanengine.com/labels/7/docs/1696710655781900?origin=left_nav
func (a *ToutiaoAdapter) DataConversionEventReportSelf(ctx context.Context, req *model2.ConversionEventReq) (resp *model2.ConversionEventResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	var result model2.ConversionEventResp
	url := req.CallbackUrl
	req.CallbackUrl = ""
	errRequest := a.Media.RequestGet(ctx, nil, url, req, &result)
	if errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
