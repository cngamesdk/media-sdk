package model

import "errors"

// CertParam 资质参数
type CertParam struct {
	CertId       int64  `json:"cert_id"`                 // 资质id，新增时传0，非0为已有资质id
	CertCategory int    `json:"cert_category,omitempty"` // 资质类型：1=营业执照 2=ICP备案 3=投放承诺函 4=行业资质
	FileToken    string `json:"file_token,omitempty"`    // 资质文件标识，上传资质接口返回的file_token
	ExpireTime   string `json:"expire_time,omitempty"`   // 资质有效期，格式yyyy-MM-dd，空字符串表示长期有效
}

// AgentAdvertiserCertSubmitReq 创建或更新账户信息和开户资质请求
type AgentAdvertiserCertSubmitReq struct {
	accessTokenReq
	AdvertiserId         int64       `json:"advertiser_id"`          // 广告主id，必填
	AgentId              int64       `json:"agent_id"`               // 代理商id，必填
	IndustryId           int64       `json:"industry_id"`            // 二级行业id，必填
	CorporationName      string      `json:"corporation_name"`       // 公司名称，必填，不超过50字符
	WebSite              string      `json:"web_site"`               // 公司网站，必填，http://或https://开头
	ProductName          string      `json:"product_name"`           // 产品名/品牌名，必填
	MarketingContentType int         `json:"marketing_content_type"` // 推广内容类型，必填：1=推广内容链接 2=推广应用
	LicenceId            string      `json:"licence_id"`             // 营业执照编号，必填，只能有数字和字母
	LicenceLocation      string      `json:"licence_location"`       // 营业执照注册地，必填，格式"xx省-xx市"
	CertList             []CertParam `json:"cert_list"`              // 开户资质列表，必填
	ContactName          string      `json:"contact_name,omitempty"` // 联系人姓名，中文或英文，不超过50字符
	Email                string      `json:"email,omitempty"`        // 邮箱
	Phone                string      `json:"phone,omitempty"`        // 手机号码，11位
}

func (receiver *AgentAdvertiserCertSubmitReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *AgentAdvertiserCertSubmitReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	if receiver.AgentId <= 0 {
		err = errors.New("agent_id is empty")
		return
	}
	if receiver.IndustryId <= 0 {
		err = errors.New("industry_id is empty")
		return
	}
	if len(receiver.CorporationName) == 0 {
		err = errors.New("corporation_name is empty")
		return
	}
	if len(receiver.WebSite) == 0 {
		err = errors.New("web_site is empty")
		return
	}
	if len(receiver.ProductName) == 0 {
		err = errors.New("product_name is empty")
		return
	}
	if receiver.MarketingContentType <= 0 {
		err = errors.New("marketing_content_type is empty")
		return
	}
	if len(receiver.LicenceId) == 0 {
		err = errors.New("licence_id is empty")
		return
	}
	if len(receiver.LicenceLocation) == 0 {
		err = errors.New("licence_location is empty")
		return
	}
	if len(receiver.CertList) == 0 {
		err = errors.New("cert_list is empty")
		return
	}
	return
}

// AgentAdvertiserCertSubmitResp 创建或更新账户信息和开户资质响应数据（仅data部分）
type AgentAdvertiserCertSubmitResp struct {
	Result bool `json:"result"` // 是否成功
}
