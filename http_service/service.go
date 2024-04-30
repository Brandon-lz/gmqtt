package httpservice

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	// "os"
	// "os/signal"
	// "syscall"

	"github.com/Brandon-lz/gmqtt/config"
	"github.com/Brandon-lz/gmqtt/http_service/routers"
	"github.com/Brandon-lz/gmqtt/utils"

	"github.com/gin-gonic/gin"

	_ "github.com/Brandon-lz/gmqtt/docs" // 引入文档目录

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	// CORS for https://foo.com and https://github.com origins, allowing:
	// - PUT and PATCH methods
	// - Origin header
	// - Credentials share
	// - Preflight requests cached for 12 hours
	router.Use(ginCors())
	router.Use(gin.CustomRecovery(ErrorHandler))
	router.Use(StructuredLogger())
	return router
}

func Start() {
	router := InitRouter()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	})

	router.GET("/health", healthCheck)

	port := config.Config.Service.Port

	url := ginSwagger.URL(fmt.Sprintf("http://localhost:%d/docs/doc.json",config.Config.Service.Port)) // The URL pointing to API definition

	// swagger:  http://localhost:8080/docs/index.html
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	v1 := router.Group("/api/v1")
	routers.RegisterRoutes(v1)

	startAndlisten(router, port)

	// router.Run("0.0.0.0:8080")

}

func startAndlisten(router *gin.Engine, port int) {

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: router,
	}
	// Graceful stop
	go func() {
		defer utils.RecoverAndLog()
		// service connections
		slog.Info(fmt.Sprintf("Start Server At Port: %d", port))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error(fmt.Sprintf("listen: %s\n", err))
			os.Exit(1)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	// quit := make(chan os.Signal)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	// signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	quit := make(chan os.Signal)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	slog.Info("Shutdown Server ...")

	// Wait for service stop
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err == nil {
		slog.Info("Server exiting")
		os.Exit(0)
	} else {
		slog.Error("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	<-ctx.Done()
	log.Println("timeout of 5 seconds.")
	log.Println("Server exiting")
}

// healthCheck 路由
// @Summary  healthCheck 路由
// @Description  healthCheck 路由
// @Tags     default
// @Accept   json
// @Produce  json
// @Success  200  {string}  pong  "pong"
// @Router   /health [get]
func healthCheck(c *gin.Context) {
	// c.Header("Content-Type", "charset=utf-8")
	c.String(200, "I am healthy")
}
