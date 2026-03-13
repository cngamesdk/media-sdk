package model

import (
	"errors"
	"github.com/cngamesdk/media-sdk/model"
	"github.com/duke-git/lancet/v2/datetime"
)

// AccountReq 账户请求
type AccountReq struct {
	accessTokenReq
	AdvertiserIds []int64  `json:"advertiser_ids,omitempty"`
	Fields        []string `json:"fields,omitempty"`
}

func (receiver *AccountReq) Convert(req *model.AccountReq) {
	receiver.AccessToken = req.AccessToken
	receiver.AdvertiserIds = append(receiver.AdvertiserIds, req.AdvertiserID)
	receiver.Fields = req.Fields
	return
}

func (receiver *AccountReq) Format() {
	receiver.accessTokenReq.Format()
	return
}

func (receiver *AccountReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if len(receiver.AdvertiserIds) <= 0 {
		err = errors.New("客户ID集合为空")
		return
	}
	return
}

// AccountResp 账户响应
type AccountResp struct {
	ID                      int64  `json:"id"`                        // 客户ID
	Name                    string `json:"name"`                      // 账户名
	Role                    string `json:"role"`                      // 角色, 详见【附录-客户角色】
	Status                  string `json:"status"`                    // 状态, 详见【附录-客户状态】
	Note                    string `json:"note"`                      // 备注
	Address                 string `json:"address"`                   // 地址
	LicenseURL              string `json:"license_url"`               // 执照预览地址(链接默认1小时内有效)
	LicenseNo               string `json:"license_no"`                // 执照编号
	LicenseProvince         string `json:"license_province"`          // 执照省份
	LicenseCity             string `json:"license_city"`              // 执照城市
	Company                 string `json:"company"`                   // 公司名
	Brand                   string `json:"brand"`                     // 经营类别
	PromotionArea           string `json:"promotion_area"`            // 运营区域
	PromotionCenterProvince string `json:"promotion_center_province"` // 运营省份
	PromotionCenterCity     string `json:"promotion_center_city"`     // 运营城市
	FirstIndustryName       string `json:"first_industry_name"`       // 一级行业名称（新版）
	SecondIndustryName      string `json:"second_industry_name"`      // 二级行业名称（新版）
	Reason                  string `json:"reason"`                    // 审核拒绝原因
	CreateTime              string `json:"create_time"`               // 创建时间
}

func (receiver *AccountResp) Convert() (resp *model.AccountResp, err error) {
	resp = &model.AccountResp{}
	resp.AdvertiserID = receiver.ID
	resp.Name = receiver.Name
	resp.Role = receiver.Role
	resp.Status = receiver.Status
	myTime, myTimeErr := datetime.FormatStrToTime(receiver.CreateTime, "yyyy-mm-dd hh:mm:ss")
	if myTimeErr != nil {
		err = myTimeErr
		return
	}
	resp.CreateTime = myTime
	resp.AdExtraData = receiver
	return
}

// AccountQueryParams 账户查询参数
type EbpAdvertiserListReq struct {
	accessTokenReq
	EnterpriseOrganizationID int64          `json:"enterprise_organization_id,omitempty"` // 升级版巨量引擎工作台ID
	AccountSource            string         `json:"account_source,omitempty"`             // 账户类型，允许值：AD 巨量营销客户账号，LOCAL 本地推
	Filtering                *AccountFilter `json:"filtering,omitempty"`                  // 过滤器
	Page                     int            `json:"page,omitempty"`                       // 页码，默认1
	PageSize                 int            `json:"page_size,omitempty"`                  // 页面大小，[1,100]，默认10
}

func (receiver *EbpAdvertiserListReq) Format() {
	receiver.accessTokenReq.Format()
	if receiver.Page <= 0 {
		receiver.Page = 1
	}
	if receiver.PageSize <= 0 || receiver.PageSize > 100 {
		receiver.PageSize = 100
	}
}

func (receiver *EbpAdvertiserListReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.EnterpriseOrganizationID <= 0 {
		err = errors.New("enterprise_organization_id is empty")
		return
	}
	return
}

// AccountFilter 账户过滤器
type AccountFilter struct {
	AccountName   string `json:"account_name,omitempty"`   // 账户名称
	ActiveAccount *bool  `json:"active_account,omitempty"` // 是否是活跃账户
}

// EbpAdvertiserListResp 账户数据
type EbpAdvertiserListResp struct {
	AccountList []AccountInfo `json:"account_list,omitempty"` // 账户列表
	PageInfo    PageInfo      `json:"page_info,omitempty"`    // 分页信息
}

// AccountInfo 账户信息
type AccountInfo struct {
	AccountID   int64  `json:"account_id"`   // 账户id
	AccountType string `json:"account_type"` // 账户类型
	AccountName string `json:"account_name"` // 账户名称
}
