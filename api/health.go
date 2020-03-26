package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Health(c *gin.Context) {
	u, _ := uuid.NewRandom()
	s := u.String()
	c.JSON(http.StatusOK, gin.H{
		"ping": s,
	})
}
