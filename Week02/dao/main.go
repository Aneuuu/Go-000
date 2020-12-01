package dao

import (
	xerrors "github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID uint64
	Name string
}

var (
	db *gorm.DB
	err error
)

func init() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/demo?charset=utf8&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		xerrors.Wrap(err, "can't connect to database")
	}
}

func QueryUser(ID uint64) (*User, error, int) {
	var u = new(User)
	result := db.Where("id = ?", ID).Find(&u)
	if result.Error != nil {
		xerrors.Wrap(result.Error, "Query Not Found!")
	}
	return u, nil, 200
}