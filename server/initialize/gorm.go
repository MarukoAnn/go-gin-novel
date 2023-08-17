package initialize

import (
	"fmt"
	"go-gin/server/global"
	"go-gin/server/models/common"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"os"
)

func Gorm() *gorm.DB {
	fmt.Println("xx-11", global.GI_CONFIG.System.DbType)
	switch global.GI_CONFIG.System.DbType {
	case "mysql":
		return GormMysql()
	default:
		return GormMysql()
	}
}

func RegisterTables() {
	db := global.GI_DB
	err := db.AutoMigrate(
		common.NovelInfo{},
	)
	if err != nil {
		global.GI_LOG.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	global.GI_LOG.Info("register table success")
}
