package utils

import (
	"fmt"
	"github.com/duke-git/lancet/v2/convertor"
	"net/url"
	"strings"
)

// ConvertStructToQueryString 结构体转换为请求串
func ConvertStructToQueryString(req interface{}) (resp string, err error) {
	reqMap, reqMapErr := convertor.StructToMap(req)
	if reqMapErr != nil {
		err = reqMapErr
		return
	}
	var result []string
	for key, value := range reqMap {
		result = append(result, fmt.Sprintf("%s=%s", key, url.QueryEscape(convertor.ToString(value))))
	}
	resp = strings.Join(result, "&")
	return
}
