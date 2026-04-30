package model

import "errors"

// AuaxProjectCreateReq 智投项目创建请求
type AuaxProjectCreateReq struct {
	accessTokenReq
	AdvertiserId     int64   `json:"advertiser_id"`      // 广告主账号ID，必填
	Name             string  `json:"name"`               // 项目名称，必填
	KolUserId        int64   `json:"kol_user_id"`        // 快手号ID，必填
	KolUserType      int     `json:"kol_user_type"`      // 快手号类型：1-普通快手号，2-蓝V服务号，3-聚星达人，必填
	AutoManageType   int     `json:"auto_manage_type"`   // 智投类型：2-短剧智投，必填
	SubjectId        int64   `json:"subject_id"`         // 短剧ID/小说ID，必填
	OcpxActionType   int     `json:"ocpx_action_type"`   // 转化目标：191-首日ROI，990-首日变现ROI（短剧），必填
	RoiRatio         float64 `json:"roi_ratio"`          // ROI系数，必填
	Description      string  `json:"description"`        // 广告语，必填
	PhotoPackageInfo []int64 `json:"photo_package_info"` // 素材包ID，必填
}

func (receiver *AuaxProjectCreateReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *AuaxProjectCreateReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	if len(receiver.Name) == 0 {
		err = errors.New("name is empty")
		return
	}
	if receiver.KolUserId <= 0 {
		err = errors.New("kol_user_id is empty")
		return
	}
	if receiver.KolUserType < 1 || receiver.KolUserType > 3 {
		err = errors.New("kol_user_type must be 1, 2 or 3")
		return
	}
	if receiver.AutoManageType != 2 {
		err = errors.New("auto_manage_type must be 2 (短剧智投)")
		return
	}
	if receiver.SubjectId <= 0 {
		err = errors.New("subject_id is empty")
		return
	}
	if receiver.OcpxActionType != 191 && receiver.OcpxActionType != 990 {
		err = errors.New("ocpx_action_type must be 191 or 990")
		return
	}
	if receiver.RoiRatio <= 0 {
		err = errors.New("roi_ratio is empty")
		return
	}
	if len(receiver.Description) == 0 {
		err = errors.New("description is empty")
		return
	}
	if len(receiver.PhotoPackageInfo) == 0 {
		err = errors.New("photo_package_info is empty")
		return
	}
	return
}

// AuaxProjectCreateResp 智投项目创建响应数据（仅data部分）
type AuaxProjectCreateResp struct {
	AuaxProjectId int64 `json:"auax_project_id"` // 智投项目ID
}
