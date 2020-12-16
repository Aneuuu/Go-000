package api

import (
	"Go-000/Week04/biz"
	"github.com/gin-gonic/gin"
	"net/http"
)

type API struct {
	Engine      *gin.Engine
	RouterGroup *gin.RouterGroup
	Biz         *biz.Biz
}

func (a *API) AddEngine() {
	a.Engine = gin.Default()
}

func (a *API) AddRouter(routerPath string) {
	a.RouterGroup = a.Engine.Group(routerPath)
}

func (a *API) AddGetMethod(path string, f func(c *gin.Context)) {
	a.RouterGroup.GET(path, f)
}

func (a *API) AddUsersRouter() {
	a.AddRouter("/users")
	g := a.RouterGroup

	g.GET("/ping", func(c *gin.Context){
		c.JSON(http.StatusOK, "ping")
	})

	g.GET("/user/:id", func(c *gin.Context){
		id := c.Param("id")
		name, err := a.Biz.GetUserNameRequest(c, id)
		if err != nil {
			c.JSON(http.StatusBadRequest, "err")
		}
		c.JSON(http.StatusOK, name)
	})
}

func (a *API) Run(addr ...string) error {
	if err := a.Engine.Run(addr...); err != nil {
		return err
	}
	return nil
}