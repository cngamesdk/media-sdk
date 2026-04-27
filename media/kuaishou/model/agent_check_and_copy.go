package model

import "errors"

// AgentCheckAndCopyReq 代理商-复制账户请求
type AgentCheckAndCopyReq struct {
	accessTokenReq
	AgentId     int64   `json:"agent_id"`          // 代理商id，必填
	AccountList []int64 `json:"account_list"`      // 被复制账户的广告主Ids，必填
	UserId      int64   `json:"user_id,omitempty"` // 复制后账户绑定的快手Id，不更换userId的情况下不传
	Code        string  `json:"code,omitempty"`    // 验证码，不更换userId的情况下不传
	Email       string  `json:"email,omitempty"`   // 邮箱，不更换userId的情况下不传
	UcType      string  `json:"uc_type"`           // 广告主开户类型：DSP_MAPI，必填
	CopyNumber  int     `json:"copy_number"`       // 需要复制出的广告主数量，必填
}

func (receiver *AgentCheckAndCopyReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *AgentCheckAndCopyReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AgentId <= 0 {
		err = errors.New("agent_id is empty")
		return
	}
	if len(receiver.AccountList) == 0 {
		err = errors.New("account_list is empty")
		return
	}
	if len(receiver.UcType) == 0 {
		err = errors.New("uc_type is empty")
		return
	}
	if receiver.CopyNumber <= 0 {
		err = errors.New("copy_number is empty")
		return
	}
	return
}

// AgentCopyAccountResult 复制账户结果
type AgentCopyAccountResult struct {
	OldAccountId   int64  `json:"old_account_id"`   // 被复制账户Id
	OldAccountName string `json:"old_account_name"` // 被复制账户名称
	NewAccountId   int64  `json:"new_account_id"`   // 复制出的账户Id
	Success        bool   `json:"success"`          // 是否复制成功
	ErrMsg         string `json:"err_msg"`          // 复制失败原因
	NewAccountName string `json:"new_account_name"` // 复制出的账户名称
	NewUserId      int64  `json:"new_user_id"`      // 复制出的账户快手Id
	NewUserName    string `json:"new_user_name"`    // 复制出的账户快手昵称
}

// AgentCheckAndCopyResp 代理商-复制账户响应数据（仅data部分）
type AgentCheckAndCopyResp []AgentCopyAccountResult
