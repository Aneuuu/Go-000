package service

import "Go-000/Week02/dao"

func QueryUser(ID uint64) (*dao.User, error, int) {
	return dao.QueryUser(ID)
}