package logic

import (
	"cloud-disk/core/define"
	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"
	"cloud-disk/core/models"
	"context"
	"fmt"
	"log/slog"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileListLogic {
	return &UserFileListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileListLogic) UserFileList(req *types.UserFileListRequest, userIdentity string) (resp *types.UserFileListReply, err error) {
	uf := make([]*types.UserFile, 0)
	resp = new(types.UserFileListReply)
	resp.List = uf
	size := req.Size
	if size == 0 {
		size = define.PageSize
	}
	page := req.Page
	if page == 0 {
		page = 1
	}
	offset := (page - 1) * size

	err = l.svcCtx.Engine.Table("user_repository").Where("parent_id = ? AND user_identity = ?", req.Id, userIdentity).
		Select("user_repository.id,user_repository.identity,user_repository.repository_identity,repository_pool.ext,repository_pool.path,repository_pool.size").Join(
		"LEFT", "repository_pool", "user_repository.repository_identity = repository_pool.identity").
		Limit(size, offset).Find(&uf)
	if err != nil {
		return nil, err
	}
	resp.List = uf
	slog.Info(fmt.Sprintf("%v ||||  %s", req, userIdentity))

	cout, err := l.svcCtx.Engine.Where("parent_id = ? AND user_identity = ?", req.Id, userIdentity).Count(new(models.UserRepository))
	resp.Count = cout
	return
}
