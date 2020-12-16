package dao

import (
	"context"
	"database/sql"
)

type DBModel struct {
	DBEngine    *sql.DB
	DBInfo      *DBInfo
	Users       *Users
	Ctx			context.Context
}

type DBInfo struct {
	DBType      string
	Host        string
	Database    string
	UserName    string `json:"username"`
	Password    string
	Charset     string
}

type Users struct {
	ID      uint64
	Name    string
}
