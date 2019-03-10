package user

import (
	"net/http"
	"github.com/gin-gonic/gin"
)
// Index 显示该用户的全部锻炼
func Index(c *gin.Context) {

	c.JSON(http.StatusOK, "1")
}