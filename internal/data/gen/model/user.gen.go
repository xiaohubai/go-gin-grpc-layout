// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameUser = "user"

// User 用户表
type User struct {
	ID        int64          `gorm:"column:id;primaryKey;autoIncrement:true;comment:ID" json:"id"`             // ID
	UID       int64          `gorm:"column:uid;not null;comment:用户ID" json:"uid"`                              // 用户ID
	Username  string         `gorm:"column:username;not null;comment:用户名" json:"username"`                     // 用户名
	Password  string         `gorm:"column:password;not null;default:123456;comment:密码" json:"password"`       // 密码
	Salt      string         `gorm:"column:salt;not null;default:abcdef;comment:加盐" json:"salt"`               // 加盐
	RoleID    int32          `gorm:"column:role_id;not null;comment:角色Id" json:"role_id"`                      // 角色Id
	Phone     string         `gorm:"column:phone;not null;comment:手机号" json:"phone"`                           // 手机号
	Email     string         `gorm:"column:email;not null;comment:邮箱" json:"email"`                            // 邮箱
	CreatedAt time.Time      `gorm:"column:created_at;not null;autoCreateTime;comment:创建时间" json:"created_at"` // 创建时间
	UpdatedAt time.Time      `gorm:"column:updated_at;not null;autoUpdateTime;comment:更新时间" json:"updated_at"` // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;comment:删除时间" json:"deleted_at"`                         // 删除时间
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
