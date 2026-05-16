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

// AddProjectFeedSelf 新建项目
// POST https://api.baidu.com/json/sms/service/ProjectFeedService/addProjectFeed
func (a *BaiduAdapter) AddProjectFeedSelf(ctx context.Context, userName string, accessToken string, req *model2.ProjectFeedAddReq) (resp *model2.ProjectFeedDataList, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	var result model2.ProjectFeedDataList
	errRequest := a.RequestPostJsonBusiness(ctx, userName, accessToken, model2.BaseUrlAPI+model2.ProjectFeedAddServiceURL, req, &result)
	if errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}

// UpdateProjectFeedSelf 更新项目
// POST https://api.baidu.com/json/sms/service/ProjectFeedService/updateProjectFeed
func (a *BaiduAdapter) UpdateProjectFeedSelf(ctx context.Context, userName string, accessToken string, req *model2.ProjectFeedUpdateReq) (resp *model2.ProjectFeedDataList, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	var result model2.ProjectFeedDataList
	errRequest := a.RequestPostJsonBusiness(ctx, userName, accessToken, model2.BaseUrlAPI+model2.ProjectFeedUpdateServiceURL, req, &result)
	if errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}

// DeleteProjectFeedSelf 删除项目（可批量）
// POST https://api.baidu.com/json/sms/service/ProjectFeedService/deleteProjectFeed
func (a *BaiduAdapter) DeleteProjectFeedSelf(ctx context.Context, userName string, accessToken string, req *model2.ProjectFeedDeleteReq) (resp *model2.ProjectFeedDeleteDataList, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	var result model2.ProjectFeedDeleteDataList
	errRequest := a.RequestPostJsonBusiness(ctx, userName, accessToken, model2.BaseUrlAPI+model2.ProjectFeedDeleteServiceURL, req, &result)
	if errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
