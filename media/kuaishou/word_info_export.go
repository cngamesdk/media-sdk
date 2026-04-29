package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// WordInfoExport 批量导出关键词
func (a *KuaishouAdapter) WordInfoExport(ctx context.Context, req *kuaishouModel.WordInfoExportReq) (resp *kuaishouModel.WordInfoExportResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.WordInfoExportResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/search/keyword/file/export", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
