package vo

type LoginReq struct {
	Code string `form:"code"`
}

type LoginResp struct {
	Nickname string `json:"nickname"`
	Headimgurl string `json:"headimgurl"`
}
