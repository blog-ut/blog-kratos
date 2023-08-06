package service

import (
	"blog-kratos/helper"
	"blog-kratos/models"
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/metadata"
	"strings"

	pb "blog-kratos/api/blog"
)

type UserService struct {
	pb.UnimplementedUserServer
}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginReply, error) {
	var username string
	if md, ok := metadata.FromServerContext(ctx); ok {
		username = md.Get("username")
		fmt.Println(username)
	}
	su := new(models.SysUser)
	err := models.DB.Where("username = ? AND password = ?", req.Username, helper.GetMd5(req.Password)).Find(&su).Error
	if err != nil {
		return nil, err
	}
	//TODO 生成token
	token, err := helper.GenToken(su.Username, su.ID)
	if err != nil {
		return nil, err
	}
	return &pb.LoginReply{
		Token: strings.Join([]string{"Beaber", token}, " "),
	}, nil
}

func (s *UserService) RegisterUser(ctx context.Context, req *pb.RegisterUserRequest) (*pb.RegisterUserReply, error) {
	//查询用户是否存在
	var cnt int64
	if err := models.DB.Model(new(models.SysUser)).Where("username = ?", req.GetUserName()).Count(&cnt).Error; err != nil {
		return nil, err
	}
	user := &models.SysUser{
		Username: req.GetUserName(),
		Password: helper.GetMd5(req.Password),
		NickName: req.GetNickName(),
		Intro:    req.GetIntro(),
		Avatar:   req.GetAvatar(),
		Phone:    req.GetPhone(),
		Email:    req.GetEmail(),
		Address:  req.GetAddress(),
		IsEnable: req.GetIsEnable(),
		IsAdmin:  req.GetIsAdmin(),
	}
	err := models.DB.Create(user).Error
	if err != nil {
		return nil, err
	}
	return &pb.RegisterUserReply{}, nil
}
func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserReply, error) {
	return &pb.CreateUserReply{}, nil
}
func (s *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserReply, error) {
	return &pb.UpdateUserReply{}, nil
}
func (s *UserService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserReply, error) {
	return &pb.DeleteUserReply{}, nil
}
func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
	user := &models.SysUser{}
	db := models.DB.Model(user)
	if req.GetUserName() != "" {
		db.Where("username = ?", req.GetUserName())
	} else if req.GetPhone() != "" {
		db.Where("phone = ?", req.GetPhone())
	}
	if err := db.First(user).Error; err != nil {
		return nil, err
	}
	item := pb.UserItem{
		UserName: user.Username,
		NickName: user.NickName,
		Intro:    user.Intro,
		Avatar:   user.Avatar,
		Phone:    user.Phone,
		Email:    user.Email,
		Address:  user.Address,
		IsEnable: user.IsEnable,
		IsAdmin:  user.IsAdmin,
		CreateAt: user.CreatedAt.Unix(),
		UpdateAt: user.UpdatedAt.Unix(),
	}
	fmt.Println(user)
	return &pb.GetUserReply{
		User: &item,
	}, nil
}
func (s *UserService) ListUser(ctx context.Context, req *pb.ListUserRequest) (*pb.ListUserReply, error) {
	return &pb.ListUserReply{}, nil
}
