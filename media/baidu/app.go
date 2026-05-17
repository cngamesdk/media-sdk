package baidu

import (
	"context"
	model2 "github.com/cngamesdk/media-sdk/media/baidu/model"
)

// GetAppFeedSelf 查询APP信息
// POST https://api.baidu.com/json/feed/v1/AppFeedService/getJsKpAppList
func (a *BaiduAdapter) GetAppFeedSelf(ctx context.Context, userName string, accessToken string) (resp *model2.AppFeedDataList, err error) {
	req := &model2.AppFeedReq{}
	var result model2.AppFeedDataList
	errRequest := a.RequestPostJsonBusiness(ctx, userName, accessToken, model2.BaseUrlAPI+model2.AppFeedServiceURL, req, &result)
	if errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
