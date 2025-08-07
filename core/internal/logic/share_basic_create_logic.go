package logic

import (
	"cloud-disk/core/helper"
	"cloud-disk/core/models"
	"context"

	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShareBasicCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShareBasicCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareBasicCreateLogic {
	return &ShareBasicCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShareBasicCreateLogic) ShareBasicCreate(req *types.ShareBasicCreateRequest, userIdentity string) (resp *types.ShareBasicCreateReply, err error) {
	data := &models.ShareBasic{
		Identity:           helper.GetUUID(),
		UserIdentity:       userIdentity,
		RepositoryIdentity: req.RepositoryIdentity,
		ExpiredTime:        req.ExpiredTime,
	}
	_, err = l.svcCtx.Engine.Insert(data)
	if err != nil {
		return nil, err
	}
	resp = &types.ShareBasicCreateReply{
		Identity: data.Identity,
	}
	return
}
