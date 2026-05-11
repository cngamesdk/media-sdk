package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// AppShareList 获取应用已共享账号列表
// https://ad.e.kuaishou.com/rest/openapi/gw/dsp/appcenter/app/share/list
func (a *KuaishouAdapter) AppShareList(ctx context.Context, req *kuaishouModel.AppShareListReq) (resp *kuaishouModel.AppShareListResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.AppShareListResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/appcenter/app/share/list", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
