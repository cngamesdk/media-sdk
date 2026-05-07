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

// ImageUploadV2 上传图片 v2
func (a *KuaishouAdapter) ImageUploadV2(ctx context.Context, req *kuaishouModel.ImageUploadV2Req) (resp *kuaishouModel.ImageUploadV2Resp, err error) {
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
	if writeErr := writer.WriteField("upload_type", strconv.Itoa(req.UploadType)); writeErr != nil {
		err = fmt.Errorf("write upload_type field error: %s", writeErr.Error())
		return
	}
	if req.Signature != "" {
		if writeErr := writer.WriteField("signature", req.Signature); writeErr != nil {
			err = fmt.Errorf("write signature field error: %s", writeErr.Error())
			return
		}
	}
	if req.Name != "" {
		if writeErr := writer.WriteField("name", req.Name); writeErr != nil {
			err = fmt.Errorf("write name field error: %s", writeErr.Error())
			return
		}
	}
	if req.Url != "" {
		if writeErr := writer.WriteField("url", req.Url); writeErr != nil {
			err = fmt.Errorf("write url field error: %s", writeErr.Error())
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

	respBytes, requestErr := a.Media.Client.Request(ctx, http.MethodPost, kuaishouModel.AdUrl+"/rest/openapi/v2/file/ad/image/upload", body, headers)
	if requestErr != nil {
		err = requestErr
		return
	}

	var baseResp kuaishouModel.BaseResp
	if unmarshalErr := json.Unmarshal(respBytes, &baseResp); unmarshalErr != nil {
		err = fmt.Errorf("unmarshal response error: %s", unmarshalErr.Error())
		return
	}
	var result kuaishouModel.ImageUploadV2Resp
	if err = a.dealResponse(baseResp, &result); err != nil {
		return
	}
	resp = &result
	return
}
