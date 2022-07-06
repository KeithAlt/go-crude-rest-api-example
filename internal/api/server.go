package api

import (
	"context"
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/config"
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
// TODO this function can be split into smaller functions
func Serve(cl *repository.Client) {
	svc := &Handler{
		Repo: &service.ProductRepository{Postgres: cl},
	}

	router := gin.Default()
	router.GET("/products", svc.FindAll)
	router.GET("/product/:guid", svc.Find)
	router.POST("/products", svc.Create)
	router.PUT("/product/:guid", svc.Update)
	router.PUT("/product/", svc.Update)
	router.DELETE("/products/:guid", svc.Delete)
	router.GET("/kill", Kill) // <- TODO add HEAVY auth

	srv := &http.Server{
		Addr:    config.Domain,
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}
	log.Println("Server exiting")
}

// Kill ungracefully shuts down our service
func Kill(ctx *gin.Context) {
	signal.NotifyContext(ctx, os.Interrupt)
}
