package initialize

import (
	"fmt"
	"go-gin/server/global"
	"gorm.io/gorm"
)

func GormMysql() *gorm.DB {
	m := global.GI_CONFIG.Mysql
	if m.Dbname == "" {
		return nil
	}
	db_config := m.Dsn()
	fmt.Println("xx", db_config)
	//db, err := gorm.Open(mysql.Open(db_config), &gorm.Config{})
	//if err != nil {
	//	println("db")
	//	panic("failed to connect mysql")
	//}
	return nil
}
