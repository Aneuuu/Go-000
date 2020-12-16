package biz

import (
	"Go-000/Week04/dao"
	"context"
)


type Biz struct {
	Ctx context.Context
	Dao *dao.DBModel
}

//func NewBiz(ctx context.Context) *Biz {
	//biz := Biz{ctx: ctx}
	//biz.Dao = dao.NewDBModel(&dao.DBInfo{
	//	DBType:   "mysql",
	//	Host:     "10.1.2.230",
	//	Database: "demorm",
	//	UserName: "root",
	//	Password: "ys@qq.com",
	//	Charset:  "utf8",
	//})
	//if err := biz.Dao.Connect();err != nil {
	//	fmt.Println(err)
	//}
	//return &biz
//}