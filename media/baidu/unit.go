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

// AddAdgroupFeedSelf 新建单元
// POST https://api.baidu.com/json/feed/v1/AdgroupFeedService/addAdgroupFeed
func (a *BaiduAdapter) AddAdgroupFeedSelf(ctx context.Context, userName string, accessToken string, req *model2.AdgroupFeedAddReq) (resp *model2.AdgroupFeedDataList, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	var result model2.AdgroupFeedDataList
	errRequest := a.RequestPostJsonBusiness(ctx, userName, accessToken, model2.BaseUrlAPI+model2.AdgroupFeedAddServiceURL, req, &result)
	if errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}

// UpdateAdgroupFeedSelf 更新单元
// POST https://api.baidu.com/json/feed/v1/AdgroupFeedService/updateAdgroupFeed
func (a *BaiduAdapter) UpdateAdgroupFeedSelf(ctx context.Context, userName string, accessToken string, req *model2.AdgroupFeedUpdateReq) (resp *model2.AdgroupFeedDataList, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	var result model2.AdgroupFeedDataList
	errRequest := a.RequestPostJsonBusiness(ctx, userName, accessToken, model2.BaseUrlAPI+model2.AdgroupFeedUpdateServiceURL, req, &result)
	if errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}

// DeleteAdgroupFeedSelf 删除单元（可批量）
// POST https://api.baidu.com/json/feed/v1/AdgroupFeedService/deleteAdgroupFeed
func (a *BaiduAdapter) DeleteAdgroupFeedSelf(ctx context.Context, userName string, accessToken string, req *model2.AdgroupFeedDeleteReq) (resp *model2.AdgroupFeedDeleteDataList, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	var result model2.AdgroupFeedDeleteDataList
	errRequest := a.RequestPostJsonBusiness(ctx, userName, accessToken, model2.BaseUrlAPI+model2.AdgroupFeedDeleteServiceURL, req, &result)
	if errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}

// GetDpaAdgroupFeedSelf 查询商品推广单元
// POST https://api.baidu.com/json/feed/v1/DpaAdgroupFeedService/getAdgroupFeed
func (a *BaiduAdapter) GetDpaAdgroupFeedSelf(ctx context.Context, userName string, accessToken string, req *model2.DpaAdgroupFeedReq) (resp *model2.DpaAdgroupFeedDataList, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	var result model2.DpaAdgroupFeedDataList
	errRequest := a.RequestPostJsonBusiness(ctx, userName, accessToken, model2.BaseUrlAPI+model2.DpaAdgroupFeedServiceURL, req, &result)
	if errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
