package kuaishou

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// ImageUploadFragment 图文件分片上传
// https://{endpoint}/api/upload/fragment?fragment_id=分片序号&upload_token=xxx
func (a *KuaishouAdapter) ImageUploadFragment(ctx context.Context, req *kuaishouModel.ImageUploadFragmentReq) (resp *kuaishouModel.ImageUploadFragmentResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}

	url := fmt.Sprintf("https://%s/api/upload/fragment?fragment_id=%d&upload_token=%s", req.Endpoint, req.FragmentId, req.UploadToken)
	headers := map[string]string{
		"Content-Type": "application/octet-stream",
	}

	respBytes, requestErr := a.Media.Client.Request(ctx, http.MethodPost, url, bytes.NewReader(req.Fragment), headers)
	if requestErr != nil {
		err = requestErr
		return
	}

	var result kuaishouModel.ImageUploadFragmentResp
	if unmarshalErr := json.Unmarshal(respBytes, &result); unmarshalErr != nil {
		err = fmt.Errorf("unmarshal response error: %s", unmarshalErr.Error())
		return
	}
	if result.Result != 1 {
		err = fmt.Errorf("image upload fragment failed: result=%d", result.Result)
		return
	}
	resp = &result
	return
}
