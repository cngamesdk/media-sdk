package tencent

import (
	"context"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
)

// ComponentsGetSelf 获取创意组件
// https://developers.e.qq.com/v3.0/docs/api/components/get
func (a *TencentAdapter) ComponentsGetSelf(ctx context.Context, req *model.ComponentsGetReq) (
	resp *model.ComponentsGetResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	var result model.ComponentsGetResp
	if requestErr := a.RequestGet(ctx, nil, model.ApiUrl3+"/components/get", req, &result); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}
