package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// NativeReportUserList 查询原生快手号列表(for 原生报表)
func (a *KuaishouAdapter) NativeReportUserList(ctx context.Context, req *kuaishouModel.NativeReportUserListReq) (resp *kuaishouModel.NativeReportUserListResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.NativeReportUserListResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/v1/native/report/user/list", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
