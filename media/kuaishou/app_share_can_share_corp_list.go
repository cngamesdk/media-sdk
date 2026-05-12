package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// AppShareCanShareCorpList 获取可共享的主体列表
// https://ad.e.kuaishou.com/rest/openapi/gw/dsp/appcenter/app/canShare/corpList
func (a *KuaishouAdapter) AppShareCanShareCorpList(ctx context.Context, req *kuaishouModel.AppShareCanShareCorpListReq) (resp *kuaishouModel.AppShareCanShareCorpListResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.AppShareCanShareCorpListResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/appcenter/app/canShare/corpList", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
