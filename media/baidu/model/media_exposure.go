package model

const (
	// MediaExposureServiceURL 查询百青藤媒体ID曝光量API端点
	MediaExposureServiceURL = "/json/feed/v1/SearchFeedService/getMedias"
	// MediaPackageServiceURL 查询媒体包ID API端点
	MediaPackageServiceURL = "/json/feed/v1/SearchFeedService/getMediaPackages"
)

// 媒体包类型枚举
const (
	MediaPackageTypePremium = 0 // 优选媒体包
	MediaPackageTypeVolume  = 1 // 冲量媒体包
	MediaPackageTypeBeta    = 3 // 公测媒体包
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

// MediaPackageReq 查询媒体包ID请求
type MediaPackageReq struct {
	IncludeUnavailable bool `json:"includeUnavailable"` // 是否包含失效的媒体包（必填）
}

// Format 格式化请求参数
func (r *MediaPackageReq) Format() {}

// Validate 校验请求参数
func (r *MediaPackageReq) Validate() error {
	return nil
}

// MediaPackageData 媒体包信息数据
type MediaPackageData struct {
	Id        int64   `json:"id"`        // 媒体包ID
	Name      string  `json:"name"`      // 媒体包名称
	Type      int     `json:"type"`      // 媒体包类型 0=优选, 1=冲量, 3=公测
	Tips      string  `json:"tips"`      // 媒体包说明
	Available bool    `json:"available"` // 媒体包是否有效
	Mediaids  []int64 `json:"mediaids"`  // 媒体包内包含的媒体ID集合
}

// MediaPackageDataList 媒体包信息数据列表
type MediaPackageDataList struct {
	Data []MediaPackageData `json:"data"`
}
