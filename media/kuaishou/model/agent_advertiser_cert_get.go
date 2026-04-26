package model

import "errors"

// AgentAdvertiserCertGetReq 获取账号信息和开户资质请求
type AgentAdvertiserCertGetReq struct {
	accessTokenReq
	AgentId      int64 `json:"agent_id"`      // 代理商id，必填，和access_token一致
	AdvertiserId int64 `json:"advertiser_id"` // 广告主id，必填
}

func (receiver *AgentAdvertiserCertGetReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *AgentAdvertiserCertGetReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AgentId <= 0 {
		err = errors.New("agent_id is empty")
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	return
}

// CertDetailItem 资质详情
type CertDetailItem struct {
	CertId       int64           `json:"cert_id"`                // 资质编号
	Url          string          `json:"url"`                    // 资质文件url
	CertCategory int             `json:"cert_category"`          // 资质类型：1=营业执照 2=ICP备案 3=投放承诺函 4=行业资质 5=投放资质
	FileToken    string          `json:"file_token"`             // 资质文件标识
	ExpireTime   string          `json:"expire_time"`            // 资质有效期，格式yyyy-MM-dd，空字符串表示永久有效
	ReviewStatus int             `json:"review_status"`          // 资质审核状态：0=待审核 1=审核中 2=审核通过 3=审核拒绝
	ReviewDetail string          `json:"review_detail"`          // 资质审核结果，拒绝时显示原因
	ReplaceCert  *CertDetailItem `json:"replace_cert,omitempty"` // 要替换的资质
	ExpireStatus int             `json:"expire_status"`          // 过期状态：1=有效 2=即将到期(30天内) 3=已过期
}

// AdBpCert 开户资质
type AdBpCert struct {
	CommonCertList []CertDetailItem `json:"common_cert_list"` // 当前开户资质列表
}

// ReviewDetailView 审核结果
type ReviewDetailView map[string]interface{}

// AgentAdvertiserCertGetResp 获取账号信息和开户资质响应数据（仅data部分）
type AgentAdvertiserCertGetResp struct {
	AdvertiserId         int64            `json:"advertiser_id"`          // 广告主id
	AgentId              int64            `json:"agent_id"`               // 代理商id
	Name                 string           `json:"name"`                   // 广告主快手昵称
	AdvertiserUserId     int64            `json:"advertiser_user_id"`     // 广告主userId
	ReviewStatus         int              `json:"review_status"`          // 账号审核状态：0=待审核 1=审核中 2=审核通过 3=审核拒绝 4=被封禁
	ReviewDetails        ReviewDetailView `json:"review_details"`         // 账号审核结果
	IndustryId           int64            `json:"industry_id"`            // 二级行业id
	Industry             string           `json:"industry"`               // 行业名称
	CorporationName      string           `json:"corporation_name"`       // 公司名称
	WebSite              string           `json:"web_site"`               // 企业网站
	ContactName          string           `json:"contact_name"`           // 联系人姓名
	Email                string           `json:"email"`                  // 联系人邮箱
	Phone                string           `json:"phone"`                  // 联系人手机号码
	CreateSource         int              `json:"create_source"`          // 创建来源：11=MAPI 6=BP
	ProductName          string           `json:"product_name"`           // 产品名/品牌名
	MarketingContentType int              `json:"marketing_content_type"` // 推广内容类型：1=推广内容链接 2=推广应用
	QualityControlStatus int              `json:"quality_control_status"` // 品控审核状态：0=待审核 1=审核中 2=审核通过 3=审核拒绝
	QualityControlDetail string           `json:"quality_control_detail"` // 品控审核详情
	LicenceId            string           `json:"licence_id"`             // 营业执照编号
	LicenceLocation      string           `json:"licence_location"`       // 营业执照注册地，格式"xx省-xx市"
	AdBpCert             *AdBpCert        `json:"ad_bp_cert"`             // 当前开户资质列表（cert_category 1/2/3/4）
	FrozenStatus         int              `json:"frozen_status"`          // 冻结状态：1=未冻结 2=冻结
	FrozenReason         string           `json:"frozen_reason"`          // 冻结原因
	AuthStatus           int              `json:"auth_status"`            // 真实性认证状态：0=无状态 1=待认证 2=认证中 3=认证通过 4=认证失败 5=认证失效
	AuthDetail           string           `json:"auth_detail"`            // 真实性认证明细
	Banned               bool             `json:"banned"`                 // 快手id封禁状态：true=封禁 false=正常
}
