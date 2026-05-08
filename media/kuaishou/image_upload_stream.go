package kuaishou

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// ImageUploadStream 图文件流式上传
// https://{endpoint}/api/upload/multipart?upload_token=xxx
func (a *KuaishouAdapter) ImageUploadStream(ctx context.Context, req *kuaishouModel.ImageUploadStreamReq) (resp *kuaishouModel.ImageUploadStreamResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

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

	url := fmt.Sprintf("https://%s/api/upload/multipart?upload_token=%s", req.Endpoint, req.UploadToken)
	headers := map[string]string{
		"Content-Type": writer.FormDataContentType(),
	}

	respBytes, requestErr := a.Media.Client.Request(ctx, http.MethodPost, url, body, headers)
	if requestErr != nil {
		err = requestErr
		return
	}

	var result kuaishouModel.ImageUploadStreamResp
	if unmarshalErr := json.Unmarshal(respBytes, &result); unmarshalErr != nil {
		err = fmt.Errorf("unmarshal response error: %s", unmarshalErr.Error())
		return
	}
	if result.Result != 1 {
		err = fmt.Errorf("image upload stream failed: result=%d", result.Result)
		return
	}
	resp = &result
	return
}
