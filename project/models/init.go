package models

import (
	"diy-server/utils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"os"
	"xorm.io/core"
)

type Address struct {
	Id       int    `xorm:"autoincr pk" json:"id,omitempty"`
	UserId   int    `xorm:"notnull default(0)" json:"user_id,omitempty"`
	Name     string `xorm:"varchar(60) notnull default('')" json:"name,omitempty"`
	Tel      string `xorm:"varchar(11) not null default('')" json:"tel,omitempty"`
	Province string `xorm:"varchar(90) default('')" json:"province,omitempty"`
	City     string `xorm:"varchar(90) default('')" json:"city,omitempty"`
	Region   string `xorm:"varchar(90) default('')" json:"region,omitempty"`
	Detail   string `xorm:"text" json:"detail,omitempty"`
	Sex      int8   `xorm:"tinyint(1) default(1)" json:"sex,omitempty"`
}

type Admin struct {
	Id           int    `xorm:"autoincr pk"`
	Username     string `xorm:"varchar(20) notnull default('')" json:"username,omitempty"`
	Password     string `xorm:"varchar(32) notnull default('')" json:"password,omitempty"`
	LastLogin    string `xorm:"varchar(11) not null default('0')" json:"last_login,omitempty"`
	CurrentLogin string `xorm:"varchar(11) not null default('0')" json:"current_login,omitempty"`
}

type Goods struct {
	Id         int     `xorm:"autoincr pk" json:"id,omitempty"`
	Name       string  `xorm:"varchar(30) default('')" json:"name,omitempty"`
	Price      float64 `xorm:"float(6,2) default(0)" json:"price,omitempty"`
	Status     int     `xorm:"tinyint(1) default(1)" json:"status,omitempty"`
	CreateTime string  `xorm:"varchar(11) default('0')" json:"create_time,omitempty"`
}

type GoodsImg struct {
	Id       int    `xorm:"autoincr pk" json:"id,omitempty"`
	GoodsId  int    `xorm:"notnull default(0)" json:"goods_id,omitempty"`
	Front    string `xorm:"varchar(100) default('')" json:"front,omitempty"`
	Contrary string `xorm:"varchar(100) default('')" json:"contrary,omitempty"`
	Color    string `xorm:"varchar(30) default('')" json:"color,omitempty"`
}

type Orders struct {
	Id                int     `xorm:"autoincr pk" json:"id,omitempty"`
	UserId            int     `xorm:"notnull default(0)" json:"user_id,omitempty"`
	GoodsId           int     `xorm:"notnull default(0)" json:"goods_id,omitempty"`
	GoodsName         string  `xorm:"notnull default('')" json:"goods_name,omitempty"`
	Front             string  `xorm:"varchar(100) default('')" json:"front,omitempty"`
	Contrary          string  `xorm:"varchar(100) default('')" json:"contrary,omitempty"`
	FrontPicture      string  `xorm:"varchar(100) default('')" json:"front_picture,omitempty"`
	ContraryPicture   string  `xorm:"varchar(100) default('')" json:"contrary_picture,omitempty"`
	FrontCompose      string  `xorm:"varchar(100) default('')" json:"front_compose,omitempty"`
	ContraryCompose   string  `xorm:"varchar(100) default('')" json:"contrary_compose,omitempty"`
	Color             string  `xorm:"varchar(30) default('')" json:"color,omitempty"`
	Size              string  `xorm:"varchar(1) default('')" json:"size,omitempty"`
	OutTradeNo        string  `xorm:"varchar(64) default('')" json:"out_trade_no,omitempty"`
	TransactionId     string  `xorm:"varchar(64) default('')" json:"transaction_id,omitempty"`
	Total             float64 `xorm:"float(6,2) default(0)" json:"total,omitempty"`
	Description       string  `xorm:"text" json:"id,omitempty"`
	ConsigneeName     string  `xorm:"varchar(60) default('')" json:"consignee_name,omitempty"`
	ConsigneeTel      string  `xorm:"varchar(11) default('')" json:"consignee_tel,omitempty"`
	ConsigneeSex      string  `xorm:"tinyint(1) default(1)" json:"consignee_sex,omitempty"`
	ConsigneeProvince string  `xorm:"varchar(90) default('')" json:"consignee_province,omitempty"`
	ConsigneeCity     string  `xorm:"varchar(90) default('')" json:"consignee_city,omitempty"`
	ConsigneeRegion   string  `xorm:"varchar(90) default('')" json:"consignee_region,omitempty"`
	ConsigneeDetail   string  `xorm:"text" json:"consignee_detail,omitempty"`
	Note              string  `xorm:"text" json:"note,omitempty"`
	CreateTime        string  `xorm:"varchar(11) default('0')" json:"create_time,omitempty"`
	PayTime           string  `xorm:"varchar(11) default('0')" json:"pay_time,omitempty"`
	DispatchinTime    string  `xorm:"varchar(11) default('0')" json:"dispatchin_time,omitempty"`
	CompleteTime      string  `xorm:"varchar(11) default('0')" json:"complete_time,omitempty"`
	Status            int     `xorm:"tinyint(1) default(0)" json:"status,omitempty"`
}

type Picture struct {
	Id         int    `xorm:"autoincr pk" json:"id,omitempty"`
	Src        string `xorm:"varchar(120) default('')" json:"src,omitempty"`
	Name       string `xorm:"varchar(90) default('')" json:"name,omitempty"`
	CreateTime string `xorm:"varchar(11) default('0')" json:"create_time,omitempty"`
}

type Users struct {
	Id           int    `xorm:"autoincr pk" json:"id,omitempty"`
	HeadPic      string `xorm:"varchar(200) notnull default('')" json:"head_pic,omitempty"`
	Nickname     string `xorm:"varchar(50) notnull default('')" json:"nickname,omitempty"`
	Openid       string `xorm:"varchar(50) notnull default('')" json:"openid,omitempty"`
	Unionid      string `xorm:"varchar(50) notnull default('')" json:"unionid,omitempty"`
	SessionKey   string `xorm:"varchar(100) notnull default('')" json:"session_key,omitempty"`
	RegisterTime string `xorm:"varchar(11) default('0')" json:"register_time,omitempty"`
}

// 日志
var log = utils.Log

// xorm引擎
var Engine *xorm.Engine

func init() {
	var err error
	Engine, err = xorm.NewEngine("mysql", "root:root@tcp(127.0.0.1:3306)/diy?charset=utf8")

	if err != nil {
		log.Error(err.Error())
		return
	}
	// 设置xorm日志
	f, err := os.Create("logs/xorm.log")
	if err != nil {
		println(err.Error())
		return
	}
	Engine.SetLogger(xorm.NewSimpleLogger(f))

	// 延迟关闭不能用，否则会报 database is closed！
	//defer Engine.Close()

	err2 := Engine.Ping()
	if err2 != nil {
		// 日志输出错误
		log.Error(err2.Error())
		return
	}
	// 打印sql语句
	Engine.ShowSQL(true)

	// 设置日志等级
	Engine.Logger().SetLevel(core.LOG_DEBUG)

	// 设置映射方式
	tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, "")
	Engine.SetTableMapper(tbMapper)

	// 创建表
	err3 := Engine.Sync2(new(Users), new(Admin), new(Orders), new(Goods), new(GoodsImg), new(Picture), new(Address))
	if err3 != nil {
		// 日志输出错误
		log.Error(err3.Error())
		return
	}

	// 判断是否有超级管理员，如果没有则进行创建
	var admin Admin
	_, err5 := Engine.Id(1).Get(&admin)
	if err5 != nil {
		log.Error(err5.Error())
	}
	// 不存在进行创建
	if admin.Id <= 0 {
		var admin = Admin{
			Id:       1,
			Username: "admin",
			Password: "e10adc3949ba59abbe56e057f20f883e",
		}
		_, err6 := Engine.Insert(&admin)
		if err6 != nil {
			log.Error(err6.Error())
		}
	}
}
