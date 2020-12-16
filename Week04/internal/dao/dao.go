package dao

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"time"
)

func NewDBModel(info *DBInfo) *DBModel{
	return &DBModel{DBInfo: info, Users: &Users{}}
}

func (m *DBModel) Connect() error {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=True&loc=Local",
		m.DBInfo.UserName,
		m.DBInfo.Password,
		m.DBInfo.Host,
		m.DBInfo.Database,
		m.DBInfo.Charset,
	)
	m.DBEngine, err = sql.Open(m.DBInfo.DBType, dsn)
	if err != nil {
		return err
	}
	
	return nil
}


func (m *DBModel) GetUserName(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	err := m.DBEngine.QueryRowContext(ctx, "SELECT name FROM users WHERE id=?", m.Users.ID).Scan(&m.Users.Name)
	switch {
	case err == sql.ErrNoRows:
		return errors.Wrap(err, fmt.Sprintf("no user with id %d\n", m.Users.ID))
	case err != nil:
		return errors.Wrap(err, "query error")
	default:
		return nil
	}
}