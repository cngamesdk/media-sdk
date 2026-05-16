package baidu

import (
	"context"
	model2 "github.com/cngamesdk/media-sdk/media/baidu/model"
)

// GetAccountFeedSelf 查询信息流账户信息
// POST https://api.baidu.com/json/feed/v1/AccountFeedService/getAccountFeed
func (a *BaiduAdapter) GetAccountFeedSelf(ctx context.Context, userName string, accessToken string, req *model2.AccountFeedReq) (resp *model2.AccountFeedDataList, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	var result model2.AccountFeedDataList
	errRequest := a.RequestPostJsonBusiness(ctx, userName, accessToken, model2.BaseUrlAPI+model2.AccountFeedServiceURL, req, &result)
	if errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}

// UpdateAccountFeedSelf 更新信息流账户信息
// POST https://api.baidu.com/json/feed/v1/AccountFeedService/updateAccountFeed
func (a *BaiduAdapter) UpdateAccountFeedSelf(ctx context.Context, userName string, accessToken string, req *model2.AccountFeedUpdateReq) (resp *model2.AccountFeedUpdateDataList, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	var result model2.AccountFeedUpdateDataList
	errRequest := a.RequestPostJsonBusiness(ctx, userName, accessToken, model2.BaseUrlAPI+model2.AccountFeedUpdateServiceURL, req, &result)
	if errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
