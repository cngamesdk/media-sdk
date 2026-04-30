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

// PhotoUpload 本地视频上传
func (a *KuaishouAdapter) PhotoUpload(ctx context.Context, req *kuaishouModel.PhotoUploadReq) (resp *kuaishouModel.PhotoUploadResp, err error) {
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
	if writeErr := writer.WriteField("authorId", strconv.FormatInt(req.AuthorId, 10)); writeErr != nil {
		err = fmt.Errorf("write authorId field error: %s", writeErr.Error())
		return
	}
	if writeErr := writer.WriteField("shieldBackwardSwitch", strconv.FormatBool(req.ShieldBackwardSwitch)); writeErr != nil {
		err = fmt.Errorf("write shieldBackwardSwitch field error: %s", writeErr.Error())
		return
	}
	if req.PhotoCaption != "" {
		if writeErr := writer.WriteField("photoCaption", req.PhotoCaption); writeErr != nil {
			err = fmt.Errorf("write photoCaption field error: %s", writeErr.Error())
			return
		}
	}
	if writeErr := writer.WriteField("nativePlcSwitch", strconv.FormatBool(req.NativePlcSwitch)); writeErr != nil {
		err = fmt.Errorf("write nativePlcSwitch field error: %s", writeErr.Error())
		return
	}
	if req.Sync != 0 {
		if writeErr := writer.WriteField("sync", strconv.Itoa(req.Sync)); writeErr != nil {
			err = fmt.Errorf("write sync field error: %s", writeErr.Error())
			return
		}
	}
	if len(req.PhotoTag) > 0 {
		for _, tag := range req.PhotoTag {
			if writeErr := writer.WriteField("photo_tag", tag); writeErr != nil {
				err = fmt.Errorf("write photo_tag field error: %s", writeErr.Error())
				return
			}
		}
	}
	if req.PhotoName != "" {
		if writeErr := writer.WriteField("photo_name", req.PhotoName); writeErr != nil {
			err = fmt.Errorf("write photo_name field error: %s", writeErr.Error())
			return
		}
	}
	if req.BlobStoreKey != "" {
		if writeErr := writer.WriteField("blob_store_key", req.BlobStoreKey); writeErr != nil {
			err = fmt.Errorf("write blob_store_key field error: %s", writeErr.Error())
			return
		}
	}
	if req.Signature != "" {
		if writeErr := writer.WriteField("signature", req.Signature); writeErr != nil {
			err = fmt.Errorf("write signature field error: %s", writeErr.Error())
			return
		}
	}

	part, partErr := writer.CreateFormFile("photo", req.FileName)
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

	respBytes, requestErr := a.Media.Client.Request(ctx, http.MethodPost, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/v1/photo/upload", body, headers)
	if requestErr != nil {
		err = requestErr
		return
	}

	var baseResp kuaishouModel.BaseResp
	if unmarshalErr := json.Unmarshal(respBytes, &baseResp); unmarshalErr != nil {
		err = fmt.Errorf("unmarshal response error: %s", unmarshalErr.Error())
		return
	}
	var result kuaishouModel.PhotoUploadResp
	if err = a.dealResponse(baseResp, &result); err != nil {
		return
	}
	resp = &result
	return
}
