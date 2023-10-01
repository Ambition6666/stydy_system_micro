package user

import (
	"context"
	"errors"
	"fmt"
	"net/smtp"
	pb "studysystem_micro/idl/user"
	gen "studysystem_micro/internal/service/user/config"
	"studysystem_micro/internal/service/user/dao"
	"studysystem_micro/internal/service/user/model"
	"studysystem_micro/pkg/init/statuscode"
	"studysystem_micro/pkg/utils"

	e "github.com/jordan-wright/email"
)

// 用户注册
func (*UserServer) Register(ctx context.Context, in *pb.UserRegister) (*pb.CommonResponse, error) {
	u := dao.Search_user(in.Email)
	if u.ID != 0 {
		return &pb.CommonResponse{Code: statuscode.ERROR, Msg: "用户已存在"}, nil
	}
	pwd := utils.Encrypt(in.Pwd)
	id := utils.GetWorker().GetID()
	n := model.NewUser(id, in.Email, pwd)
	dao.Create_user(n)
	return &pb.CommonResponse{Code: 200, Msg: "注册成功"}, nil
}

// 获取验证码
func (*UserServer) GetAuthCode(ctx context.Context, in *pb.UserGetAuthCode) (*pb.CommonResponse, error) {
	err := SendEmail(in.Email)
	if err != nil {
		return &pb.CommonResponse{Code: statuscode.ERROR, Msg: "发送失败"}, err
	}
	return &pb.CommonResponse{Code: statuscode.OK, Msg: "发送成功"}, nil
}

// 用户通过验证码登录
func (*UserServer) Login(ctx context.Context, in *pb.UserLoginImformationsByAuthCode) (*pb.UserToken, error) {
	u := dao.Search_user(in.Email)
	if u.ID == 0 {
		return &pb.UserToken{
			Code:  statuscode.Loginfailed,
			Token: "",
		}, errors.New("用户不存在")
	}
	code := IdentifyCode(in.Email, in.AuthCode)
	if code == 401 {
		return &pb.UserToken{
			Code:  statuscode.Loginfailed,
			Token: "",
		}, errors.New("验证码错误")
	}
	token, err := utils.GetToken(u.ID, u.Role)
	if err != nil {
		return &pb.UserToken{
			Code:  statuscode.ERROR,
			Token: "",
		}, err
	}
	return &pb.UserToken{
		Code:  statuscode.OK,
		Token: token,
	}, nil
}

// 用户通过密码登录
func (*UserServer) LoginByPwd(ctx context.Context, in *pb.UserLoginImformationsByPwd) (*pb.UserToken, error) {
	u := dao.Search_user(in.Email)
	if u.ID == 0 {
		return &pb.UserToken{
			Code:  statuscode.Loginfailed,
			Token: "",
		}, errors.New("用户不存在")
	}
	s := utils.Encrypt(in.Pwd)
	if s != u.PassWord {
		return &pb.UserToken{
			Code:  statuscode.Loginfailed,
			Token: "",
		}, errors.New("密码错误")
	}
	token, err := utils.GetToken(u.ID, u.Role)
	if err != nil {
		return &pb.UserToken{
			Code:  statuscode.ERROR,
			Token: "",
		}, err
	}
	return &pb.UserToken{
		Code:  statuscode.OK,
		Token: token,
	}, nil
}

// 获取用户信息
func (*UserServer) GetUserInfo(ctx context.Context, in *pb.GETUserInfo) (*pb.GetUserInfoResponse, error) {
	u := dao.Search_user_by_id(in.Id)
	v := &pb.UserInfo{
		Uid:              u.ID,
		NickName:         u.NickName,
		Avatar:           u.Avatar,
		IndividualResume: u.IndividualResume,
		Role:             int32(u.Role),
		Email:            u.Email,
	}
	return &pb.GetUserInfoResponse{
		Code:     statuscode.OK,
		UserInfo: v,
	}, nil
}

// 更新用户信息
func (*UserServer) UpdateUserInfo(ctx context.Context, in *pb.UpdateUserInfoReq) (*pb.CommonResponse, error) {
	dao.Update_user(in.Id, in.Action, in.Data)
	return &pb.CommonResponse{Code: statuscode.OK, Msg: "更新成功"}, nil
}

// 发送邮件
func SendEmail(to string) error {
	subject := "【study_system】邮箱验证"
	html := fmt.Sprintf(`<div style="text-align: center;">
		<h2 style="color: #333;">欢迎使用，你的验证码为：</h2>
		<h1 style="margin: 1.2em 0;">%s</h1>
		<p style="font-size: 12px; color: #666;">请在5分钟内完成验证，过期失效，请勿告知他人，以防个人信息泄露</p>
	</div>`, CreateAuthCode(to))
	em := e.NewEmail()
	// 设置 sender 发送方 的邮箱 ， 此处可以填写自己的邮箱
	em.From = gen.C.Email.From

	// 设置 receiver 接收方 的邮箱  此处也可以填写自己的邮箱， 就是自己发邮件给自己
	em.To = []string{to}

	// 设置主题
	em.Subject = subject

	// 简单设置文件发送的内容，暂时设置成纯文本
	em.HTML = []byte(html)
	//fmt.Println(config.EmailAddr, config.Email, config.EmailAuth, config.EmailHost, config.EmailFrom)
	//设置服务器相关的配置
	err := em.Send(gen.C.Email.Addr, smtp.PlainAuth("", gen.C.Email.Email, gen.C.Email.Auth, gen.C.Email.Host))
	return err
}

// 创建验证码
func CreateAuthCode(em string) string {
	code := fmt.Sprintf("%d", utils.Randnum(900000)+100000)
	dao.SetAuthCode(em, code)
	return code
}

// 校验验证码
func IdentifyCode(em string, authcode string) int {
	res, err := dao.GetAuthCode(em)
	if err != nil {
		//logs.SugarLogger.Debugf("获取验证码失败:", err)
		return 401
	}
	if res != authcode {
		return 401
	}
	return 200
}
