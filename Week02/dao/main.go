package dao

import (
	xerrors "github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Users struct {
	ID   uint64
	Name string
}

var (
	db  *gorm.DB
	err error
)

func init() {
	dsn := "root:ys@qq.com@tcp(10.1.2.230:3306)/demorm?charset=utf8&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		xerrors.Wrap(err, "can't connect to database")
	}
}

func QueryUser(ID uint64) (*Users, error, int) {
	var u = new(Users)
	result := db.Where("id = ?", ID).Find(&u).Take(&u)
	if result.Error != nil {
		return u, xerrors.WithMessage(result.Error, "Query Not Found!"), 500
	}
	return u, nil, 200
}
