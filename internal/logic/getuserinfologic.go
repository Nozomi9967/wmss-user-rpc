package logic

import (
	"WMSS/user/rpc/model"
	"context"

	"WMSS/user/rpc/internal/svc"
	"WMSS/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取用户详情信息
func (l *GetUserInfoLogic) GetUserInfo(in *__.GetUserRequest) (*__.GetUserResponse, error) {
	var sysUser *model.SysUser
	var err error
	sysUser, err = l.svcCtx.SysUserModel.FindOne(l.ctx, in.UserId)
	if err != nil {
		return nil, err
	}
	return &__.GetUserResponse{
		UserId:       sysUser.UserId,
		UserName:     sysUser.UserName,
		RealName:     sysUser.RealName,
		RoleId:       sysUser.RoleId,
		Department:   sysUser.Department.String,
		Position:     sysUser.Position.String,
		ContactPhone: sysUser.ContactPhone.String,
		UserStatus:   string(sysUser.UserStatus),
		//LastLoginTime:      ,
		//PasswordExpireTime: "",
		CreateTime: sysUser.CreateTime.Format("2006-01-02 15:04:05"),
		UpdateTime: sysUser.UpdateTime.Format("2006-01-02 15:04:05"),
	}, nil
}
