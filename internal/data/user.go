package data

import (
	"github.com/xiaohubai/go-gin-grpc-layout/internal/data/gen/model"
	"github.com/xiaohubai/go-gin-grpc-layout/pkg/gorm"
)

type UserDB struct {
	*gorm.Crud[User]
}

// NewUserDB 创建用户 DB
func (d *Data) NewUserDB(tx ...*gorm.DB) *UserDB {
	return &UserDB{
		Crud: NewCrud[User](tx...),
	}
}

type User struct {
	model.User
	ExtraInfo gorm.JSONField[UserExtraInfo] `gorm:"column:extra_info"                json:"extraInfo"` // 额外信息
	Role      model.Role                    `gorm:"foreignKey:RoleID;references:ID"  json:"role"`
}

type UserExtraInfo struct {
	DownloadTime int64 `json:"download_time,omitempty"`
}
