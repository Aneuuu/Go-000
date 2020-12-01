package main

import (
	"Go-000/Week02/api"
	"fmt"
	xerrors "github.com/pkg/errors"
	"os"
)

func main() {
	err := api.Run()
	if err != nil {
		fmt.Printf("original error: %v \n", xerrors.Cause(err))
		fmt.Printf("stack error: \n+%v\n", err)
	}
	os.Exit(129)
}
