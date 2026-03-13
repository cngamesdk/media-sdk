package model

import (
	"errors"
	"strings"
)

const (
	BaseUrlOpen = "https://open.oceanengine.com"
	BaseUrlApi  = "https://api.oceanengine.com"
	BaseUrlAd   = "https://ad.oceanengine.com"
)

type AccountRole string

const (
	RoleAdvertiser                   AccountRole = "ADVERTISER"                           // 客户
	RoleCustomerAdmin                AccountRole = "CUSTOMER_ADMIN"                       // 普通版工作台-管理员
	RoleCustomerOperator             AccountRole = "CUSTOMER_OPERATOR"                    // 普通版工作台-协作者
	RoleAgent                        AccountRole = "AGENT"                                // 代理商
	RoleChildAgent                   AccountRole = "CHILD_AGENT"                          // 二级代理商
	RolePlatformStar                 AccountRole = "PLATFORM_ROLE_STAR"                   // 星图账户
	RolePlatformShopAccount          AccountRole = "PLATFORM_ROLE_SHOP_ACCOUNT"           // 抖音店铺账户
	RolePlatformQianchuanAgent       AccountRole = "PLATFORM_ROLE_QIANCHUAN_AGENT"        // 千川代理商
	RolePlatformStarAgent            AccountRole = "PLATFORM_ROLE_STAR_AGENT"             // 星图代理商
	RolePlatformAweme                AccountRole = "PLATFORM_ROLE_AWEME"                  // 抖音号
	RolePlatformStarMCN              AccountRole = "PLATFORM_ROLE_STAR_MCN"               // 星图MCN机构
	RolePlatformStarISV              AccountRole = "PLATFORM_ROLE_STAR_ISV"               // 星图服务商
	RoleAgentSystemAccount           AccountRole = "AGENT_SYSTEM_ACCOUNT"                 // 代理商系统账户
	RolePlatformLocalAgent           AccountRole = "PLATFORM_ROLE_LOCAL_AGENT"            // 本地推代理商
	RolePlatformYuntuBrandISVAdmin   AccountRole = "PLATFORM_ROLE_YUNTU_BRAND_ISV_ADMIN"  // 云图品牌服务商管理员
	RolePlatformLife                 AccountRole = "PLATFORM_ROLE_LIFE"                   // 抖音家客账户
	RolePlatformEnterpriseBPAdmin    AccountRole = "PLATFORM_ROLE_ENTERPRISE_BP_ADMIN"    // 升级版工作台管理员
	RolePlatformEnterpriseBPOperator AccountRole = "PLATFORM_ROLE_ENTERPRISE_BP_OPERATOR" // 升级版工作台协作者
)

// AccountSource 账户类型常量
const (
	AccountSourceAD    = "AD"    // 巨量营销客户账号
	AccountSourceLOCAL = "LOCAL" // 本地推
)

// AccountType 账户类型常量
const (
	AccountTypeNormal  = "AD_NORMAL" // 普通客户账号
	AccountTypeDouPlus = "DOU_PLUS"  // DOU+类客户账号
	AccountTypeLocal   = "LOCAL"     // 本地推客户账号
)

type accessTokenReq struct {
	AccessToken string `json:"access_token,omitempty"`
}

func (a *accessTokenReq) Format() {
	a.AccessToken = strings.TrimSpace(a.AccessToken)
}

func (a *accessTokenReq) Validate() (err error) {
	if len(a.AccessToken) <= 0 {
		err = errors.New("access token is empty")
		return
	}
	return
}

func (a *accessTokenReq) GetHeaders() map[string]string {
	headers := make(map[string]string)
	headers["Access-Token"] = a.AccessToken
	a.AccessToken = ""
	return headers
}

type BaseResp struct {
	Code      int         `json:"code"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
	RequestId string      `json:"request_id"`
}

// PageInfo 分页信息
type PageInfo struct {
	Page        int `json:"page"`         // 页码
	PageSize    int `json:"page_size"`    // 页面大小
	TotalPage   int `json:"total_page"`   // 总页数
	TotalNumber int `json:"total_number"` // 总数
}
