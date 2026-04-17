package model

import "errors"

// ========== 新增线索 ==========
// https://developers.e.qq.com/v3.0/docs/api/leads/add

// 线索匹配类型枚举
const (
	LeadsMatchTypeNone    = "NONE"    // 不匹配
	LeadsMatchTypeContact = "CONTACT" // 联系方式匹配
	LeadsMatchTypeClickId = "CLICKID" // 点击id匹配
)

// 线索类型枚举
const (
	LeadsTypeForm            = "LEADS_TYPE_FORM"             // 表单
	LeadsTypeMakePhoneCall   = "LEADS_TYPE_MAKE_PHONE_CALL"  // 拨打电话
	LeadsTypePageScanCode    = "LEADS_TYPE_PAGE_SCAN_CODE"   // 扫码
	LeadsTypePromotionFollow = "LEADS_TYPE_PROMOTION_FOLLOW" // 关注
)

// 线索性别枚举
const (
	LeadsGenderTypeUnknown = "GENDER_TYPE_UNKNOWN" // 未知
	LeadsGenderTypeFemale  = "GENDER_TYPE_FEMALE"  // 女
	LeadsGenderTypeMale    = "GENDER_TYPE_MALE"    // 男
)

// 线索信息列表长度限制
const (
	LeadsInfoListMinLength = 1
	LeadsInfoListMaxLength = 50
)

// 自定义标签集合长度限制
const (
	LeadsCustomizedTagsMinLength = 1
	LeadsCustomizedTagsMaxLength = 50
	LeadsTagNameListMaxLength    = 100
)

// LeadsAddReq 新增线索请求
type LeadsAddReq struct {
	GlobalReq
	AccountID     int64           `json:"account_id"`           // 广告主账号id (必填)
	MatchType     string          `json:"match_type,omitempty"` // 线索匹配类型，不填认为是NONE
	LeadsInfoList []*LeadsAddInfo `json:"leads_info_list"`      // 导入的线索信息列表 (必填)，1-50条
}

// LeadsAddInfo 导入的线索信息
type LeadsAddInfo struct {
	OuterLeadsId             string                `json:"outer_leads_id,omitempty"`              // 外部线索id
	LeadsId                  int64                 `json:"leads_id,omitempty"`                    // 线索id
	LeadsTel                 string                `json:"leads_tel,omitempty"`                   // 手机号，1-32字节
	LeadsQq                  int64                 `json:"leads_qq,omitempty"`                    // QQ号
	LeadsWechat              string                `json:"leads_wechat,omitempty"`                // 微信号，1-64字节
	ClickId                  string                `json:"click_id,omitempty"`                    // 点击id，1-64字节
	LeadsType                string                `json:"leads_type"`                            // 线索类型 (必填)
	LeadsUserId              string                `json:"leads_user_id,omitempty"`               // 线索用户id，1-64字节
	LeadsUserType            string                `json:"leads_user_type,omitempty"`             // 线索用户类型
	LeadsUserWechatAppid     string                `json:"leads_user_wechat_appid,omitempty"`     // 线索用户的微信AppId，1-64字节
	LeadsActionTime          string                `json:"leads_action_time,omitempty"`           // 线索生成时间，格式：yyyy-MM-dd HH:mm:ss，0-32字节
	LeadsName                string                `json:"leads_name,omitempty"`                  // 姓名，0-128字节
	LeadsGender              string                `json:"leads_gender,omitempty"`                // 性别
	LeadsEmail               string                `json:"leads_email,omitempty"`                 // 邮箱，0-64字节
	LeadsArea                string                `json:"leads_area,omitempty"`                  // 所在地，0-128字节
	Bundle                   string                `json:"bundle,omitempty"`                      // 其他线索信息，0-1024字节
	Memo                     string                `json:"memo,omitempty"`                        // 备注，1-128字节
	LeadsAge                 string                `json:"leads_age,omitempty"`                   // 年龄，0-32字节
	LeadsIdNumber            string                `json:"leads_id_number,omitempty"`             // 身份证，1-24字节
	LeadsNationality         string                `json:"leads_nationality,omitempty"`           // 国籍，1-32字节
	LeadsAddress             string                `json:"leads_address,omitempty"`               // 详细地址，1-128字节
	LeadsCompany             string                `json:"leads_company,omitempty"`               // 公司，1-128字节
	LeadsProfession          string                `json:"leads_profession,omitempty"`            // 职业，1-16字节
	LeadsWorkingYears        string                `json:"leads_working_years,omitempty"`         // 工作年限，1-8字节
	LeadsPageId              string                `json:"leads_page_id,omitempty"`               // 落地页id，1-256字节
	LeadsPageName            string                `json:"leads_page_name,omitempty"`             // 落地页名称，1-256字节
	LeadsPageUrl             string                `json:"leads_page_url,omitempty"`              // 落地页链接，1-2096字节
	LeadsConvertType         string                `json:"leads_convert_type,omitempty"`          // 线索状态
	LeadsIneffectReason      string                `json:"leads_ineffect_reason,omitempty"`       // 无效原因
	OuterLeadsConvertType    string                `json:"outer_leads_convert_type,omitempty"`    // 外部线索状态，1-64字节
	OuterLeadsIneffectReason string                `json:"outer_leads_ineffect_reason,omitempty"` // 外部无效原因，1-32字节
	CustomizedTags           []*LeadsCustomizedTag `json:"customized_tags,omitempty"`             // 自定义标签集合，1-50条
}

// LeadsCustomizedTag 自定义标签
type LeadsCustomizedTag struct {
	TagGroupName string   `json:"tag_group_name,omitempty"` // 标签组名称，0-32字节
	TagNameList  []string `json:"tag_name_list,omitempty"`  // 标签集合，最大100条
}

func (p *LeadsAddReq) Format() {
	p.GlobalReq.Format()
}

// Validate 验证新增线索请求参数
func (p *LeadsAddReq) Validate() error {
	if validateErr := p.GlobalReq.Validate(); validateErr != nil {
		return validateErr
	}

	// 验证account_id
	if p.AccountID == 0 {
		return errors.New("account_id为必填")
	}

	// 验证match_type
	if p.MatchType != "" && p.MatchType != LeadsMatchTypeNone && p.MatchType != LeadsMatchTypeContact && p.MatchType != LeadsMatchTypeClickId {
		return errors.New("match_type值无效，可选值：NONE、CONTACT、CLICKID")
	}

	// 验证leads_info_list
	if len(p.LeadsInfoList) == 0 {
		return errors.New("leads_info_list为必填")
	}
	if len(p.LeadsInfoList) < LeadsInfoListMinLength || len(p.LeadsInfoList) > LeadsInfoListMaxLength {
		return errors.New("leads_info_list数组长度必须在1-50之间")
	}

	// 验证每条线索信息
	for _, info := range p.LeadsInfoList {
		if info.LeadsType == "" {
			return errors.New("leads_info_list.leads_type为必填")
		}
		if info.LeadsType != LeadsTypeForm && info.LeadsType != LeadsTypeMakePhoneCall &&
			info.LeadsType != LeadsTypePageScanCode && info.LeadsType != LeadsTypePromotionFollow {
			return errors.New("leads_info_list.leads_type值无效，可选值：LEADS_TYPE_FORM、LEADS_TYPE_MAKE_PHONE_CALL、LEADS_TYPE_PAGE_SCAN_CODE、LEADS_TYPE_PROMOTION_FOLLOW")
		}

		// 验证customized_tags
		if len(info.CustomizedTags) > LeadsCustomizedTagsMaxLength {
			return errors.New("leads_info_list.customized_tags数组长度不能超过50")
		}
		for _, tag := range info.CustomizedTags {
			if len(tag.TagNameList) > LeadsTagNameListMaxLength {
				return errors.New("leads_info_list.customized_tags.tag_name_list数组长度不能超过100")
			}
		}
	}

	return nil
}

// LeadsAddResp 新增线索响应
type LeadsAddResp struct {
	FailOuterLeadIdList []*LeadsAddFailItem    `json:"fail_outer_lead_id_list,omitempty"` // 返回失败的信息列表
	SuccessLeadIdList   []*LeadsAddSuccessItem `json:"success_lead_id_list,omitempty"`    // 返回成功的线索列表
}

// LeadsAddFailItem 失败的线索信息
type LeadsAddFailItem struct {
	Index           int    `json:"index,omitempty"`             // 线索在请求参数中的索引
	OuterLeadsId    string `json:"outer_leads_id,omitempty"`    // 外部线索id
	DetailedErrCode int    `json:"detailed_err_code,omitempty"` // 具体错误码
	DetailedErrMsg  string `json:"detailed_err_msg,omitempty"`  // 错误信息
}

// LeadsAddSuccessItem 成功的线索信息
type LeadsAddSuccessItem struct {
	Index        int    `json:"index,omitempty"`          // 线索在请求参数中的索引
	OuterLeadsId string `json:"outer_leads_id,omitempty"` // 外部线索id
	LeadsId      int64  `json:"leads_id,omitempty"`       // 线索id
}
