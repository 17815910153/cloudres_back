package param

//前端进行数据传递、 手机号+验证码

type SMSLoginParam struct {
	Phone string `json:"phone"`
	Code string `json:"code"`
}
