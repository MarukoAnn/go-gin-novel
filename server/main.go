package main

import (
	"fmt"
	"go-gin/server/core"
	"go-gin/server/global"
	"go-gin/server/initialize"
	"go-gin/server/models/common"
)

func main() {
	//
	global.GI_VP = core.Viper()
	fmt.Println(global.GI_VP.AllKeys())
	// 初始化数据库链接
	global.GI_DB = initialize.Gorm()
	data := common.NovelInfo{}
	if global.GI_DB != nil {
		initialize.RegisterTables() // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.GI_DB.DB()
		defer db.Close()
	}
	res := global.GI_DB.Limit(10).Find(&data)
	if res.Error != nil || res.RowsAffected <= 0 {
		fmt.Println("错误")
	}
	fmt.Println(data)
}
