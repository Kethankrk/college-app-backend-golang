package students

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func StudentLogin(c *gin.Context) {
	c.String(http.StatusOK, "hello")
}
