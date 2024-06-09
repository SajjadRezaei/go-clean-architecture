package api

import (
	"fmt"
	"github.com/SajjadRezaei/go-clean-architecture/api/middleware"
	"github.com/SajjadRezaei/go-clean-architecture/api/routes"
	"github.com/SajjadRezaei/go-clean-architecture/api/validation"
	"github.com/SajjadRezaei/go-clean-architecture/config"
	"github.com/SajjadRezaei/go-clean-architecture/pkg/logging"
	"github.com/SajjadRezaei/go-clean-architecture/pkg/metrics"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/prometheus/client_golang/prometheus"
	_ "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
)

var logger = logging.NewLogger(config.GetConfig())

func InitServer(cfg *config.Config) {
	gin.SetMode(cfg.Server.RunMode)
	r := gin.New()

	RegisterValidators()
	RegisterPrometheus()

	r.Use(middleware.DefaultStructuredLogger(cfg))
	r.Use(middleware.Cors(cfg))
	r.Use(middleware.Prometheus())
	r.Use(gin.Logger(), gin.CustomRecovery(middleware.ErrorHandler) /*middleware.TestMiddleware()*/, middleware.LimitByRequest())

	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			health := v1.Group("/health")
			// Test
			testRouter := v1.Group("/test" /*middleware.Authentication(cfg), middleware.Authorization([]string{"admin"})*/)

			// User
			users := v1.Group("/users")

			// Test
			router.Health(health)
			router.TestRouter(testRouter)

			// User
			router.User(users, cfg)

		}

		//v1.GET("/health", func(c *gin.Context) {
		//	c.JSON(http.StatusOK, "Working!")
		//	return
		//})

		health := api.Group("/health")
		router.Health(health)

	}

	r.Run(fmt.Sprintf(""))
}

func RegisterValidators() {
	val, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		err := val.RegisterValidation("mobile", validation.IranianMobileNumberValidator, true)
		if err != nil {
			logger.Error(logging.Validation, logging.Startup, err.Error(), nil)
		}
		err = val.RegisterValidation("password", validation.PasswordValidator, true)
		if err != nil {
			logger.Error(logging.Validation, logging.Startup, err.Error(), nil)
		}
	}
}

//func RegisterSwagger(r *gin.Engine, cfg *config.Config) {
//	docs.SwaggerInfo.Title = "golang web api"
//	docs.SwaggerInfo.Description = "golang web api"
//	docs.SwaggerInfo.Version = "1.0"
//	docs.SwaggerInfo.BasePath = "/api"
//	docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%s", cfg.Server.ExternalPort)
//	docs.SwaggerInfo.Schemes = []string{"http"}
//
//	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
//}

func RegisterPrometheus() {
	err := prometheus.Register(metrics.DbCall)
	if err != nil {
		logger.Error(logging.Prometheus, logging.Startup, err.Error(), nil)
	}

	err = prometheus.Register(metrics.HttpDuration)
	if err != nil {
		logger.Error(logging.Prometheus, logging.Startup, err.Error(), nil)
	}
}
