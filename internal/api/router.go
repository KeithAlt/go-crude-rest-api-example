package api

import (
	"context"
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/internal/service"
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/internal/service/repository"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Serve begins serving our resources
func Serve(db *repository.Client) {
	router := createRouter(db)
	server := createHTTPServer(router)

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	select {
	case <-ctx.Done():
		log.Println("Timing out in 5 seconds!")
	}
	log.Println("Server shutting down...")
}

// createRouter creates our routes & router
func createRouter(db *repository.Client) *gin.Engine {
	svc := &service.Products{
		Repo: db,
	}

	router := gin.Default()
	router.GET("/products", svc.FindAll)
	router.GET("/product/:guid", svc.Find)
	router.POST("/products", svc.Create)
	router.PUT("/products/:guid", svc.Update)
	router.PUT("/products/", svc.Update)
	router.DELETE("/products/:guid", svc.Delete)
	router.GET("/kill", Kill) // <- TODO add HEAVY auth
	return router
}

// Kill ungracefully shuts down our service
func Kill(ctx *gin.Context) {
	signal.NotifyContext(ctx, os.Interrupt)
}
