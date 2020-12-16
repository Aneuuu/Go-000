package biz

import (
	"context"
	"github.com/pkg/errors"
	"strconv"
)

func (b *Biz) GetUserNameRequest(ctx context.Context, id string) (string, error) {
	ID, err := strconv.Atoi(id)
	if err != nil {
		return "", errors.WithMessage(err, "string to id error")
	}
	b.Dao.Users.ID = uint64(ID)
	if err := b.Dao.GetUserName(ctx); err != nil {
		return "", err
	}
	return b.Dao.Users.Name, nil
}