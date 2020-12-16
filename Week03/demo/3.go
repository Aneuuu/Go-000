package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"time"
)

func main() {
	g, ctx := errgroup.WithContext(context.Background())
	g.Go(func() error {
		return errors.New("test")
	})
	
	g.Go(func() error {
		time.Sleep(5000 * time.Millisecond)
		return errors.New("test2")
	})
	
	err := g.Wait()
	fmt.Println(err)
	fmt.Println(ctx.Err())
	log.Fatal()
}
