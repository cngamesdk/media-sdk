package toutiao

import (
	"context"
	model2 "github.com/cngamesdk/media-sdk/media/toutiao/model"
)

// PromotionCreateSelf 创建项目
// https://open.oceanengine.com/labels/7/docs/1740946299496459?origin=left_nav
func (a *ToutiaoAdapter) PromotionCreateSelf(ctx context.Context, req *model2.PromotionCreateReq) (resp *model2.PromotionCreateResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result model2.PromotionCreateResp
	errRequest := a.RequestPostJson(ctx, headers, model2.BaseUrlApi+"/open_api/v3.0/promotion/create/", req, &result)
	if errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}

// PromotionUpdateSelf 修改单元
// https://open.oceanengine.com/labels/7/docs/1740952287987719?origin=left_nav
func (a *ToutiaoAdapter) PromotionUpdateSelf(ctx context.Context, req *model2.PromotionUpdateReq) (resp *model2.PromotionUpdateResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result model2.PromotionUpdateResp
	errRequest := a.RequestPostJson(ctx, headers, model2.BaseUrlApi+"/open_api/v3.0/promotion/update/", req, &result)
	if errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}

// PromotionListSelf 修改单元
// https://open.oceanengine.com/labels/7/docs/1741028841006095?origin=left_nav
func (a *ToutiaoAdapter) PromotionListSelf(ctx context.Context, req *model2.PromotionListReq) (resp *model2.PromotionListResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result model2.PromotionListResp
	errRequest := a.RequestGet(ctx, headers, model2.BaseUrlApi+"/open_api/v3.0/promotion/list/", req, &result)
	if errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
