package model

import "errors"

// AuaxProjectUpdateReq 智投项目更新请求
type AuaxProjectUpdateReq struct {
	accessTokenReq
	AdvertiserId     int64   `json:"advertiser_id"`      // 广告主账号ID，必填
	AuaxProjectId    int64   `json:"auax_project_id"`    // 智投项目ID，必填
	OcpxActionType   int     `json:"ocpx_action_type"`   // 转化目标：191-首日ROI，990-首日变现ROI（短剧），必填
	RoiRatio         float64 `json:"roi_ratio"`          // ROI系数，必填
	PhotoPackageInfo []int64 `json:"photo_package_info"` // 素材包ID，必填
}

func (receiver *AuaxProjectUpdateReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *AuaxProjectUpdateReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	if receiver.AuaxProjectId <= 0 {
		err = errors.New("auax_project_id is empty")
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
	if len(receiver.PhotoPackageInfo) == 0 {
		err = errors.New("photo_package_info is empty")
		return
	}
	return
}
