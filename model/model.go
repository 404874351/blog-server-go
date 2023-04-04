package model

import (
	"blog-server-go/conf"
	"blog-server-go/utils"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var DB *gorm.DB

type Model struct {
	// id
	ID        	int 		        `json:"id"          gorm:"primary_key"`
	// 创建时间
	CreateTime 	utils.SystemTime	`json:"createTime"`
	// 更新时间
	UpdateTime 	utils.SystemTime	`json:"updateTime"`
	// 逻辑删除 数据禁用标记
	Deleted 	int8 		        `json:"deleted"     gorm:"index"`
}

const (
	MODEL_ACTIVED       int8 = 0
	MODEL_DEACTIVED     int8 = 1
)

//
// init 初始化DB
//  @Description: 初始化数据库连接
//
func init() {
	// 拼接数据源
	path := conf.GlobalConfig.Mysql.Username +
		":" + conf.GlobalConfig.Mysql.Password +
		"@" + conf.GlobalConfig.Mysql.Url
	// 尝试连接
	var err error
	DB, err = gorm.Open("mysql", path)
	if err != nil {
		panic(err)
	}

	// 表名禁用复数
	DB.SingularTable(true)
	// 自定义时间更新插件
	DB.Callback().Create().Replace("gorm:update_time_stamp", fillTimeStampForCreate)
	DB.Callback().Update().Replace("gorm:update_time_stamp", fillTimeStampForUpdate)
	// 连接数配置
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)

}

//
// fillTimeStampForCreate
//  @Description: 插入时自动填充时间戳
//  @param scope
//
func fillTimeStampForCreate(scope *gorm.Scope) {
	if scope.HasError() {
		return
	}
	now := utils.SystemTime{Time: time.Now()}
	if createTimeField, ok := scope.FieldByName("CreateTime"); ok {
		err := createTimeField.Set(now)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	if updateTimeField, ok := scope.FieldByName("UpdateTime"); ok {
		err := updateTimeField.Set(now)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

}

//
// fillTimeStampForUpdate
//  @Description: 更新时自动填充时间戳
//  @param scope
//
func fillTimeStampForUpdate(scope *gorm.Scope)  {
	if scope.HasError() {
		return
	}
	now := utils.SystemTime{Time: time.Now()}
	if _, ok := scope.FieldByName("UpdateTime"); ok {
		//err := updateTimeField.Set(now)
		err := scope.SetColumn("UpdateTime", now)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

}

