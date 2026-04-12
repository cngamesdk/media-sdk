package tencent

import (
	"context"
	"strconv"

	"github.com/cngamesdk/media-sdk/media/tencent/model"
	"github.com/cngamesdk/media-sdk/utils"
)

// XijingPageInteractiveAddSelf 蹊径创建互动落地页
// https://developers.e.qq.com/v3.0/docs/api/xijing_page_interactive/add
func (a *TencentAdapter) XijingPageInteractiveAddSelf(ctx context.Context, req *model.XijingPageInteractiveAddReq) (
	resp *model.XijingPageInteractiveAddResp, err error) {
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
	fields["is_auto_submit"] = strconv.Itoa(req.IsAutoSubmit)
	fields["page_type"] = req.PageType
	fields["interactive_page_type"] = req.InteractivePageType
	fields["page_title"] = req.PageTitle
	fields["page_name"] = req.PageName
	fields["mobile_app_id"] = req.MobileAppID
	if req.TransformType != "" {
		fields["transform_type"] = req.TransformType
	}
	if req.PageConfig != "" {
		fields["page_config"] = req.PageConfig
	}

	apiURL := model.ApiUrl3 + "/xijing_page_interactive/add?" + globalQuery
	var result model.XijingPageInteractiveAddResp

	if len(req.FileData) > 0 {
		if requestErr := a.RequestPostMultipart(ctx, apiURL, fields, "file", req.FileName, req.FileData, &result); requestErr != nil {
			err = requestErr
			return
		}
	} else {
		if requestErr := a.RequestPostMultipartFields(ctx, apiURL, fields, &result); requestErr != nil {
			err = requestErr
			return
		}
	}
	resp = &result
	return
}
