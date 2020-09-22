package service

import (
	"CloudRes/dao"
	"CloudRes/model"
	"CloudRes/param"
	"CloudRes/tool"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
)

type MemberService struct {

}

// 使用账号密码进行登录
func (ms *MemberService) PwdLogin(loginparam param.PwdLoginParam) *model.Member  {
	md := dao.MemberDao{tool.DbEngine}
	//进行密码的加密
	md5Pwd := tool.Md5(loginparam.Pwd)
	member := md.QueryByPhoneAndPwd(loginparam.Phone, md5Pwd)
	if member.Id == 0{
		return nil
	}
	return member

}
//进行登录 手机号+验证码
func (ms *MemberService) SmsLogin(loginparam param.SMSLoginParam) *model.Member  {
	// 1.获取到手机号和验证码

	// 2.验证手机号member表中查询记录

	md := dao.MemberDao{tool.DbEngine}
	sms := md.ValidateSmsCode(loginparam.Phone, loginparam.Code)
	if sms.Id == 0 {
		return nil
	}

	// 3, 根据手机号member中查询记录
	member := md.QueryByPhone(loginparam.Phone)
	if member.Id != 0 {
		return member
	}
	// 4. 新建一个member记录，并保存
	user := model.Member{}
	user.UserName = loginparam.Phone
	user.Mobile = loginparam.Phone
	user.RegisterTime = time.Now().Unix()

	user.Id = md.InsertMember(user)

	return &user

}

//发送手机验证码
func (ms *MemberService) SendCode (phone string) bool {

	// 1.产生一个验证码
	code := fmt.Sprintf("%06v",rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000))
	// 2.调用阿里云sdk，完成发送

	config := tool.GetConfig().Sms
	client, err := dysmsapi.NewClientWithAccessKey(config.RegionId, config.AppKey, config.AppSecret)

	if err != nil {
		panic(err)
		return false
	}
	//阿里云官方文档，设置请求参数
	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"
	request.SignName = config.SignName
	request.TemplateCode = config.TemplateCode
	request.PhoneNumbers = phone
	par, err := json.Marshal(map[string]interface{}{
		"code": code,
	})

	request.TemplateParam = string(par)

	response, err := client.SendSms(request)
	fmt.Println(response)
	if err != nil {
		fmt.Print(err.Error())
	}

	// 3.接收返回的信息，并判断发送状态
	// 短信验证码发送成功
	fmt.Println("返回的",response.Code)
	if response.Code == "OK" {
		//将验证码保存到数据库里面
		smsCode := model.SmsCode{Phone: phone, Code: code, BizId: response.BizId, CreateTime: time.Now().Unix()}

		memberDao := dao.MemberDao{Orm: tool.DbEngine}
		result := memberDao.InsertCode(smsCode)

		return result >0
	}


	return false
}