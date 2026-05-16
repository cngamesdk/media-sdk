package baidu

import (
	"context"
	model2 "github.com/cngamesdk/media-sdk/media/baidu/model"
)

// GetProjectFeedSelf 查询项目
// POST https://api.baidu.com/json/sms/service/ProjectFeedService/getProjectFeed
func (a *BaiduAdapter) GetProjectFeedSelf(ctx context.Context, userName string, accessToken string, req *model2.ProjectFeedReq) (resp *model2.ProjectFeedDataList, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	var result model2.ProjectFeedDataList
	errRequest := a.RequestPostJsonBusiness(ctx, userName, accessToken, model2.BaseUrlAPI+model2.ProjectFeedServiceURL, req, &result)
	if errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
