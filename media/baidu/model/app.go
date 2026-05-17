package model

const (
	// AppFeedServiceURL 查询APP信息API端点
	AppFeedServiceURL = "/json/feed/v1/AppFeedService/getJsKpAppList"
)

// 应用审核状态枚举
const (
	AppStatusRejected = 0 // 审核拒绝
	AppStatusApproved = 1 // 审核通过
	AppStatusAuditing = 2 // 审核中
)

// AppFeedReq 查询APP信息请求（请求体为空）
type AppFeedReq struct{}

func (r *AppFeedReq) Format()         {}
func (r *AppFeedReq) Validate() error { return nil }

// AppFeedData 可投放APP信息
type AppFeedData struct {
	AppName        string `json:"appName"`        // 推广APP名称
	ApkName        string `json:"apkName"`        // 推广APP包名（仅Android）
	AppURL         string `json:"appUrl"`         // 推广APP链接（iOS=iTunes地址）
	DocID          int64  `json:"docId"`          // 推广APP docId（仅Android）
	ChannelID      int64  `json:"channelId"`      // 渠道包ID（仅Android）
	ChannelPackage string `json:"channelPackage"` // 渠道包名称（仅Android）
	AppStatus      int    `json:"appStatus"`      // 应用审核状态：0-拒绝, 1-通过, 2-审核中（仅Android）
}

// AppFeedDataList APP信息列表
type AppFeedDataList struct {
	Data []AppFeedData `json:"data"`
}
