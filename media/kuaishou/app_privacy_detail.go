package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// AppPrivacyDetail 获取应用中心隐私详情
// https://ad.e.kuaishou.com/rest/openapi/gw/dsp/appcenter/privacy/detail
func (a *KuaishouAdapter) AppPrivacyDetail(ctx context.Context, req *kuaishouModel.AppPrivacyDetailReq) (resp *kuaishouModel.AppPrivacyDetailResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.AppPrivacyDetailResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/appcenter/privacy/detail", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
