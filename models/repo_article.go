// Package models
// Time : 2023/8/6 16:05
// Author : PTJ
// File : repo_article
// Software: GoLand
package models

import "gorm.io/gorm"

type BizArticle struct {
	gorm.Model
	UserId        int64  `gorm:"not null;Index;comment:用户ID" json:"userId"`
	SubjectId     int64  `gorm:"Index;comment:主题ID" json:"subjectId"`
	Title         string `gorm:"Index;comment:文章标题" json:"title"`
	Intro         string `gorm:"type:text;comment:简介" json:"intro"`
	Cover         string `gorm:"comment:封面" json:"cover"`
	Content       string `gorm:"type:text;comment:文章内容" json:"content"`
	ContentMd     string `gorm:"type:text;comment:文章内容-markdown" json:"contentMd"`
	Code          string `gorm:"not null;comment:文章编码" json:"code"`
	KeyWords      int64  `gorm:"comment:关键字" json:"keyWords"`
	Sort          int64  `gorm:"comment:排序" json:"sort"`
	IsElite       int64  `gorm:"comment:是否推荐" json:"isElite"`
	Hits          int64  `gorm:"comment:浏览量" json:"hits"`
	ArticleStatus int64  `gorm:"tinyint(1);comment:文章状态0-发布1-草稿;" json:"articleStatus"`
}

func (*BizArticle) TableName() string {
	return "biz_article"
}
