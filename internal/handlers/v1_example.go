package handlers

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {
    name := c.Param("name")
    if name == "" {
        name = "world"
    }
    c.JSON(http.StatusOK, gin.H{"message": "Hello " + name})
}

func Echo(c *gin.Context) {
    var payload map[string]any
    if err := c.BindJSON(&payload); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"received": payload})
}

