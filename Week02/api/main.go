package api

import (
	"Go-000/Week02/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Run() {
	r := gin.New()
	r.GET("/query/:ID", QueryUser)
	r.Run()
}

func QueryUser(c *gin.Context) {
	ID := 8
	user, _, code := service.QueryUser(uint64(ID))

	if code != 200 {
		c.String(http.StatusOK, "default string")
		return
	}
	c.JSON(code, gin.H{
		"userName": user.Name,
	})
}
