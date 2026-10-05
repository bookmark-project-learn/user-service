package infrastructure

import (
	"github.com/bookmark-project-learn/user-service/constant"
	"github.com/bookmark-project-learn/user-service/internal/api"
	"github.com/bookmark-project-learn/user-service/internal/config"
	"github.com/bookmark-project-learn/user-service/internal/connection"
	"github.com/gin-gonic/gin"

	"github.com/bookmark-project-learn/bookmark-common-libs/pkg/sqldb"
)

func CreateApi() api.Engine {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic("Failed to load config: " + err.Error())
	}

	// create redis client
	rdClient := CreateRedisClient()

	// create SQL client
	sqlClient := CreateSqlClient()

	migrator := sqldb.BuildMigrate(sqlClient, constant.MigrationPath)
	migrator.SetLogging()
	err = migrator.MigrateUp()
	if err != nil {
		panic("start migration failed " + err.Error())
	}

	// connector
	connector := connection.NewDBConnector(rdClient, sqlClient)
	jwtGenerator, jwtValidator := CreateJwtProvider()
	apiEngine := api.NewEngine(&api.EnginOpt{
		App:          gin.New(),
		Cfg:          cfg,
		Connector:    connector,
		JwtGenerator: jwtGenerator,
		JwtValidator: jwtValidator,
	})
	return apiEngine
}
