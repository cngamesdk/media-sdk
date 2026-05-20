package baidu

import (
	"context"
	model2 "github.com/cngamesdk/media-sdk/media/baidu/model"
)

// GetMediaExposureSelf 查询百青藤媒体ID曝光量
// POST https://api.baidu.com/json/feed/v1/SearchFeedService/getMedias
func (a *BaiduAdapter) GetMediaExposureSelf(ctx context.Context, userName string, accessToken string, req *model2.MediaExposureReq) (resp *model2.MediaExposureDataList, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	var result model2.MediaExposureDataList
	errRequest := a.RequestPostJsonBusiness(ctx, userName, accessToken, model2.BaseUrlAPI+model2.MediaExposureServiceURL, req, &result)
	if errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}

// GetMediaPackageSelf 查询媒体包ID
// POST https://api.baidu.com/json/feed/v1/SearchFeedService/getMediaPackages
func (a *BaiduAdapter) GetMediaPackageSelf(ctx context.Context, userName string, accessToken string, req *model2.MediaPackageReq) (resp *model2.MediaPackageDataList, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	var result model2.MediaPackageDataList
	errRequest := a.RequestPostJsonBusiness(ctx, userName, accessToken, model2.BaseUrlAPI+model2.MediaPackageServiceURL, req, &result)
	if errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
