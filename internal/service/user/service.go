package user

import (
	"context"
	pb "studysystem_micro/idl/user"
)

func (*UserServer) Register(ctx context.Context, in *pb.UserRegister) (*pb.CommonResponse, error) {
	return &pb.CommonResponse{Code: 200, Msg: "success"}, nil
}
func (*UserServer) GetAuthCode(ctx context.Context, in *pb.UserGetAuthCode) (*pb.CommonResponse, error) {
	return nil, nil
}
func (*UserServer) Login(ctx context.Context, in *pb.UserLoginImformationsByAuthCode) (*pb.UserToken, error) {
	return nil, nil
}
func (*UserServer) LoginByPwd(ctx context.Context, in *pb.UserLoginImformationsByPwd) (*pb.UserToken, error) {
	return nil, nil
}
func (*UserServer) GetUserInfo(ctx context.Context, in *pb.UserInfo) (*pb.GetUserInfoResponse, error) {
	return nil, nil
}
func (*UserServer) UpdateUserInfo(ctx context.Context, in *pb.UpdateUserInfoReq) (*pb.CommonResponse, error) {
	return nil, nil
}
