package baidu

import (
	"context"
	model2 "github.com/cngamesdk/media-sdk/media/baidu/model"
)

// GetCampaignFeedSelf 查询计划
// POST https://api.baidu.com/json/feed/v1/CampaignFeedService/getCampaignFeed
func (a *BaiduAdapter) GetCampaignFeedSelf(ctx context.Context, userName string, accessToken string, req *model2.CampaignFeedReq) (resp *model2.CampaignFeedDataList, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	var result model2.CampaignFeedDataList
	errRequest := a.RequestPostJsonBusiness(ctx, userName, accessToken, model2.BaseUrlAPI+model2.CampaignFeedServiceURL, req, &result)
	if errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}

// AddCampaignFeedSelf 新建计划
// POST https://api.baidu.com/json/feed/v1/CampaignFeedService/addCampaignFeed
func (a *BaiduAdapter) AddCampaignFeedSelf(ctx context.Context, userName string, accessToken string, req *model2.CampaignFeedAddReq) (resp *model2.CampaignFeedDataList, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	var result model2.CampaignFeedDataList
	errRequest := a.RequestPostJsonBusiness(ctx, userName, accessToken, model2.BaseUrlAPI+model2.CampaignFeedAddServiceURL, req, &result)
	if errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}

// UpdateCampaignFeedSelf 更新计划
// POST https://api.baidu.com/json/feed/v1/CampaignFeedService/updateCampaignFeed
func (a *BaiduAdapter) UpdateCampaignFeedSelf(ctx context.Context, userName string, accessToken string, req *model2.CampaignFeedUpdateReq) (resp *model2.CampaignFeedDataList, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	var result model2.CampaignFeedDataList
	errRequest := a.RequestPostJsonBusiness(ctx, userName, accessToken, model2.BaseUrlAPI+model2.CampaignFeedUpdateServiceURL, req, &result)
	if errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}

// DeleteCampaignFeedSelf 删除计划（可批量）
// POST https://api.baidu.com/json/feed/v1/CampaignFeedService/deleteCampaignFeed
func (a *BaiduAdapter) DeleteCampaignFeedSelf(ctx context.Context, userName string, accessToken string, req *model2.CampaignFeedDeleteReq) (resp *model2.CampaignFeedDeleteDataList, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	var result model2.CampaignFeedDeleteDataList
	errRequest := a.RequestPostJsonBusiness(ctx, userName, accessToken, model2.BaseUrlAPI+model2.CampaignFeedDeleteServiceURL, req, &result)
	if errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
