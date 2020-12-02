package service

import (
	"Go-000/Week02/dao"
	"fmt"
	xerrors "github.com/pkg/errors"
)

func QueryUser(ID uint64) (*dao.Users, error, int) {
	user, err, code := dao.QueryUser(ID)
	if err != nil {
		fmt.Printf("original error: %v \n", xerrors.Cause(err))
		fmt.Printf("stack error: \n %v\n", err)
	}

	return user, nil, code
}
