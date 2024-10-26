package rest

import (
	"log"
	"net/http"
	"sync"

	"github.com/ValkyrieKia/golang-deals-test-project/internal/controller/common"
	"github.com/ValkyrieKia/golang-deals-test-project/internal/controller/middleware"
	"github.com/ValkyrieKia/golang-deals-test-project/internal/controller/rest"
	"github.com/gin-gonic/gin"
)

func InitHttpServer(providers *common.HttpServerProviders) {
	r := gin.Default()
	cfg := *providers.Config

	r.Use(middleware.UseErrorHandler(providers.Config))

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// initialize app module routes
	rest.NewItemHttpController(r.Group("/item"), providers)

	// 🔥 Start the HTTP Server
	startServer(r, cfg.GenericConfig.AppPort)
}

func startServer(engine *gin.Engine, port string) {
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: engine,
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Printf("🚀 Server is listening on %s", srv.Addr)
		if err := srv.ListenAndServe(); err != nil {
			log.Fatalf("ListenAndServe: %v", err)
		}
	}()
	wg.Wait()
}
