package model

import (
	"errors"
	"github.com/spf13/cast"
)

const (
	ActionTypeCompleteOrder         = "COMPLETE_ORDER"          // 下单
	ActionTypeReservation           = "RESERVATION"             // 表单预约
	ActionTypeRegister              = "REGISTER"                // 注册
	ActionTypeActivateApp           = "ACTIVATE_APP"            // 激活
	ActionTypeStartApp              = "START_APP"               // 次日留存 / 7日留存
	ActionTypePurchase              = "PURCHASE"                // 付费 / 首次付费
	ActionTypeConfirmEffectiveLeads = "CONFIRM_EFFECTIVE_LEADS" // 有效综合线索
	ActionTypeCredit                = "CREDIT"                  // 授信
	ActionTypeFollow                = "FOLLOW"                  // 关注
	ActionTypeScanCode              = "SCANCODE"                // 加企业微信客服
	ActionTypeViewContent           = "VIEW_CONTENT"            // 关键页面访问
	ActionTypeAdPurchase            = "AD_PURCHASE"             // 广告变现
	ActionTypeProductView           = "PRODUCT_VIEW"            // 商品详情页浏览
	ActionTypeLandingPageClick      = "LANDING_PAGE_CLICK"      // 跳转按钮点击
	ActionTypeScanCodeWx            = "SCANCODE_WX"             // 扫码加粉
	ActionTypeOnlineConsult         = "ONLINE_CONSULT"          // 在线咨询
	ActionTypeAddGroup              = "ADD_GROUP"               // 加群
	ActionTypeAddDesktop            = "ADD_DESKTOP"             // 快应用加桌面
	ActionTypeMakePhoneCall         = "MAKE_PHONE_CALL"         // 电话拨打
	ActionTypeCreateRole            = "CREATE_ROLE"             // 小游戏创角
	ActionTypeClaimOffer            = "CLAIM_OFFER"             // 领券/综合线索收集
	ActionTypeReservationCheck      = "RESERVATION_CHECK"       // 意向表单（第三张图片）
	ActionTypeApply                 = "APPLY"                   // 进件
	ActionTypePrePay                = "PRE_PAY"                 // 预付定金
	ActionTypeConsult               = "CONSULT"                 // 主动一句话咨询 / 三句话咨询

	// 特殊关键行为
	ActionTypeCustom       = "CUSTOM"         // 自定义 action_type
	ActionTypeUvCoreAction = "UV_CORE_ACTION" // 关键行为（custom_action 上报）
)

// 转化上报
type ConversionReportReq struct {
	Callback string                   `json:"callback"`
	Data     *ConversionReportActions `json:"data"`
}

func (a *ConversionReportReq) Validate() (err error) {
	if len(a.Callback) <= 0 {
		err = errors.New("callback is empty")
		return
	}
	if a.Data == nil {
		err = errors.New("data is empty")
		return
	}
	if len(a.Data.Actions) <= 0 {
		err = errors.New("data.actions is empty")
		return
	}
	for _, item := range a.Data.Actions {
		if item == nil {
			err = errors.New("data.actions.item exists nil pointer")
			return
		}
		if validateErr := item.Validate(); validateErr != nil {
			err = validateErr
			return
		}
	}
	return
}

type ConversionReportInterface interface {
	Validate() error
}

type ConversionReportActions struct {
	Actions []ConversionReportInterface `json:"actions,omitempty"`
}

type ConversionReportActionCommon struct {
	OuterActionId string                       `json:"outer_action_id,omitempty"`
	ActionTime    int64                        `json:"action_time,omitempty"`
	UserId        ConversionReportInterface    `json:"user_id,omitempty"`
	ActionType    string                       `json:"action_type,omitempty"`
	ActionParam   *ConversionReportActionParam `json:"action_param,omitempty"`
}

func (receiver *ConversionReportActionCommon) Validate() (err error) {
	if len(receiver.ActionType) <= 0 {
		err = errors.New("action_type is empty")
		return
	}
	if receiver.ActionType == ActionTypeCompleteOrder || receiver.ActionType == ActionTypePurchase {
		if receiver.ActionParam == nil {
			err = errors.New("action_param is empty")
			return
		}
		if cast.ToInt(receiver.ActionParam.Value) <= 0 {
			err = errors.New("action_param.value is empty")
			return
		}
	}
	if receiver.UserId == nil {
		err = errors.New("user_id is empty")
		return
	}
	if validateErr := receiver.UserId.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	return
}

type ConversionReportActionApp struct {
	ConversionReportActionCommon
}

type ConversionReportActionUserIdApp struct {
	HashImei      string `json:"hash_imei,omitempty"`
	HashIdfa      string `json:"hash_idfa,omitempty"`
	HashAndroidId string `json:"hash_android_id,omitempty"`
	Oaid          string `json:"oaid,omitempty"`
	HashOaid      string `json:"hash_oaid,omitempty"`
	Caid          string `json:"caid,omitempty"`
	CaidVersion   string `json:"caid_version,omitempty"`
}

func (receiver ConversionReportActionUserIdApp) Validate() (err error) {
	return
}

type ConversionReportActionWeb struct {
	ConversionReportActionCommon
	Url   string `json:"url"`
	Trace struct {
		ClickId string `json:"click_id"`
	} `json:"trace"`
}

func (receiver *ConversionReportActionWeb) Validate() (err error) {
	if validateErr := receiver.ConversionReportActionCommon.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if len(receiver.Url) <= 0 {
		err = errors.New("url is empty")
		return
	}
	if len(receiver.Trace.ClickId) <= 0 {
		err = errors.New("click_id is empty")
		return
	}
	return
}

type ConversionReportActionUserIdWeb struct {
	HashImei      string `json:"hash_imei,omitempty"`
	HashIdfa      string `json:"hash_idfa,omitempty"`
	HashAndroidId string `json:"hash_android_id,omitempty"`
	Oaid          string `json:"oaid,omitempty"`
	HashOaid      string `json:"hash_oaid,omitempty"`
	Ip            string `json:"ip,omitempty"`
	UserAgent     string `json:"user_agent,omitempty"`
}

type ConversionReportActionMiniGame struct {
	ConversionReportActionCommon
}

type ConversionReportActionUserIdMiniGame struct {
	HashImei      string `json:"hash_imei,omitempty"`
	HashIdfa      string `json:"hash_idfa,omitempty"`
	HashAndroidId string `json:"hash_android_id,omitempty"`
	Oaid          string `json:"oaid,omitempty"`
	HashOaid      string `json:"hash_oaid,omitempty"`
	WechatAppId   string `json:"wechat_app_id,omitempty"`
	WechatOpenid  string `json:"wechat_openid,omitempty"`
}

func (receiver *ConversionReportActionUserIdMiniGame) Validate() (err error) {
	if len(receiver.WechatAppId) <= 0 {
		err = errors.New("wechat_app_id is empty")
		return
	}
	if len(receiver.WechatOpenid) <= 0 {
		err = errors.New("wechat_openid is empty")
		return
	}
	return
}

type ConversionReportActionMiniProgram struct {
	ConversionReportActionCommon
}

type ConversionReportActionUserIdMiniProgram struct {
	ConversionReportActionUserIdMiniGame
	WechatUnionid string `json:"wechat_unionid,omitempty"`
}

func (receiver *ConversionReportActionUserIdMiniProgram) Validate() (err error) {
	if len(receiver.WechatAppId) <= 0 {
		err = errors.New("wechat_app_id is empty")
		return
	}
	if len(receiver.WechatOpenid) <= 0 && len(receiver.WechatUnionid) <= 0 {
		err = errors.New("wechat_openid is empty and wechat_unionid is empty")
		return
	}
	return
}

// ConversionReportActionOfficialAccount 微信公众号
type ConversionReportActionOfficialAccount struct {
	ConversionReportActionCommon
	Trace struct {
		ClickId string `json:"click_id"`
	} `json:"trace"`
}

func (receiver *ConversionReportActionOfficialAccount) Validate() (err error) {
	if validateErr := receiver.ConversionReportActionCommon.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if len(receiver.Trace.ClickId) <= 0 {
		err = errors.New("trace.click_id is empty")
		return
	}
	return
}

type ConversionReportActionUserIdOfficialAccount struct {
	WechatOpenid  string `json:"wechat_openid,omitempty"`
	WechatUnionid string `json:"wechat_unionid,omitempty"`
	WechatAppId   string `json:"wechat_app_id,omitempty"`
}

func (receiver *ConversionReportActionUserIdOfficialAccount) Validate() (err error) {
	if len(receiver.WechatAppId) <= 0 {
		err = errors.New("wechat_app_id is empty")
		return
	}
	if len(receiver.WechatOpenid) <= 0 && len(receiver.WechatUnionid) <= 0 {
		err = errors.New("wechat_openid is empty and wechat_unionid is empty")
		return
	}
	return
}

// ConversionReportActionWeCom 企业微信
type ConversionReportActionWeCom struct {
	ConversionReportActionCommon
	Trace struct {
		ClickId string `json:"click_id"`
	} `json:"trace"`
}

func (receiver *ConversionReportActionWeCom) Validate() (err error) {
	if validateErr := receiver.ConversionReportActionCommon.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if len(receiver.Trace.ClickId) <= 0 {
		err = errors.New("trace.click_id is empty")
		return
	}
	return
}

type ConversionReportActionUserIdWeCom struct {
	WechatOpenid  string `json:"wechat_openid,omitempty"`
	WechatUnionid string `json:"wechat_unionid,omitempty"`
	WechatAppId   string `json:"wechat_app_id,omitempty"`
}

func (receiver *ConversionReportActionUserIdWeCom) Validate() (err error) {
	if len(receiver.WechatAppId) <= 0 {
		err = errors.New("wechat_app_id is empty")
		return
	}
	if len(receiver.WechatOpenid) <= 0 && len(receiver.WechatUnionid) <= 0 {
		err = errors.New("wechat_openid is empty and wechat_unionid is empty")
		return
	}
	return
}

type ConversionReportActionParam struct {
	Value interface{} `json:"value,omitempty"`
}

type ConversionReportResp struct {
}
