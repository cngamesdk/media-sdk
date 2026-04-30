package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// AuthorizationShardResult 获取共享授权结果
func (a *KuaishouAdapter) AuthorizationShardResult(ctx context.Context, req *kuaishouModel.AuthorizationShardResultReq) (resp *kuaishouModel.AuthorizationShardResultResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.AuthorizationShardResultResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/authorization/shard/result", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
