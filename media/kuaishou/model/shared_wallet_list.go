package model

import "errors"

// SharedWalletListReq 客户共享钱包列表请求
type SharedWalletListReq struct {
	accessTokenReq
	AgentId  string `json:"agent_id"`            // 代理商ID，必填
	PageNum  int64  `json:"page_num,omitempty"`  // 页码
	PageSize int64  `json:"page_size,omitempty"` // 每页大小
}

func (receiver *SharedWalletListReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *SharedWalletListReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AgentId == "" {
		err = errors.New("agent_id is empty")
		return
	}
	return
}

// SharedWalletListInfo 共享钱包信息
type SharedWalletListInfo struct {
	WalletId                  string `json:"wallet_id"`                    // 钱包ID
	WalletName                string `json:"wallet_name"`                  // 钱包名称
	AppId                     int64  `json:"app_id"`                       // 业务线ID
	AgentId                   string `json:"agent_id"`                     // 代理商ID
	AgentName                 string `json:"agent_name"`                   // 代理商名称
	TotalBalance              string `json:"total_balance"`                // 总余额
	Cash                      string `json:"cash"`                         // 现金余额
	CreateTime                string `json:"create_time"`                  // 创建时间时间戳
	AccountSharedWalletStatus int64  `json:"account_shared_wallet_status"` // 钱包状态：1=已生效 2=失效中 3=已失效
}

// SharedWalletListResp 客户共享钱包列表响应数据（仅data部分）
type SharedWalletListResp []SharedWalletListInfo
