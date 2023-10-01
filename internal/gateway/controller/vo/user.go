package vo

type Register_resquest struct {
	Email string `form:"email"`
	Pwd   string `form:"pwd"`
	Code  string `form:"authcode"`
}

type Register_response struct {
	Code int32    `json:"code"`
	Msg  string `json:"msg"`
}
type Login_resquest struct {
	Email string `form:"email"`
	Pwd   string `form:"pwd"`
}

type Login_by_auth_code_request struct {
	Email    string `form:"email"`
	AuthCode string `form:"authcode"`
}

type Login_response struct {
	Code  int32    `json:"code"`
	Msg   string `json:"msg"`
	Token string `json:"token"`
}

type Get_user_info_response struct {
	Code             int32  `json:"code"`
	ID               int64  `json:"uid" gorm:"primarykey"`
	NickName         string `json:"nickname"`
	Email            string `json:"email"`
	Avatar           string `json:"avatar"`
	IndividualResume string `json:"individual_resume"`
	Role             int32  `json:"role"`
}

type Update_user_request struct {
	Data   string `form:"data"`
	Action int32    `form:"action"`
}

type Update_user_response struct {
	Code int32    `json:"code"`
	Data string `json:"data"`
}

type Delete_user_response struct {
	Code int32  `json:"code"`
	Data string `json:"data"`
}
