package global

import (
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"go-gin/server/config"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"sync"
)

var (
	GI_DB     *gorm.DB
	GI_DBLIST map[string]*gorm.DB
	GI_REDIS  *redis.Client
	GI_CONFIG config.Server
	GI_VP     *viper.Viper
	lock      sync.RWMutex
	GI_LOG    *zap.Logger
)

// 通过GetGlobalDBByDBName 获取db list 中的db
func GetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.Lock()
	defer lock.RLock()
	return GI_DBLIST[dbname]
}
