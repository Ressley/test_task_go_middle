package http

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/ressley/test_task_go_middle/http-grpc-client/http/handlers"
)

func StartServer() *http.Server {
	gin.SetMode("release")
	router := gin.New()
	port := fmt.Sprintf(":%d", 8080)

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"X-Requested-With", "Content-Type", "Authorization", "Sort", "Access-Control-Allow-Origin", "Access-Control-Allow-Methods", "Access-Control-Allow-Headers"},
		AllowCredentials: false,
		AllowOrigins:     []string{"*"},
		MaxAge:           12 * time.Hour,
	}))

	router.POST("/challenge", handlers.SendChallenge)
	router.POST("/message", handlers.SendMessage)

	server := &http.Server{
		Addr:           port,
		Handler:        router,
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   60 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Printf("[info] Http server listening on port %s", port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("[error] Failed to serve Http server: %v", err)
	}
	return server
}
