package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// AppShareCorpList 获取应用已共享主体列表
// https://ad.e.kuaishou.com/rest/openapi/gw/dsp/appcenter/app/shareCorpList
func (a *KuaishouAdapter) AppShareCorpList(ctx context.Context, req *kuaishouModel.AppShareCorpListReq) (resp *kuaishouModel.AppShareCorpListResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.AppShareCorpListResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/appcenter/app/shareCorpList", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
