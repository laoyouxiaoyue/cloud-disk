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

type UserRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRegisterLogic) UserRegister(req *types.UserRegisterRequest) (resp *types.UserRegisterReply, err error) {
	code, err := l.svcCtx.RDB.Get(l.ctx, req.Email).Result()
	if err != nil {
		return nil, err
	}
	if code != req.Code {
		return nil, errors.New("验证码错误")
	}
	cnt, _ := l.svcCtx.Engine.Where("name = ?", req.Name).Count(new(models.UserBasic))
	if cnt > 0 {
		return nil, errors.New("用户名已存在")
	}
	user := &models.UserBasic{
		Name:     req.Name,
		Email:    req.Email,
		Password: helper.Md5(req.Password),
		Identity: helper.GetUUID(),
	}
	_, err = l.svcCtx.Engine.Insert(user)
	if err != nil {
		return nil, err
	}
	return
}
