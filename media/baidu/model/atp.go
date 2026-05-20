package model

const (
	// AtpFeedServiceURL 查询定向包API端点
	AtpFeedServiceURL = "/json/feed/v1/AtpFeedService/getAtpFeed"
	// AtpFeedAddServiceURL 新增定向包API端点
	AtpFeedAddServiceURL = "/json/feed/v1/AtpFeedService/addAtpFeed"
	// AtpFeedUpdateServiceURL 更新定向包API端点
	AtpFeedUpdateServiceURL = "/json/feed/v1/AtpFeedService/updateAtpFeed"
)

// AtpFeedReq 查询定向包请求
type AtpFeedReq struct {
	AtpFeedFields []string `json:"atpFeedFields"`       // 待查询的定向包属性（必填）
	Ids           []int64  `json:"ids,omitempty"`       // 定向包ID集合 [0, 100]，不填=全部
	Key           string   `json:"key,omitempty"`       // 定向包名称查询关键字 [1, 60]
	PageNo        int      `json:"pageNo,omitempty"`    // 页码，从1开始
	PageSize      int      `json:"pageSize,omitempty"`  // 页面大小
	SortOrder     string   `json:"sortOrder,omitempty"` // 排序方式 asc=升序, desc=降序
	SortField     string   `json:"sortField,omitempty"` // 排序字段
}

// Format 格式化请求参数
func (r *AtpFeedReq) Format() {}

// Validate 校验请求参数
func (r *AtpFeedReq) Validate() error {
	return nil
}

// AtpFeedData 定向包信息数据
type AtpFeedData struct {
	AtpFeedId           int64             `json:"atpFeedId"`           // 定向包ID
	AtpFeedName         string            `json:"atpFeedName"`         // 定向包名称 [1, 60]
	AtpFeedDesc         string            `json:"atpFeedDesc"`         // 定向包描述 [1, 80]
	Ftypes              []int             `json:"ftypes"`              // 投放范围
	Subject             int               `json:"subject"`             // 推广对象
	RelatedAdgroupFeeds []int64           `json:"relatedAdgroupFeeds"` // 定向包关联单元ID [0, 100]
	Audience            map[string]string `json:"audience"`            // 定向设置
	DeliveryType        []int             `json:"deliveryType"`        // 投放场景
	MiniProgramType     int               `json:"miniProgramType"`     // 小程序子类型
}

// AtpFeedDataList 定向包信息数据列表
type AtpFeedDataList struct {
	Data []AtpFeedData `json:"data"`
}

// AtpFeedType 新增定向包对象
type AtpFeedType struct {
	AtpFeedName     string            `json:"atpFeedName"`               // 定向包名称（必填）[1, 60]
	AtpFeedDesc     string            `json:"atpFeedDesc,omitempty"`     // 定向包描述 [1, 80]
	Ftypes          []int             `json:"ftypes"`                    // 投放范围（必填）
	Subject         int               `json:"subject,omitempty"`         // 推广对象
	Audience        map[string]string `json:"audience,omitempty"`        // 定向设置
	DeliveryType    []int             `json:"deliveryType,omitempty"`    // 投放场景
	MiniProgramType int               `json:"miniProgramType,omitempty"` // 小程序子类型 3=微信小程序
}

// AtpFeedAddReq 新增定向包请求
type AtpFeedAddReq struct {
	AtpFeedTypes []AtpFeedType `json:"atpFeedTypes"` // 批量添加定向包集合 [0, 100]
}

// Format 格式化请求参数
func (r *AtpFeedAddReq) Format() {}

// Validate 校验请求参数
func (r *AtpFeedAddReq) Validate() error {
	return nil
}

// AtpFeedUpdateType 更新定向包对象
type AtpFeedUpdateType struct {
	AtpFeedId    int64             `json:"atpFeedId"`              // 定向包ID（必填）
	AtpFeedName  string            `json:"atpFeedName,omitempty"`  // 定向包名称 [1, 60]
	AtpFeedDesc  string            `json:"atpFeedDesc,omitempty"`  // 定向包描述 [1, 80]
	Ftypes       []int             `json:"ftypes,omitempty"`       // 投放范围（未绑定单元时可修改）
	Subject      int               `json:"subject,omitempty"`      // 推广对象（未绑定单元时可修改）
	Audience     map[string]string `json:"audience,omitempty"`     // 定向设置
	DeliveryType []int             `json:"deliveryType,omitempty"` // 投放场景（未绑定单元时可修改）
}

// AtpFeedUpdateReq 更新定向包请求
type AtpFeedUpdateReq struct {
	AtpFeedTypes []AtpFeedUpdateType `json:"atpFeedTypes"` // 定向包对象数组
}

// Format 格式化请求参数
func (r *AtpFeedUpdateReq) Format() {}

// Validate 校验请求参数
func (r *AtpFeedUpdateReq) Validate() error {
	return nil
}
