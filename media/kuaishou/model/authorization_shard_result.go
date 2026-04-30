package model

import "errors"

// AuthorizationShardResultReq 共享授权结果查询请求
type AuthorizationShardResultReq struct {
	accessTokenReq
	ShardAuthId  int64 `json:"shard_auth_id"` // 共享授权任务ID，必填
	AdvertiserId int64 `json:"advertiser_id"` // 广告主id，必填
}

func (receiver *AuthorizationShardResultReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *AuthorizationShardResultReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	if receiver.ShardAuthId <= 0 {
		err = errors.New("shard_auth_id is empty")
		return
	}
	return
}

// AuthAccount 授权账户信息
type AuthAccount struct {
	AccountId   int64  `json:"account_id"`   // 账户ID
	AccountName string `json:"account_name"` // 账户名称
	FailReason  string `json:"fail_reason"`  // 授权失败原因
}

// AuthorizationShardResultResp 共享授权结果响应数据（仅data部分）
type AuthorizationShardResultResp struct {
	Status         int           `json:"status"`          // 授权任务状态：0-处理中，1-全部成功，2-部分成功，3-全部失败
	SuccessAccount []AuthAccount `json:"success_account"` // 共享授权成功的账户信息
	FailAccount    []AuthAccount `json:"fail_account"`    // 共享授权失败的账户信息
}
