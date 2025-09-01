package handlers

import (
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
)

func Health(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "status":  "ok",
        "time":    time.Now().UTC(),
    })
}

func Ready(c *gin.Context) {
    // Extend with checks (DB, downstream, etc.)
    c.JSON(http.StatusOK, gin.H{
        "ready":   true,
        "time":    time.Now().UTC(),
    })
}

