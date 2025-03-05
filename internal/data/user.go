package data

import "github.com/xiaohubai/go-gin-grpc-layout/pkg/gorm"

type UserDB struct {
	*gorm.Crud[User]
}

// NewUserDB 创建用户 DB
func NewUserDB(tx ...*gorm.DB) *UserDB {
	return &UserDB{
		Crud: NewCrud[User](tx...),
	}
}

type User struct {
	ID       uint   `gorm:"primary_key"`
	Name     string `json:"name" gorm:"column:name;type:varchar(255);not null;comment:'用户名'"`
	Password string `json:"password" gorm:"column:password;type:varchar(255);not null;comment:'密码'"`
}
