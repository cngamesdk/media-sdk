package baidu

import (
	"context"
	model2 "github.com/cngamesdk/media-sdk/media/baidu/model"
)

// GetAdgroupFeedSelf 查询单元
// POST https://api.baidu.com/json/feed/v1/AdgroupFeedService/getAdgroupFeed
func (a *BaiduAdapter) GetAdgroupFeedSelf(ctx context.Context, userName string, accessToken string, req *model2.AdgroupFeedReq) (resp *model2.AdgroupFeedDataList, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	var result model2.AdgroupFeedDataList
	errRequest := a.RequestPostJsonBusiness(ctx, userName, accessToken, model2.BaseUrlAPI+model2.AdgroupFeedServiceURL, req, &result)
	if errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
