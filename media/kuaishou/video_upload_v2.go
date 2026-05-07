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

// VideoUploadV2 上传视频 v2
func (a *KuaishouAdapter) VideoUploadV2(ctx context.Context, req *kuaishouModel.VideoUploadV2Req) (resp *kuaishouModel.VideoUploadV2Resp, err error) {
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
	if writeErr := writer.WriteField("signature", req.Signature); writeErr != nil {
		err = fmt.Errorf("write signature field error: %s", writeErr.Error())
		return
	}
	if writeErr := writer.WriteField("type", strconv.Itoa(req.Type)); writeErr != nil {
		err = fmt.Errorf("write type field error: %s", writeErr.Error())
		return
	}
	if writeErr := writer.WriteField("sync", strconv.Itoa(req.Sync)); writeErr != nil {
		err = fmt.Errorf("write sync field error: %s", writeErr.Error())
		return
	}
	if req.PhotoName != "" {
		if writeErr := writer.WriteField("photo_name", req.PhotoName); writeErr != nil {
			err = fmt.Errorf("write photo_name field error: %s", writeErr.Error())
			return
		}
	}
	if req.PhotoTag != "" {
		if writeErr := writer.WriteField("photo_tag", req.PhotoTag); writeErr != nil {
			err = fmt.Errorf("write photo_tag field error: %s", writeErr.Error())
			return
		}
	}
	if req.BlobStoreKey != "" {
		if writeErr := writer.WriteField("blob_store_key", req.BlobStoreKey); writeErr != nil {
			err = fmt.Errorf("write blob_store_key field error: %s", writeErr.Error())
			return
		}
	}
	if req.ShieldBackwardSwitch != nil {
		if writeErr := writer.WriteField("shield_backward_switch", strconv.FormatBool(*req.ShieldBackwardSwitch)); writeErr != nil {
			err = fmt.Errorf("write shield_backward_switch field error: %s", writeErr.Error())
			return
		}
	}
	if req.NativePlcSwitch != nil {
		if writeErr := writer.WriteField("native_plc_switch", strconv.FormatBool(*req.NativePlcSwitch)); writeErr != nil {
			err = fmt.Errorf("write native_plc_switch field error: %s", writeErr.Error())
			return
		}
	}
	if len(req.File) > 0 {
		part, partErr := writer.CreateFormFile("file", req.FileName)
		if partErr != nil {
			err = fmt.Errorf("create form file error: %s", partErr.Error())
			return
		}
		if _, writeErr := part.Write(req.File); writeErr != nil {
			err = fmt.Errorf("write file error: %s", writeErr.Error())
			return
		}
	}

	if closeErr := writer.Close(); closeErr != nil {
		err = fmt.Errorf("close multipart writer error: %s", closeErr.Error())
		return
	}

	headers := map[string]string{
		"Access-Token": req.AccessToken,
		"Content-Type": writer.FormDataContentType(),
	}

	respBytes, requestErr := a.Media.Client.Request(ctx, http.MethodPost, kuaishouModel.AdUrl+"/rest/openapi/v2/file/ad/video/upload", body, headers)
	if requestErr != nil {
		err = requestErr
		return
	}

	var baseResp kuaishouModel.BaseResp
	if unmarshalErr := json.Unmarshal(respBytes, &baseResp); unmarshalErr != nil {
		err = fmt.Errorf("unmarshal response error: %s", unmarshalErr.Error())
		return
	}
	var result kuaishouModel.VideoUploadV2Resp
	if err = a.dealResponse(baseResp, &result); err != nil {
		return
	}
	resp = &result
	return
}
