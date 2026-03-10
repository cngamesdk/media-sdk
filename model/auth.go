package model

type AuthRedirectReq struct {
	State       string
	AuthType    string
	CallbackUrl string
}
