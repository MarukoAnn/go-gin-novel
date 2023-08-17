package core

import (
	"fmt"
	"go-gin/server/global"
	"go-gin/server/utils"
	"go.uber.org/zap"
	"os"
)

func Zap(logger *zap.Logger) {
	if ok, _ := utils.PathExists(global.GI_CONFIG.Zap.Director); !ok {
		fmt.Printf("create %v director\n", global.GI_CONFIG.Zap.Director)
		// os.ModePerm 最高权限
		_ = os.Mkdir(global.GI_CONFIG.Zap.Director, os.ModePerm)
	}

	//cores := inter
}
