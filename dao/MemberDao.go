package dao

import (
	"CloudRes/model"
	"CloudRes/tool"
	"fmt"
)

type MemberDao struct {
	*tool.Orm
}
//根据手机号和密码进行登录
func (md *MemberDao) QueryByPhoneAndPwd(phone string, pwd string) *model.Member  {
	var member model.Member
	if _, err := md.Where("mobile = ? and password = ?", phone, pwd).Get(&member); err != nil{
		fmt.Println(err.Error())
	}
	return &member

}
// 验证用户的验证码是否与手机匹配
func (md *MemberDao)ValidateSmsCode(phone string, code string) *model.SmsCode {
	var sms model.SmsCode
	if _, err := md.Where("phone = ? and code = ?", phone, code).Get(&sms); err != nil{
		fmt.Println(err.Error())

	}

	return &sms

}
//根据手机号码查询用户
func (md *MemberDao) QueryByPhone(phone string ) *model.Member {
	var member model.Member
	if _, err := md.Where("mobile = ?", phone).Get(&member); err != nil {
		fmt.Println(err.Error())
	}
	return &member

}
//新增用户
func (md *MemberDao) InsertMember(member model.Member) int64  {
	result, err := md.InsertOne(&member)
	if err != nil {
		fmt.Println(err.Error())
		return 0

	}
	return result
}
func (md *MemberDao) InsertCode(sms model.SmsCode) int64  {
	result, err := md.InsertOne(&sms)
	if err != nil {
		panic(err.Error())
	}
	return result

}
