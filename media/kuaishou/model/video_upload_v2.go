package model

import "errors"

// VideoUploadV2Req 上传视频v2请求
type VideoUploadV2Req struct {
	accessTokenReq
	AdvertiserId         int64  `json:"-"` // 广告主ID，必填
	Signature            string `json:"-"` // 视频MD5值，必填
	File                 []byte `json:"-"` // 视频文件二进制内容，与BlobStoreKey二选一
	FileName             string `json:"-"` // 视频文件名，上传文件时填写
	BlobStoreKey         string `json:"-"` // 视频存储地址token，流式/分片上传时使用，与File二选一
	PhotoName            string `json:"-"` // 视频名称，最多150个字符
	PhotoTag             string `json:"-"` // 视频标签，单个标签最多10个字符
	Sync                 int    `json:"-"` // 上传模式，0:异步(默认) 1:同步
	Type                 int    `json:"-"` // 视频上传类型，1:信息流竖版(默认) 2:信息流横版 5-9:开屏视频
	ShieldBackwardSwitch *bool  `json:"-"` // 是否自动同步到个人主页，false:屏蔽 true:不屏蔽
	NativePlcSwitch      *bool  `json:"-"` // 是否挂载原生PLC组件，false:不挂载 true:挂载
}

func (receiver *VideoUploadV2Req) Format() {
	receiver.accessTokenReq.Format()
	if receiver.Type <= 0 {
		receiver.Type = 1
	}
}

func (receiver *VideoUploadV2Req) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	if receiver.Signature == "" {
		err = errors.New("signature is empty")
		return
	}
	if len(receiver.File) == 0 && receiver.BlobStoreKey == "" {
		err = errors.New("file or blob_store_key is required")
		return
	}
	return
}

// VideoUploadV2Resp 上传视频v2响应数据（仅data部分）
type VideoUploadV2Resp struct {
	PhotoId   string `json:"photo_id"`  // 视频ID（加密）
	PhotoIdV1 string `json:"photoId"`   // 视频ID（旧字段，已废弃）
	Signature string `json:"signature"` // MD5签名
}
