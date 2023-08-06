// Package models
// Time : 2023/8/6 15:54
// Author : PTJ
// File : repo_basic
// Software: GoLand
package models

import "gorm.io/gorm"

type SysUser struct {
	gorm.Model
	Username string `gorm:"not null;type:varchar(255);uniqueIndex;comment:用户名" json:"username"`
	NickName string `gorm:"not null;comment:昵称" json:"nickName"`
	Password string `gorm:"not null;comment:密码" json:"password"`
	Intro    string `gorm:"comment:简介" json:"intro"`
	Avatar   string `gorm:"comment:头像" json:"avatar"`
	Phone    string `gorm:"comment:电话号码" json:"phone"`
	Email    string `gorm:"comment:邮箱" json:"email"`
	Address  string `gorm:"comment:地址" json:"address"`
	IsEnable uint32 `gorm:"not null;type:tinyint(1);comment:是否禁用 0无效 1有效" json:"isEnable"`
	IsAdmin  uint32 `gorm:"not null;type:tinyint(1);comment:是否是管理员 0无效 1有效" json:"isAdmin"`
}

func (table *SysUser) TableName() string {
	return "sys_user"
}
