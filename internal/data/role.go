package data

import (
	"github.com/xiaohubai/go-gin-grpc-layout/internal/data/gen/model"
	"github.com/xiaohubai/go-gin-grpc-layout/pkg/gorm"
)

type RoleDB struct {
	*gorm.Crud[Role]
}

// NewUserDB 创建用户 DB
func (d *Data) NewRoleDB(tx ...*gorm.DB) *RoleDB {
	return &RoleDB{
		Crud: NewCrud[Role](tx...),
	}
}

type Role struct {
	model.Role
	ExtraInfo gorm.JSONField[RoleExtraInfo] `gorm:"column:extra_info"                json:"extraInfo"` // 额外信息
}

type RoleExtraInfo struct {
	DownloadTime int64 `json:"download_time,omitempty"`
}
