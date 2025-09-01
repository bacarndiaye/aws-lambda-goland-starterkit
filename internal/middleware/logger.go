package middleware

import (
    "log"
    "time"

    "github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
    return func(c *gin.Context) {
        start := time.Now()
        c.Next()
        rid := c.GetString(RequestIDHeader)
        latency := time.Since(start)
        status := c.Writer.Status()
        method := c.Request.Method
        path := c.Request.URL.Path
        log.Printf("rid=%s status=%d method=%s path=%s latency=%s", rid, status, method, path, latency)
    }
}

