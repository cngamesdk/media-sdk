package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// OperationRecordList 账户操作记录信息查询
func (a *KuaishouAdapter) OperationRecordList(ctx context.Context, req *kuaishouModel.OperationRecordListReq) (resp *kuaishouModel.OperationRecordListResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.OperationRecordListResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/v1/tool/operation_record/list", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
