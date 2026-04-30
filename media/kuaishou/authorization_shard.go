package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// AuthorizationShard 共享授权
func (a *KuaishouAdapter) AuthorizationShard(ctx context.Context, req *kuaishouModel.AuthorizationShardReq) (resp *kuaishouModel.AuthorizationShardResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.AuthorizationShardResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/authorization/shard", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
