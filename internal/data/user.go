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

}
