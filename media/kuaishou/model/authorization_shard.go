package model

import "errors"

// AuthorizationShardReq 共享授权请求
type AuthorizationShardReq struct {
	accessTokenReq
	AdvertiserId        int64  `json:"advertiser_id"`         // 广告主id，必填
	AuthId              int64  `json:"auth_id"`               // 要共享的授权ID，必填
	ShardAuthorizeScope int    `json:"shard_authorize_scope"` // 共享授权范围：0-共享快手号授权，1-共享快手号授权+组件授权，必填
	ShardAccountId      string `json:"shard_account_id"`      // 要共享给账户，accountId逗号分隔，必填
	ShardUserType       int    `json:"shard_user_type"`       // 共享授权类型：1-同主体同代理商，2-同主体，7-同代理
}

func (receiver *AuthorizationShardReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *AuthorizationShardReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	if receiver.AuthId <= 0 {
		err = errors.New("auth_id is empty")
		return
	}
	if receiver.ShardAuthorizeScope < 0 || receiver.ShardAuthorizeScope > 1 {
		err = errors.New("shard_authorize_scope must be 0 or 1")
		return
	}
	if len(receiver.ShardAccountId) == 0 {
		err = errors.New("shard_account_id is empty")
		return
	}
	return
}

// AuthorizationShardResp 共享授权响应数据（仅data部分）
type AuthorizationShardResp struct {
	AuthId int64 `json:"auth_id"` // 共享授权记录ID
}
