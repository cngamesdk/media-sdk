package tencent

import (
	"context"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
	"github.com/cngamesdk/media-sdk/utils"
)

// VideoChannelLeadsDataGetSelf 获取线索数据
// https://developers.e.qq.com/v3.0/docs/api/video_channel_leads_data/get
func (a *TencentAdapter) VideoChannelLeadsDataGetSelf(ctx context.Context, req *model.VideoChannelLeadsDataGetReq) (
	resp *model.VideoChannelLeadsDataGetResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	globalQuery, globalQueryErr := utils.ConvertStructToQueryString(req.GlobalReq)
	if globalQueryErr != nil {
		err = globalQueryErr
		return
	}
	req.GlobalReq.Clear()
	headers := make(model.Headers)
	headers.Json()
	var result model.VideoChannelLeadsDataGetResp
	if requestErr := a.RequestPostJson(ctx, headers, model.ApiUrl3+"/video_channel_leads_data/get?"+globalQuery, req, &result); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}
