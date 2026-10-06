package infrastructure

import (
	redisPkg "github.com/bookmark-project-learn/bookmark-common-libs/pkg/redis"
	"github.com/bookmark-project-learn/bookmark-common-libs/pkg/sqldb"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func CreateSqlClient() *gorm.DB {
	sqlClient, err := sqldb.NewSqlDB()
	if err != nil {
		panic(err)
	}
	return sqlClient
}
func CreateRedisClient() *redis.Client {
	rdClient, err := redisPkg.NewRedisClient()
	if err != nil {
		panic(err)
	}
	return rdClient
}
