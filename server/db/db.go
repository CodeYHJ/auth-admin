package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"sync"
)

var (
	Once     sync.Once
	pgdbpool *gorm.DB
	pgerr    error
)

func InitDB() {
	Once.Do(func() {
		dsn := os.Getenv("DB_URI")
		pgdbpool, pgerr = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if pgerr != nil {
			panic(fmt.Errorf("连接数据库出错:%v", pgerr))
		}
	})

}

func GetDB() *gorm.DB {
	return pgdbpool
}
