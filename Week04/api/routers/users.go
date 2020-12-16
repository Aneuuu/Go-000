package routers

import (
	"Go-000/Week04/api"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Users struct {
	api *api.API
}

func (u *Users) AddUsersRouter() {
	u.api.AddRouter("/users")
	g := u.api.RouterGroup
	
	g.GET("/ping", func(c *gin.Context){
		c.JSON(http.StatusOK, "ping")
	})
	
	g.GET("/user/:id", func(c *gin.Context){
		id := c.Param("id")
		name, err := u.api.Biz.GetUserNameRequest(c, id)
		if err != nil {
			c.JSON(http.StatusBadRequest, "err")
		}
		c.JSON(http.StatusOK, name)
	})
}