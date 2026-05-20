package baidu

import (
	"context"
	model2 "github.com/cngamesdk/media-sdk/media/baidu/model"
)

// GetTransTraceSelf 查询转化追踪
// POST https://api.baidu.com/json/feed/v1/SearchFeedService/getOcpcTransFeed
func (a *BaiduAdapter) GetTransTraceSelf(ctx context.Context, userName string, accessToken string, req *model2.TransTraceReq) (resp *model2.TransTraceDataList, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	var result model2.TransTraceDataList
	errRequest := a.RequestPostJsonBusiness(ctx, userName, accessToken, model2.BaseUrlAPI+model2.TransTraceServiceURL, req, &result)
	if errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}

// AddTransTraceSelf 新增转化追踪
// POST https://api.baidu.com/json/feed/v1/OcpcTransFeedService/addOcpcTransFeed
func (a *BaiduAdapter) AddTransTraceSelf(ctx context.Context, userName string, accessToken string, req *model2.TransTraceAddReq) (resp *model2.TransTraceDataList, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	var result model2.TransTraceDataList
	errRequest := a.RequestPostJsonBusiness(ctx, userName, accessToken, model2.BaseUrlAPI+model2.TransTraceAddServiceURL, req, &result)
	if errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
