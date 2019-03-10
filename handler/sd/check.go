package v1

import (
	"github.com/gin-gonic/gin"
)




func HealthCheck(c *gin.Context) {
	c.String(200, "ok")
}