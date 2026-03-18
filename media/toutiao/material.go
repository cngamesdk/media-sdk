package toutiao

import (
	"context"
	model2 "github.com/cngamesdk/media-sdk/media/toutiao/model"
)

// EbpVideoUploadSelf 升级版工作台上传视频
// https://open.oceanengine.com/labels/7/docs/1855448450527623?origin=left_nav
func (a *ToutiaoAdapter) EbpVideoUploadSelf(ctx context.Context, req *model2.EbpVideoUploadReq) (resp *model2.EbpVideoUploadResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result model2.EbpVideoUploadResp
	errRequest := a.RequestPostJson(ctx, headers, model2.BaseUrlApi+"/open_api/v3.0/tools/ebp/video/upload/", req, &result)
	if errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}

// FileImageAdSelf 上传广告图片
// https://open.oceanengine.com/labels/7/docs/1696710600176640?origin=left_nav
func (a *ToutiaoAdapter) FileImageAdSelf(ctx context.Context, req *model2.FileImageAdReq) (resp *model2.FileImageAdResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result model2.FileImageAdResp
	errRequest := a.RequestPostJson(ctx, headers, model2.BaseUrlApi+"/open_api/2/file/image/ad/", req, &result)
	if errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
