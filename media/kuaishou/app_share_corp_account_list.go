package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// AppShareCorpAccountList 获取单个主体下共享账号列表
// https://ad.e.kuaishou.com/rest/openapi/gw/dsp/appcenter/app/share/listCorpAccount
func (a *KuaishouAdapter) AppShareCorpAccountList(ctx context.Context, req *kuaishouModel.AppShareCorpAccountListReq) (resp *kuaishouModel.AppShareCorpAccountListResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.AppShareCorpAccountListResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/appcenter/app/share/listCorpAccount", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
