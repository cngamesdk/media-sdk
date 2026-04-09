package tencent

import (
	"context"
	"strconv"

	"github.com/cngamesdk/media-sdk/media/tencent/model"
	"github.com/cngamesdk/media-sdk/utils"
)

// BrandAddSelf 创建品牌形象
// https://developers.e.qq.com/v3.0/docs/api/brand/add
func (a *TencentAdapter) BrandAddSelf(ctx context.Context, req *model.BrandAddReq) (
	resp *model.BrandAddResp, err error) {
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

	fields := map[string]string{
		"account_id": strconv.FormatInt(req.AccountID, 10),
		"name":       req.Name,
	}

	var result model.BrandAddResp
	if requestErr := a.RequestPostMultipart(
		ctx,
		model.ApiUrl3+"/brand/add?"+globalQuery,
		fields,
		"brand_image_file",
		req.BrandImageFileName,
		req.BrandImageFile,
		&result,
	); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}
