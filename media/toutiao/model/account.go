package model

import (
	"errors"
	"github.com/cngamesdk/media-sdk/model"
	"github.com/duke-git/lancet/v2/datetime"
)

// AccountReq 账户请求
type AccountReq struct {
	*model.AccountReq
	AdvertiserIds []int64           `json:"advertiser_ids,omitempty"`
	Headers       map[string]string `json:"-"`
}

func (receiver *AccountReq) Convert(req *model.AccountReq) {
	receiver.AccountReq = req
	return
}

func (receiver *AccountReq) Format() {
	if receiver.AdvertiserID > 0 {
		receiver.AdvertiserIds = append(receiver.AdvertiserIds, receiver.AdvertiserID)
		receiver.AdvertiserID = 0
	}
	if receiver.Headers == nil {
		receiver.Headers = make(map[string]string)
	}
	receiver.Headers["Access-Token"] = receiver.AccessToken
	receiver.AccessToken = ""
	return
}

func (receiver *AccountReq) Validate() (err error) {
	if len(receiver.AdvertiserIds) <= 0 {
		err = errors.New("客户ID集合为空")
		return
	}
	accessToken, accessTokenOk := receiver.Headers["Access-Token"]
	if !accessTokenOk {
		err = errors.New("Access-Token不存在")
		return
	}
	if len(accessToken) <= 0 {
		err = errors.New("Access-Token为空")
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
