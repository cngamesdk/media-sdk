package kuaishou

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// ImageUploadComplete 图合并分片
// https://{endpoint}/api/upload/complete?fragment_count=分片总数&upload_token=xxxx
func (a *KuaishouAdapter) ImageUploadComplete(ctx context.Context, req *kuaishouModel.ImageUploadCompleteReq) (resp *kuaishouModel.ImageUploadCompleteResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}

	url := fmt.Sprintf("https://%s/api/upload/complete?fragment_count=%d&upload_token=%s", req.Endpoint, req.FragmentCount, req.UploadToken)
	headers := map[string]string{
		"Content-Type": "application/json",
	}

	respBytes, requestErr := a.Media.Client.Request(ctx, http.MethodPost, url, nil, headers)
	if requestErr != nil {
		err = requestErr
		return
	}

	var result kuaishouModel.ImageUploadCompleteResp
	if unmarshalErr := json.Unmarshal(respBytes, &result); unmarshalErr != nil {
		err = fmt.Errorf("unmarshal response error: %s", unmarshalErr.Error())
		return
	}
	if result.Result != 1 {
		err = fmt.Errorf("image upload complete failed: result=%d", result.Result)
		return
	}
	resp = &result
	return
}
