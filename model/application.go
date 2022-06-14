package model

import (
	"gorm.io/gorm"
)

//const (
//	Accept = 1 //接收申请
//	Reject = 0 //拒绝申请
//)

type Application struct {
	gorm.Model

	//Uid       uint   `json:"uid"`
	Username  string `json:"username"`
	Sid       int    `json:"sid"` //档案id
	Grayscale string `json:"grayscale"`
	Reason    string `json:"reason"`
	Flag      *bool  `json:"flag" gorm:"default:false"` //true表示审批通过，false表示未通过
	Reviewer  string `json:"reviewer"`                  //负责审批的员工姓名
}

type CreateApplicationInput struct {
	Username  string `json:"username" binding:"required"`
	Sid       int    `json:"sid" binding:"required"`
	Grayscale string `json:"grayscale" binding:"required"`
	Reason    string `json:"reason"`
}

type UpdateApplicationInput struct {
	//Id       uint     `json:"id" binding:"required"`
	Flag     *bool  `json:"flag" structs:"flag"`
	Reviewer string `json:"reviewer" binding:"required" structs:"reviewer"`
}
