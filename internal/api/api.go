package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/bookmark-project-learn/bookmark-common-libs/middleware"
	"github.com/bookmark-project-learn/bookmark-common-libs/pkg/helpers/hasher"
	jwt_pkg "github.com/bookmark-project-learn/bookmark-common-libs/pkg/jwt"
	"github.com/bookmark-project-learn/user-service/docs"
	_ "github.com/bookmark-project-learn/user-service/docs"

	"github.com/bookmark-project-learn/user-service/internal/config"
	"github.com/bookmark-project-learn/user-service/internal/connection"
	health_check_handler "github.com/bookmark-project-learn/user-service/internal/handler/health_check"
	user_handler "github.com/bookmark-project-learn/user-service/internal/handler/user"
	health_check_repository "github.com/bookmark-project-learn/user-service/internal/repository/health_check"
	userRepository "github.com/bookmark-project-learn/user-service/internal/repository/user"
	health_check_service "github.com/bookmark-project-learn/user-service/internal/service/health_check"
	user_service "github.com/bookmark-project-learn/user-service/internal/service/user"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Engine interface for app engine
type Engine interface {
	Run() error
	ServeHTTP(w http.ResponseWriter, req *http.Request)
}

// engine struct for app engine
type engine struct {
	app          *gin.Engine
	cfg          *config.Config
	connector    connection.DBConnector
	jwtGenerator jwt_pkg.JwtGenerator
	jwtValidator jwt_pkg.JwtValidator
}

type EnginOpt struct {
	App          *gin.Engine
	Cfg          *config.Config
	Connector    connection.DBConnector
	JwtGenerator jwt_pkg.JwtGenerator
	JwtValidator jwt_pkg.JwtValidator
}

// NewEngine creates a new engine instance
func NewEngine(opt *EnginOpt) Engine {
	api := &engine{
		app:          opt.App,
		cfg:          opt.Cfg,
		connector:    opt.Connector,
		jwtGenerator: opt.JwtGenerator,
		jwtValidator: opt.JwtValidator,
	}

	api.initRoutes(opt.Cfg)
	return api
}

// config Run starts the app engine
func (e *engine) Run() error {
	return e.app.Run(fmt.Sprintf(":%s", e.cfg.AppPort))
}

// override config ServeHTTP serves the app engine
func (e *engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	e.app.ServeHTTP(w, req)
}

type handlers struct {
	healthCheck health_check_handler.HealthCheck
	user        user_handler.UserHandler
	config      *config.Config
}

func (e *engine) InitHandlers(cfg *config.Config) handlers {
	serviceName := cfg.ServiceName
	instanceID := cfg.InstanceID
	redisClient := e.connector.GetRedisClient()
	sqlDB := e.connector.GetSqlDB()
	// create helper
	hasher := hasher.NewHasher()
	// create repository
	healthCheckRepository := health_check_repository.NewPing(redisClient)

	userRepository := userRepository.NewUserRepository(sqlDB)
	// create service
	healthCheckService := health_check_service.NewHealthCheck(serviceName, instanceID, healthCheckRepository)
	userService := user_service.NewUserService(userRepository, hasher, e.jwtGenerator)

	// create handler
	healthCheckHandler := health_check_handler.NewHealthCheck(healthCheckService)
	userHandler := user_handler.NewUserHandler(userService)

	return handlers{healthCheckHandler, userHandler, cfg}
}

func (e *engine) initRoutes(cfg *config.Config) {
	allHandlers := e.InitHandlers(cfg)

	e.app.GET("/health-check", allHandlers.healthCheck.Ping)

	docs.SwaggerInfo.BasePath = allHandlers.config.BasePath
	e.app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	jwtMiddleware := middleware.NewJwtAuthMiddleware(e.jwtValidator)

	v1Routes := e.app.Group("/v1")
	{
		// --- Public API
		v1Routes.POST("/users/login", allHandlers.user.Login)
		v1Routes.POST("/users/register", allHandlers.user.Register)

		// -- Private Api
		v1Routes.Use(jwtMiddleware.JwtAuth()) // middelware
		v1Routes.GET("/self/info", allHandlers.user.GetUserInfo)
		v1Routes.PUT("/self/info", allHandlers.user.UpdateUserInfo)

	}
}
