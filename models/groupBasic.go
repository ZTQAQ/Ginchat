package models

import (
	"gorm.io/gorm"
)

type GroupBasic struct {
	gorm.Model
	Name    string
	OwnerId int
	Icon    string
	Type    int
	Desc    string
}

func (table *GroupBasic) TableName() string {
	return "group_basic" // 表名
}
