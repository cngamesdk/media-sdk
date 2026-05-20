package model

const (
	// MediaExposureServiceURL 查询百青藤媒体ID曝光量API端点
	MediaExposureServiceURL = "/json/feed/v1/SearchFeedService/getMedias"
)

// MediaExposureReq 查询百青藤媒体ID曝光量请求
type MediaExposureReq struct {
	NewMediaids []int64 `json:"newMediaids"` // 要查询的媒体ID集合（必填）
}

// Format 格式化请求参数
func (r *MediaExposureReq) Format() {}

// Validate 校验请求参数
func (r *MediaExposureReq) Validate() error {
	return nil
}

// MediaExposureData 百青藤媒体曝光量数据
type MediaExposureData struct {
	Id   int64  `json:"id"`   // 媒体ID
	Name string `json:"name"` // 媒体名称
	Pv   int    `json:"pv"`   // 媒体流量（单位：万），精确到百万。0表示曝光量小于50万
}

// MediaExposureDataList 百青藤媒体曝光量数据列表
type MediaExposureDataList struct {
	Data []MediaExposureData `json:"data"`
}
