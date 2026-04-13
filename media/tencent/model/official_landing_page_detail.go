package model

import "errors"

// ========== 官方落地页-获取落地页详情 ==========
// https://developers.e.qq.com/v3.0/docs/api/official_landing_page_detail/get

// OfficialLandingPageDetailGetReq 官方落地页获取落地页详情请求（GET）
type OfficialLandingPageDetailGetReq struct {
	GlobalReq
	AccountId    int64 `json:"account_id" form:"account_id"`       // 广告主帐号 id (必填)
	PageId       int64 `json:"page_id" form:"page_id"`             // 落地页 id (必填)
	ProtoVersion int   `json:"proto_version" form:"proto_version"` // 协议版本，0 或 1，默认 0
}

func (r *OfficialLandingPageDetailGetReq) Format() {
	r.GlobalReq.Format()
}

// Validate 验证获取落地页详情请求参数
func (r *OfficialLandingPageDetailGetReq) Validate() error {
	if r.AccountId == 0 {
		return errors.New("account_id为必填")
	}
	if r.PageId == 0 {
		return errors.New("page_id为必填")
	}
	if r.ProtoVersion != 0 && r.ProtoVersion != 1 {
		return errors.New("proto_version须为0或1")
	}
	return r.GlobalReq.Validate()
}

// OfficialLandingPageDetailConfig 落地页详情中的页面配置
// 与创建接口的 OfficialLandingPageConfig 结构基本相同，
// 区别在于此处 share_thumb_url 为已上传图片的 URL，而创建接口使用 share_thumburl_material_id（素材ID）
type OfficialLandingPageDetailConfig struct {
	PageName            string `json:"page_name"`             // 落地页名称（管理用）
	PageTitle           string `json:"page_title"`            // 落地页标题（展示用）
	IosAppId            string `json:"ios_app_id"`            // ios App ID
	AndroidAppId        string `json:"android_app_id"`        // android App ID
	EnableAndroidMarket bool   `json:"enable_android_market"` // 是否开启自动跳转应用市场
	ShareTitle          string `json:"share_title"`           // 分享标题
	ShareDescription    string `json:"share_description"`     // 分享描述
	ShareThumbUrl       string `json:"share_thumb_url"`       // 分享缩略图 URL（返回值为图片 URL，而非素材 ID）
	BgColor             string `json:"bg_color"`              // 页面背景颜色
}

// OfficialLandingPageDetailRespData 获取落地页详情响应数据
type OfficialLandingPageDetailRespData struct {
	PageId        int64                            `json:"page_id"`         // 落地页 id
	LandingPageId int                              `json:"landing_page_id"` // 落地页 id（旧字段，与 page_id 对应）
	PageConfig    *OfficialLandingPageDetailConfig `json:"page_config"`     // 落地页配置
	PageElements  []*OfficialLandingPageElement    `json:"page_elements"`   // 落地页组件列表，复用创建接口的组件结构
}

// OfficialLandingPageDetailGetResp 获取落地页详情响应
type OfficialLandingPageDetailGetResp struct {
	Code      int                                `json:"code"`
	Message   string                             `json:"message"`
	MessageCn string                             `json:"message_cn"`
	Data      *OfficialLandingPageDetailRespData `json:"data"`
}
