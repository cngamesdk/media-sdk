package kuaishou

import (
	"context"
	"fmt"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// DmpPopulationDelete 人群包删除
// https://ad.e.kuaishou.com/rest/openapi/v1/dmp/population/delete
func (a *KuaishouAdapter) DmpPopulationDelete(ctx context.Context, req *kuaishouModel.DmpPopulationDeleteReq) (resp *kuaishouModel.DmpPopulationDeleteResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()

	// data 为字符串，单独解析，不走 dealResponse
	var baseResp struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Data    string `json:"data"`
	}
	if errRequest := a.Media.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/v1/dmp/population/delete", req, &baseResp); errRequest != nil {
		err = errRequest
		return
	}
	if baseResp.Code != 0 {
		err = fmt.Errorf("kuaishou api error: code=%d, message:%s", baseResp.Code, baseResp.Message)
		return
	}
	resp = &kuaishouModel.DmpPopulationDeleteResp{
		DeleteMsg: baseResp.Data,
	}
	return
}
