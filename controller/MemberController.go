package controller

import (
	"CloudRes/param"
	"CloudRes/service"
	"CloudRes/tool"
	"fmt"
	"github.com/gin-gonic/gin"
)

type MemberController struct {

}

var  (
	ms = service.MemberService{}
)

//类似构造函数 ,主要进行路由解析
func (mc *MemberController) Router(e *gin.Engine)  {
	e.GET("/api/sendcode", mc.sendSmsCode)
	e.POST("/api/login_sms", mc.smsLogin)
	e.POST("/api/login_pwd", mc.pwdLogin)

}


// http:localhost:8090/api/sendcode?phone = 15362124173
func (mc *MemberController) sendSmsCode(c *gin.Context) {
	//假数据进行测试
	//c.JSON(http.StatusOK, "短信发送成功")
	//发送验证码
	phone, ok := c.GetQuery("phone")
	//fmt.Println("收到的手机号为", phone)
	if !ok {
		tool.Failed(c,"参数解析失败")
		return
	}

	issend := ms.SendCode(phone)
	if issend {
		tool.Success(c,"发送成功")

		return
	}
	tool.Failed(c,"发送失败")
}

//手机号 + 短信 登录的方法
func (mc *MemberController) smsLogin(c *gin.Context)  {
	//提取参数
	var  smsLoginParam param.SMSLoginParam
	err := tool.Decode(c.Request.Body, &smsLoginParam)
	if err != nil {
		tool.Failed(c,"参数解析失败")
	}

	// 完成手机+验证码登录
	member := ms.SmsLogin(smsLoginParam)
	if member != nil {
		tool.Success(c,member)
		fmt.Println("登录成功",member)
		return
	}

	tool.Failed(c,"登录失败")
	fmt.Println("登录失败")
}

// 用户通过密码进行登录
func (mc *MemberController) pwdLogin(c *gin.Context)  {

	//提取参数
	var pwdLoginParam param.PwdLoginParam
	err := tool.Decode(c.Request.Body, &pwdLoginParam)
	if err !=nil {
		tool.Failed(c,"参数解析失败")
	}

	// 进行密码 + 手机号的登录
	member := ms.PwdLogin(pwdLoginParam)
	if member != nil {
		tool.Success(c, member)
		fmt.Println("登录成功",member)
		return
	}
	tool.Failed(c,"登录失败")

}