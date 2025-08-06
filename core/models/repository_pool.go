package models

import "time"

type RepositoryPool struct {
	Id          int
	Identity    string
	Hash        string
	Name        string
	Ext         string
	Size        int64
	Path        string
	CreatedTime time.Time `xorm:"created_at"`
	UpdatedTime time.Time `xorm:"updated_at"`
	DeletedTime time.Time `xorm:"deleted_at"`
}

func (r *RepositoryPool) TableName() string {
	return "repository_pool"
}
