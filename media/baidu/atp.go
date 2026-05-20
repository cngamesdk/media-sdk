package baidu

import (
	"context"
	model2 "github.com/cngamesdk/media-sdk/media/baidu/model"
)

// GetAtpFeedSelf 查询定向包
// POST https://api.baidu.com/json/feed/v1/AtpFeedService/getAtpFeed
func (a *BaiduAdapter) GetAtpFeedSelf(ctx context.Context, userName string, accessToken string, req *model2.AtpFeedReq) (resp *model2.AtpFeedDataList, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	var result model2.AtpFeedDataList
	errRequest := a.RequestPostJsonBusiness(ctx, userName, accessToken, model2.BaseUrlAPI+model2.AtpFeedServiceURL, req, &result)
	if errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}

// BindAtpFeedSelf 定向包绑定单元
// POST https://api.baidu.com/json/feed/v1/AtpFeedService/bindAtpFeed
func (a *BaiduAdapter) BindAtpFeedSelf(ctx context.Context, userName string, accessToken string, req *model2.AtpFeedBindReq) (resp *model2.AtpBindFeedDataList, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	var result model2.AtpBindFeedDataList
	errRequest := a.RequestPostJsonBusiness(ctx, userName, accessToken, model2.BaseUrlAPI+model2.AtpFeedBindServiceURL, req, &result)
	if errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}

// DeleteAtpFeedSelf 删除定向包（可批量）
// POST https://api.baidu.com/json/feed/v1/AtpFeedService/deleteAtpFeed
func (a *BaiduAdapter) DeleteAtpFeedSelf(ctx context.Context, userName string, accessToken string, req *model2.AtpFeedDeleteReq) (resp *model2.AtpFeedDeleteDataList, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	var result model2.AtpFeedDeleteDataList
	errRequest := a.RequestPostJsonBusiness(ctx, userName, accessToken, model2.BaseUrlAPI+model2.AtpFeedDeleteServiceURL, req, &result)
	if errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}

// UpdateAtpFeedSelf 更新定向包
// POST https://api.baidu.com/json/feed/v1/AtpFeedService/updateAtpFeed
func (a *BaiduAdapter) UpdateAtpFeedSelf(ctx context.Context, userName string, accessToken string, req *model2.AtpFeedUpdateReq) (resp *model2.AtpFeedDataList, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	var result model2.AtpFeedDataList
	errRequest := a.RequestPostJsonBusiness(ctx, userName, accessToken, model2.BaseUrlAPI+model2.AtpFeedUpdateServiceURL, req, &result)
	if errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}

// AddAtpFeedSelf 新增定向包
// POST https://api.baidu.com/json/feed/v1/AtpFeedService/addAtpFeed
func (a *BaiduAdapter) AddAtpFeedSelf(ctx context.Context, userName string, accessToken string, req *model2.AtpFeedAddReq) (resp *model2.AtpFeedDataList, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	var result model2.AtpFeedDataList
	errRequest := a.RequestPostJsonBusiness(ctx, userName, accessToken, model2.BaseUrlAPI+model2.AtpFeedAddServiceURL, req, &result)
	if errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
