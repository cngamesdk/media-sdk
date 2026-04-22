package tencent

import (
	"context"
	"strconv"

	"github.com/cngamesdk/media-sdk/media/tencent/model"
	"github.com/cngamesdk/media-sdk/utils"
)

// IllegalComplaintAdd 新增广告主违规申述
// https://developers.e.qq.com/v3.0/docs/api/illegal_complaint/add
func (a *TencentAdapter) IllegalComplaintAdd(ctx context.Context, req *model.IllegalComplaintAddReq) (
	resp *model.IllegalComplaintAddResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	globalQuery, globalQueryErr := utils.ConvertStructToQueryString(req.GlobalReq)
	if globalQueryErr != nil {
		err = globalQueryErr
		return
	}
	req.GlobalReq.Clear()

	fields := make(map[string]string)
	fields["account_id"] = strconv.FormatInt(req.AccountID, 10)
	fields["illegal_order_id"] = req.IllegalOrderID
	fields["complaint_reason"] = req.ComplaintReason

	var result model.IllegalComplaintAddResp
	if requestErr := a.RequestPostMultipart(
		ctx,
		model.ApiUrl3+"/illegal_complaint/add?"+globalQuery,
		fields,
		"file",
		req.FileName,
		req.File,
		&result,
	); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}
