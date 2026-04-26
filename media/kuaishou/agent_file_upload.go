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

// AgentFileUploadSelf 上传资质文件
func (a *KuaishouAdapter) AgentFileUploadSelf(ctx context.Context, req *kuaishouModel.AgentFileUploadReq) (resp *kuaishouModel.AgentFileUploadResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}

	// 构建 multipart body
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// 添加 agent_id 字段
	if writeErr := writer.WriteField("agent_id", strconv.FormatInt(req.AgentId, 10)); writeErr != nil {
		err = fmt.Errorf("write agent_id field error: %s", writeErr.Error())
		return
	}

	// 添加 file 文件字段
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

	// 构建 headers
	headers := map[string]string{
		"Access-Token": req.AccessToken,
		"Content-Type": writer.FormDataContentType(),
	}

	// 发送请求
	respBytes, requestErr := a.Media.Client.Request(ctx, http.MethodPost, kuaishouModel.AdUrl+"/rest/openapi/v1/agent/file/upload", body, headers)
	if requestErr != nil {
		err = requestErr
		return
	}

	// 解析响应
	var baseResp kuaishouModel.BaseResp
	if unmarshalErr := json.Unmarshal(respBytes, &baseResp); unmarshalErr != nil {
		err = fmt.Errorf("unmarshal response error: %s", unmarshalErr.Error())
		return
	}
	if err = a.dealResponse(baseResp, &resp); err != nil {
		return
	}
	return
}
