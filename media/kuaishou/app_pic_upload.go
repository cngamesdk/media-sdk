package kuaishou

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"strconv"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// AppPicUpload 上传图片
// https://ad.e.kuaishou.com/rest/openapi/gw/dsp/appcenter/upload/pic
func (a *KuaishouAdapter) AppPicUpload(ctx context.Context, req *kuaishouModel.AppPicUploadReq) (resp *kuaishouModel.AppPicUploadResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	if writeErr := writer.WriteField("advertiser_id", strconv.FormatInt(req.AdvertiserId, 10)); writeErr != nil {
		err = fmt.Errorf("write advertiser_id field error: %s", writeErr.Error())
		return
	}
	if writeErr := writer.WriteField("type", strconv.Itoa(req.Type)); writeErr != nil {
		err = fmt.Errorf("write type field error: %s", writeErr.Error())
		return
	}

	part, partErr := writer.CreateFormFile("file", req.FileName)
	if partErr != nil {
		err = fmt.Errorf("create form file error: %s", partErr.Error())
		return
	}
	if _, writeErr := part.Write(req.File); writeErr != nil {
		err = fmt.Errorf("write file error: %s", writeErr.Error())
		return
	}

	if closeErr := writer.Close(); closeErr != nil {
		err = fmt.Errorf("close multipart writer error: %s", closeErr.Error())
		return
	}

	headers := map[string]string{
		"Access-Token": req.AccessToken,
		"Content-Type": writer.FormDataContentType(),
	}

	respBytes, requestErr := a.Media.Client.Request(ctx, http.MethodPost, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/appcenter/upload/pic", body, headers)
	if requestErr != nil {
		err = requestErr
		return
	}

	var baseResp kuaishouModel.BaseResp
	if unmarshalErr := json.Unmarshal(respBytes, &baseResp); unmarshalErr != nil {
		err = fmt.Errorf("unmarshal response error: %s", unmarshalErr.Error())
		return
	}
	var result kuaishouModel.AppPicUploadResp
	if err = a.dealResponse(baseResp, &result); err != nil {
		return
	}
	resp = &result
	return
}
