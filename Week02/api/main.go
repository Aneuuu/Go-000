package api

import (
	"Go-000/Week02/service"
	"github.com/gin-gonic/gin"
	xerrors "github.com/pkg/errors"
	"net/http"
)


func Run() error {
	r := gin.Default()
	r.GET("/query/:ID", QueryUser)
	err := r.Run()
	if err != nil {
		xerrors.WithMessage(err, "Gin Server Runnig Failed!")
	}
	return nil
}

func QueryUser(c *gin.Context){
	ID := 8
	user, _, code := service.QueryUser(uint64(ID))

	if code != 200 {
		c.Redirect(http.StatusOK, QueryDefault())
	}
	c.JSON(code, gin.H{
		"userName": user.Name,
	})
}

func QueryDefault() string{
	return "default string"
}