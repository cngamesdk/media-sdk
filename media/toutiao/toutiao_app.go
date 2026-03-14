package toutiao

import (
	model2 "github.com/cngamesdk/media-sdk/media/toutiao/model"
	"golang.org/x/net/context"
)

// EbpAppListSelf 获取安卓应用列表
// https://open.oceanengine.com/labels/7/docs/1846773030696265?origin=left_nav
func (a *ToutiaoAdapter) EbpAppListSelf(ctx context.Context, req *model2.EbpAppListReq) (resp *model2.EbpAppListResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result model2.EbpAppListResp
	errRequest := a.RequestGet(ctx, headers, model2.BaseUrlApi+"/open_api/v3.0/tools/ebp/app/list/", req, &result)
	if errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}

// EbpAppExtendCreate 创建安卓分包
// https://open.oceanengine.com/labels/7/docs/1846773756545432?origin=left_nav
func (a *ToutiaoAdapter) EbpAppExtendCreateSelf(ctx context.Context, req *model2.EbpAppExtendCreateReq) (resp *model2.EbpAppExtendCreateResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result model2.EbpAppExtendCreateResp
	errRequest := a.RequestGet(ctx, headers, model2.BaseUrlApi+"/open_api/v3.0/tools/ebp/app_extend/create/", req, &result)
	if errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
