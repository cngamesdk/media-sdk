package model

import "errors"

// SharedWalletChargeReq 客户共享钱包账户充值/退款请求
type SharedWalletChargeReq struct {
	accessTokenReq
	WalletId      string `json:"wallet_id"`           // 钱包ID，必填
	WalletName    string `json:"wallet_name"`         // 钱包名称，必填
	AgentId       string `json:"agent_id"`            // 代理商ID，必填
	AgentName     string `json:"agent_name"`          // 代理商名称，必填
	TradeNo       string `json:"trade_no"`            // 交易流水号，必填。格式：mapi_{共享钱包ID}_{代理商ID}_{自定义编号}
	TradeType     int64  `json:"trade_type"`          // 交易类型，必填：16=充值(转入) 17=退款(转出)
	RelatedFlowNo string `json:"related_flow_no"`     // 关联业务方流水号，必填
	BizTradeTime  string `json:"biz_trade_time"`      // 交易时间戳，必填
	TotalAmount   int64  `json:"total_amount"`        // 交易金额，必填
	UserId        string `json:"user_id"`             // 操作人快手ID，必填
	Operator      string `json:"operator"`            // 操作人快手昵称，必填
	AppId         int64  `json:"app_id"`              // 业务线，必填，磁力智投默认传7
	SignCompany   string `json:"sign_company"`        // 签约主体信息，必填
	FundsOpType   int64  `json:"funds_op_type"`       // 资金操作类型，必填：1=现金操作 2=信用操作
	UserData      string `json:"user_data,omitempty"` // 用户自定义数据，可选
}

func (receiver *SharedWalletChargeReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *SharedWalletChargeReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.WalletId == "" {
		err = errors.New("wallet_id is empty")
		return
	}
	if receiver.AgentId == "" {
		err = errors.New("agent_id is empty")
		return
	}
	if receiver.TradeNo == "" {
		err = errors.New("trade_no is empty")
		return
	}
	if receiver.TotalAmount <= 0 {
		err = errors.New("total_amount is empty")
		return
	}
	return
}

// SharedWalletChargeResp 客户共享钱包账户充值/退款响应数据（仅data部分）
type SharedWalletChargeResp struct {
	TradeStatus int64  `json:"trade_status"` // 交易状态
	TradeNo     string `json:"trade_no"`     // 接入方流水号
}
