package v1

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	token := "aaa"
	r := struct {
		token: token
		
	}
	c.JSON(http.StatusOK, c.H)
}

    GIT_COMMITTER_DATE="2018-06-09 19:23:08" \
    git commit . \
    -m"init program"