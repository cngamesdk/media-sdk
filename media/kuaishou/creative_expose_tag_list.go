package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// CreativeExposeTagList 查询创意推荐理由
func (a *KuaishouAdapter) CreativeExposeTagList(ctx context.Context, req *kuaishouModel.CreativeExposeTagListReq) (resp *kuaishouModel.CreativeExposeTagListResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.CreativeExposeTagListResp
	if errRequest := a.RequestGet(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/v1/tool/expose_tags/list", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
