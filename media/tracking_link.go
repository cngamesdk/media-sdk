package media

import (
	"fmt"
	"strings"
)

type CustomMacros map[string]string

// 重置系统默认参数
func (m CustomMacros) Reset(keySrc string, keyDst string) {
	value, ok := m[keySrc]
	if ok {
		delete(m, keySrc)
		m[keyDst] = value
	}
}

// 添加新字段
func (m CustomMacros) Add(key string, value string) {
	m[key] = value
}

// 构建请求参数
func (m CustomMacros) BuildQueryString() string {
	var container []string
	for key, value := range m {
		tempValue := strings.TrimSpace(value)
		container = append(container, fmt.Sprintf("%s=%s", strings.TrimSpace(strings.ToLower(key)), tempValue))
	}
	return strings.Join(container, "&")
}

// 构建请求URL
func (m CustomMacros) BuildUrl(url string) string {
	connectStr := "?"
	if strings.Contains(url, "?") {
		connectStr = "&"
	}
	return url + connectStr + m.BuildQueryString()
}
