package logic

import (
	"cloud-disk/core/define"
	"cloud-disk/core/helper"
	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type RefreshTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRefreshTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefreshTokenLogic {
	return &RefreshTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RefreshTokenLogic) RefreshToken(req *types.RefreshTokenRequest, oldtoken string) (resp *types.RefreshTokenReply, err error) {
	//
	uc, err := helper.AnalyzeToken(oldtoken)
	if err != nil {
		return nil, err
	}
	token, err := helper.GenerateToken(uc.Id, uc.Identity, uc.Name, int64(define.TokenExpire))
	if err != nil {
		return nil, err
	}
	refreshToken, err := helper.GenerateToken(uc.Id, uc.Identity, uc.Name, int64(define.RefreshTokenExpire))
	if err != nil {
		return nil, err
	}
	resp = &types.RefreshTokenReply{
		Token:        token,
		RefreshToken: refreshToken,
	}
	return
}
