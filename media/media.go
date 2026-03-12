package media

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/utils"
	"github.com/duke-git/lancet/v2/netutil"
	"github.com/spf13/cast"
	"io"
	"net/http"
	"reflect"
	"strings"
)

type Media struct {
	Config *config.Config
	Client *utils.HTTPClient
}

// RequestPostJson 发送POST请求
func (a *Media) RequestPostJson(ctx context.Context, headers map[string]string, url string, body interface{}, result interface{}) error {
	var ioReader io.Reader
	if body != nil {
		bodyData, errJson := json.Marshal(body)
		if errJson != nil {
			return errJson
		}
		ioReader = bytes.NewReader(bodyData)
	}
	if headers == nil {
		headers = make(map[string]string)
	}

	if _, ok := headers["Content-Type"]; !ok {
		headers["Content-Type"] = "application/json"
	}

	resp, err := a.Client.Request(ctx, http.MethodPost, url, ioReader, headers)
	if err != nil {
		return err
	}

	return json.Unmarshal(resp, result)
}

// RequestGet 发送GET请求
func (a *Media) RequestGet(ctx context.Context, headers map[string]string, url string, data interface{}, result interface{}) (err error) {
	var query string
	if data != nil {
		kind := reflect.TypeOf(data).Kind()
		switch kind {
		case reflect.String:
			query = cast.ToString(data)
			break
		case reflect.Map:
			mapData, ok := data.(map[string]interface{})
			if !ok {
				err = errors.New("data格式为map[string]interface{}")
				return
			}
			query = netutil.ConvertMapToQueryString(mapData)
			break
		case reflect.Struct:
		case reflect.Ptr:
			query, err = utils.ConvertStructToQueryString(data)
			if err != nil {
				return
			}
			break
		default:
			err = errors.New("data类型不支持." + kind.String())
			return
		}
	}
	if len(query) > 0 {
		if strings.Index(url, "?") < 0 {
			url += "?"
		}
		url += query
	}
	response, errResponse := a.Client.Request(ctx, http.MethodGet, url, nil, headers)
	if errResponse != nil {
		return errResponse
	}

	return json.Unmarshal(response, result)
}
