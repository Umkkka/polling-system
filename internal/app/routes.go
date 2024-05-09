package app

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
	slokmetrics "github.com/slok/go-http-metrics/metrics/prometheus"
	"github.com/slok/go-http-metrics/middleware"
	ginmiddleware "github.com/slok/go-http-metrics/middleware/gin"

	"polling-system/internal/config"
	"polling-system/internal/transport/handler"
)

// initRouter - инициализирует gin router
func initRouter(config *config.Config) *gin.Engine {
	if config.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	engine := gin.New()
	engine.Use(gin.Recovery())

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPatch},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	metricsMiddleware := middleware.New(middleware.Config{
		Recorder: slokmetrics.NewRecorder(slokmetrics.Config{}),
	})

	engine.Use(ginmiddleware.Handler("", metricsMiddleware), corsMiddleware)

	engine.NoMethod(func(ctx *gin.Context) {
		ctx.JSON(http.StatusMethodNotAllowed, "method not allowed")
	})
	engine.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusMethodNotAllowed, "no route found")
	})

	return engine
}

func setupTechnicalRoutes(engine *gin.Engine) {
	engine.GET("/technical/ping", func(ctx *gin.Context) {
		ctx.JSON(
			http.StatusOK,
			gin.H{"success": true, "result": "pong"},
		)
	})

	engine.GET("/metrics", func(ctx *gin.Context) {
		promhttp.HandlerFor(prometheus.DefaultGatherer, promhttp.HandlerOpts{DisableCompression: true}).
			ServeHTTP(ctx.Writer, ctx.Request)
	})
}

// setupApiRoutes - определяет api группу и ее пути
func setupApiRoutes(engine *gin.Engine, pollHandler *handler.Handler) {
	api := engine.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.POST("/poll", pollHandler.CreatePoll)
		}
	}
}
