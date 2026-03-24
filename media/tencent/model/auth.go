package model

import (
	"errors"
	"github.com/cngamesdk/media-sdk/model"
	"net/url"
	"strings"
	"time"
)

type OAuth2AuthorizeReq struct {
	ClientID    int64  `json:"client_id"`              // 应用id (必填)
	RedirectURI string `json:"redirect_uri"`           // 应用回调地址 (必填)
	State       string `json:"state,omitempty"`        // 验证请求有效性参数
	Scope       string `json:"scope,omitempty"`        // 授权范围
	AccountType string `json:"account_type,omitempty"` // 授权账号类型，默认ACCOUNT_TYPE_QQ
}

// 常量定义 - 账号类型
const (
	AccountTypeWechat = "ACCOUNT_TYPE_WECHAT" // 微信账号
	AccountTypeQQ     = "ACCOUNT_TYPE_QQ"     // QQ账号（默认）
)

// 常量定义 - 授权范围
const (
	ScopeAdsManagement      = "ads_management"      // 广告投放
	ScopeAdsInsights        = "ads_insights"        // 数据洞察
	ScopeAccountManagement  = "account_management"  // 账号服务
	ScopeAudienceManagement = "audience_management" // 人群管理
	ScopeUserActions        = "user_actions"        // 用户行为数据接入
)

// 默认值常量
const (
	DefaultAccountType = AccountTypeQQ
)

// 长度限制常量
const (
	MinRedirectURILength = 1
	MaxRedirectURILength = 1024
	MaxStateLength       = 512
	MinScopeLength       = 1
	MaxScopeLength       = 64
)

func (p *OAuth2AuthorizeReq) Convert(req interface{}) (err error) {
	result, ok := req.(*OAuth2AuthorizeReq)
	if !ok {
		err = errors.New("请求参数类型不一致")
		return
	}
	*p = *result
	return
}

func (p *OAuth2AuthorizeReq) Format() {

}

// Validate 验证OAuth2授权参数
func (p *OAuth2AuthorizeReq) Validate() error {
	// 1. 验证client_id
	if p.ClientID == 0 {
		return errors.New("client_id为必填")
	}

	// 2. 验证redirect_uri
	if p.RedirectURI == "" {
		return errors.New("redirect_uri为必填")
	}
	if len(p.RedirectURI) < MinRedirectURILength || len(p.RedirectURI) > MaxRedirectURILength {
		return errors.New("redirect_uri长度必须在1-1024字节之间")
	}
	if err := validateRedirectURI(p.RedirectURI); err != nil {
		return err
	}

	// 3. 验证state
	if p.State != "" && len(p.State) > MaxStateLength {
		return errors.New("state长度不能超过512字节")
	}

	// 4. 验证scope
	if p.Scope != "" {
		if len(p.Scope) < MinScopeLength || len(p.Scope) > MaxScopeLength {
			return errors.New("scope长度必须在1-64字节之间")
		}
		if err := validateScope(p.Scope); err != nil {
			return err
		}
	}

	// 5. 设置默认账号类型
	if p.AccountType == "" {
		p.AccountType = DefaultAccountType
	}
	if p.AccountType != AccountTypeWechat && p.AccountType != AccountTypeQQ {
		return errors.New("account_type值无效，允许值：ACCOUNT_TYPE_WECHAT、ACCOUNT_TYPE_QQ")
	}

	return nil
}

// validateRedirectURI 验证回调地址
func validateRedirectURI(uri string) error {
	// 必须是http或https协议
	if !strings.HasPrefix(uri, "http://") && !strings.HasPrefix(uri, "https://") {
		return errors.New("redirect_uri仅支持http和https协议")
	}

	// 解析URL
	parsedURL, err := url.Parse(uri)
	if err != nil {
		return errors.New("redirect_uri格式无效")
	}

	// 不支持指定端口号（即不能包含端口号或端口号为默认值）
	if parsedURL.Port() != "" {
		return errors.New("redirect_uri不支持指定端口号")
	}

	return nil
}

// validateScope 验证授权范围
func validateScope(scope string) error {
	// 允许的scope值
	validScopes := map[string]bool{
		ScopeAdsManagement:      true,
		ScopeAdsInsights:        true,
		ScopeAccountManagement:  true,
		ScopeAudienceManagement: true,
		ScopeUserActions:        true,
	}

	// 可能多个scope用空格分隔
	scopes := strings.Fields(scope)
	for _, s := range scopes {
		if !validScopes[s] {
			return errors.New("scope包含无效值，允许值：ads_management、ads_insights、account_management、audience_management、user_actions")
		}
	}

	return nil
}

// OAuth2AuthorizeResp OAuth2授权响应
type OAuth2AuthorizeResp struct {
	RedirectURL string `json:"redirect_url"` // 授权跳转地址
}

type AccessTokenReq struct {
	model.AccessTokenReq
}

func (receiver *AccessTokenReq) Convert(req *model.AccessTokenReq) {
	receiver.AccessTokenReq = *req
}

func (receiver *AccessTokenReq) Format() {
	receiver.AccessTokenReq.Format()
}

func (receiver *AccessTokenReq) Validate() (err error) {
	if validateErr := receiver.AccessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AppId <= 0 {
		err = errors.New("app_id is empty")
		return
	}
	if len(receiver.Secret) <= 0 {
		err = errors.New("secret is empty")
		return
	}
	if len(receiver.AuthCode) <= 0 {
		err = errors.New("AuthCode为空")
		return
	}
	return
}

type AccessTokenResp struct {
	model.AccessTokenResp
}

func (receiver *AccessTokenResp) Convert() (*model.AccessTokenResp, error) {
	return &receiver.AccessTokenResp, nil
}

type RefreshTokenReq struct {
	model.RefreshTokenReq
}

func (receiver *RefreshTokenReq) Format() {
	receiver.Secret = strings.TrimSpace(receiver.Secret)
	receiver.RefreshToken = strings.TrimSpace(receiver.RefreshToken)
}

func (receiver *RefreshTokenReq) Validate() (err error) {
	if receiver.AppId <= 0 {
		err = errors.New("app_id is empty")
		return
	}
	if len(receiver.Secret) <= 0 {
		err = errors.New("secret is empty")
		return
	}
	if len(receiver.RefreshToken) <= 0 {
		err = errors.New("refresh token is empty")
		return
	}
	return
}

func (receiver *RefreshTokenReq) Convert(req *model.RefreshTokenReq) {
	receiver.RefreshTokenReq = *req
	return
}

type RefreshTokenResp struct {
	model.RefreshTokenResp
}

func (receiver *RefreshTokenResp) Convert() (resp *model.RefreshTokenResp, err error) {
	receiver.ExpireTime = time.Now().Add(time.Duration(receiver.ExpiresIn) * time.Second)
	receiver.RefreshTokenExpireTime = time.Now().Add(time.Duration(receiver.RefreshTokenExpireIn) * time.Second)
	resp = &receiver.RefreshTokenResp
	return
}
