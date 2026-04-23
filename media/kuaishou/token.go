package kuaishou

import (
	"context"
	"github.com/cngamesdk/media-sdk/media/kuaishou/model"
	model3 "github.com/cngamesdk/media-sdk/media/tencent/model"
	"github.com/cngamesdk/media-sdk/utils"
)

func (a *KuaishouAdapter) AuthSelf(ctx context.Context, req *model.AuthReq) (resp *model.AuthResp, err error) {
	_ = ctx
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	convertResult, convertErr := utils.ConvertStructToQueryString(req)
	if convertErr != nil {
		err = convertErr
		return
	}
	result := model.AuthResp(model3.DevelopersUrl + "/tools/authorize?" + convertResult)
	resp = &result
	return
}
