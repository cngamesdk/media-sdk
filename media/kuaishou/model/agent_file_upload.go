package model

import "errors"

// AgentFileUploadReq 上传资质文件请求
type AgentFileUploadReq struct {
	accessTokenReq
	AgentId  int64  `json:"agent_id"` // 代理商id，必填
	File     []byte `json:"-"`        // 文件二进制内容，必填，支持jpg、jpeg、pdf、gif格式，小于5MB
	FileName string `json:"-"`        // 文件名，必填
}

func (receiver *AgentFileUploadReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *AgentFileUploadReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AgentId <= 0 {
		err = errors.New("agent_id is empty")
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

// AgentFileUploadResp 上传资质文件响应数据（仅data部分）
type AgentFileUploadResp struct {
	FileToken string `json:"file_token"` // 资质文件标识
	Url       string `json:"url"`        // 文件预览链接
}
