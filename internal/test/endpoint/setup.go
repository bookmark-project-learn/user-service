package user_endpoint

import (
	"testing"

	jwt_pkg "github.com/bookmark-project-learn/bookmark-common-libs/pkg/jwt"
	"github.com/bookmark-project-learn/user-service/internal/api"
	"github.com/bookmark-project-learn/user-service/internal/config"
	"github.com/bookmark-project-learn/user-service/internal/connection"
	"github.com/bookmark-project-learn/user-service/internal/test/data/fixture"
	"github.com/gin-gonic/gin"
)

func BuildApiEngine(t *testing.T, fix fixture.Fixture, engOpt *api.EnginOpt) api.Engine {
	connectorMock, errConnector := connection.InitDBConnectorMock(t, fix)
	if errConnector != nil {
		t.Fatal(errConnector)
	}
	engOpt.App = gin.New()
	engOpt.Connector = connectorMock
	apiEngine := api.NewEngine(engOpt)
	return apiEngine
}

func BuildUserHandlerFull(testItem *testing.T, cfg *config.Config) api.Engine {
	db := fixture.NewUserTestCase(testItem)
	jwtMock := jwt_pkg.NewMockJwt()
	jwtGenerator := jwtMock.JwtGenarate
	jwtValidator := jwtMock.JwtValidate
	apiEngine := BuildApiEngine(testItem, db, &api.EnginOpt{
		Cfg:          cfg,
		JwtGenerator: jwtGenerator,
		JwtValidator: jwtValidator,
	})
	return apiEngine

}
