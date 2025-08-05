package logic

import (
	"cloud-disk/core/helper"
	"cloud-disk/core/models"
	"context"
	"errors"

	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.LoginRequest) (resp *types.LoginReply, err error) {

	user := new(models.UserBasic)
	get, err := models.Engine.Where("name = ? AND password = ?", req.Name, helper.Md5(req.Password)).Get(user)

	if err != nil {
		return &types.LoginReply{Message: "服务器错误"}, errors.New("GetUserErr")
	}
	if !get {
		return &types.LoginReply{Message: "用户或密码错误"}, nil
	}
	token, err := helper.GenerateToken(int64(user.Id), user.Identity, user.Name)
	if err != nil {
		return nil, err
	}
	return &types.LoginReply{Message: token}, nil
}
