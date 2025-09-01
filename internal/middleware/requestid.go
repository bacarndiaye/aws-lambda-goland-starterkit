package middleware

import (
    "crypto/rand"
    "encoding/hex"

    "github.com/gin-gonic/gin"
)

const RequestIDHeader = "X-Request-ID"

func RequestID() gin.HandlerFunc {
    return func(c *gin.Context) {
        rid := c.GetHeader(RequestIDHeader)
        if rid == "" {
            rid = newID()
        }
        c.Writer.Header().Set(RequestIDHeader, rid)
        c.Set(RequestIDHeader, rid)
        c.Next()
    }
}

func newID() string {
    b := make([]byte, 16)
    _, _ = rand.Read(b)
    return hex.EncodeToString(b)
}

