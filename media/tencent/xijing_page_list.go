package tencent

import (
	"context"

	"github.com/cngamesdk/media-sdk/media/tencent/model"
)

// XijingPageListGetSelf 蹊径获取落地页列表
// https://developers.e.qq.com/v3.0/docs/api/xijing_page_list/get
func (a *TencentAdapter) XijingPageListGetSelf(ctx context.Context, req *model.XijingPageListGetReq) (
	resp *model.XijingPageListGetResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	var result model.XijingPageListGetResp
	if requestErr := a.RequestGet(ctx, nil, model.ApiUrl3+"/xijing_page_list/get", req, &result); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}
