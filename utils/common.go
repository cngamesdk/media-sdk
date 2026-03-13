package utils

import (
	"encoding/json"
	"fmt"
	"github.com/duke-git/lancet/v2/convertor"
	"net/url"
	"strings"
)

// ConvertStructToQueryString 结构体转换为请求串
func ConvertStructToQueryString(req interface{}) (resp string, err error) {
	reqJson, reqJsonErr := json.Marshal(req)
	if reqJsonErr != nil {
		err = reqJsonErr
		return
	}
	var toMap map[string]interface{}
	if toMapErr := json.Unmarshal(reqJson, &toMap); toMapErr != nil {
		err = toMapErr
		return
	}
	var result []string
	for key, value := range toMap {
		result = append(result, fmt.Sprintf("%s=%s", key, url.QueryEscape(convertor.ToString(value))))
	}
	resp = strings.Join(result, "&")
	return
}
