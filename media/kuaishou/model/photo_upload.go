package model

import "errors"

// PhotoUploadReq 本地视频上传请求
type PhotoUploadReq struct {
	accessTokenReq
	AdvertiserId         int64    `json:"advertiser_id"`        // 广告主id，必填
	AuthorId             int64    `json:"authorId"`             // 原生上传至达人的快手号，必填
	ShieldBackwardSwitch bool     `json:"shieldBackwardSwitch"` // 上传后是否自动同步至个人主页，false=屏蔽
	PhotoCaption         string   `json:"photoCaption"`         // 视频描述，最长30字符
	NativePlcSwitch      bool     `json:"nativePlcSwitch"`      // 挂载plc组件
	Sync                 int      `json:"sync"`                 // 1=同步，0=异步（默认同步）
	PhotoTag             []string `json:"photo_tag"`            // 标签
	PhotoName            string   `json:"photo_name"`           // 视频名称
	BlobStoreKey         string   `json:"blob_store_key"`       // blob存储key
	Signature            string   `json:"signature"`            // md5
	File                 []byte   `json:"-"`                    // 视频文件二进制内容，必填
	FileName             string   `json:"-"`                    // 视频文件名，必填
}

func (receiver *PhotoUploadReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *PhotoUploadReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	if receiver.AuthorId <= 0 {
		err = errors.New("authorId is empty")
		return
	}
	if len(receiver.File) == 0 {
		err = errors.New("file is empty")
		return
	}
	if len(receiver.FileName) == 0 {
		err = errors.New("file_name is empty")
		return
	}
	return
}

// PhotoUploadResp 本地视频上传响应数据（仅data部分）
type PhotoUploadResp struct {
	PhotoId int64 `json:"photo_id"` // 视频ID（加密）
}
