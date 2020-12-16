//+build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"context"
	"github.com/google/wire"
)

func InitializeApp(ctx context.Context) App {
	wire.Build(NewApp, NewApi, NewBiz, NewDBEngine)
	return App{}
}