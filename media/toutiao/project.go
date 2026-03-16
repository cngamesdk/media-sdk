package toutiao

import (
	"context"
	model2 "github.com/cngamesdk/media-sdk/media/toutiao/model"
)

// ProjectCreateSelf 创建项目
// https://open.oceanengine.com/labels/7/docs/1740868093375503?origin=left_nav
func (a *ToutiaoAdapter) ProjectCreateSelf(ctx context.Context, req *model2.ProjectCreateReq) (resp *model2.ProjectCreateResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result model2.ProjectCreateResp
	errRequest := a.RequestPostJson(ctx, headers, model2.BaseUrlApi+"/open_api/v3.0/project/create/", req, &result)
	if errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}

// ProjectUpdateSelf 更新项目
// https://open.oceanengine.com/labels/7/docs/1740936504522831?origin=left_nav
func (a *ToutiaoAdapter) ProjectUpdateSelf(ctx context.Context, req *model2.ProjectUpdateReq) (resp *model2.ProjectUpdateResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result model2.ProjectUpdateResp
	errRequest := a.RequestPostJson(ctx, headers, model2.BaseUrlApi+"/open_api/v3.0/project/update/", req, &result)
	if errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
