package router

import (
    "github.com/gin-gonic/gin"

    "nexus-enedis-m23-measure-lambda/internal/handlers"
    "nexus-enedis-m23-measure-lambda/internal/middleware"
)

// New wires Gin engine, middleware, and routes.
func New() *gin.Engine {
    r := gin.New()

    // Core middleware
    r.Use(gin.Recovery())
    r.Use(middleware.RequestID())
    r.Use(middleware.Logger())
    r.Use(middleware.CORS())

    // Health endpoints
    r.GET("/healthz", handlers.Health)
    r.GET("/readyz", handlers.Ready)

    // v1 API group
    v1 := r.Group("/v1")
    {
        v1.GET("/hello/:name", handlers.Hello)
        v1.POST("/echo", handlers.Echo)
    }

    return r
}

