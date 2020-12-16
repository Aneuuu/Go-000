package main

import (
	"Go-000/Week04/api"
	"Go-000/Week04/biz"
	"Go-000/Week04/dao"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"os"
	"os/signal"
	"syscall"
)

type App struct {
	api *api.API
}

func NewDBEngine(ctx context.Context) *dao.DBModel{
	return &dao.DBModel{DBInfo: &dao.DBInfo{
		DBType:   "mysql",
		Host:     "10.1.2.230",
		Database: "demorm",
		UserName: "root",
		Password: "ys@qq.com",
		Charset:  "utf8",
	}, Users: &dao.Users{}, Ctx: ctx}
}

func NewBiz(d *dao.DBModel) *biz.Biz{
	return &biz.Biz{Ctx: d.Ctx, Dao: d}
}

func NewApi(b *biz.Biz) *api.API {
	return &api.API{Engine: gin.Default(), Biz: b}
}

func NewApp(a *api.API) App {
	return App{api: a}
}

func (app App) Start() {
	if err := app.api.Run(":8080"); err != nil{
		panic(err)
	}
}

func main() {
	g, ctx := errgroup.WithContext(context.Background())
	sigs := make(chan os.Signal)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	g.Go(func() error{
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
				return errors.New("other quit")
			}
		}
	})

	app := InitializeApp(ctx)
	app.Start()

	if err := g.Wait(); err != nil {
		fmt.Println(err)
	}
}