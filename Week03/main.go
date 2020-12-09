package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type server struct {
	*http.Server
}

type a struct{}

type b struct{}

func (a *a) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "a")
}

func (b *b) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "b")
}

func (s *server) start() error {
	return s.ListenAndServe()
}
func (s *server) stop() error {
	return s.Shutdown(context.Background())
}

func newServer(addr string, handler http.Handler) *server {
	return &server{
		&http.Server{
			Addr:    addr,
			Handler: handler,
		}}
}

func main() {
	a := new(a)
	b := new(b)


	sigs := make(chan os.Signal)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	server1 := newServer("127.0.0.1:8080", a)
	server2 := newServer("127.0.0.1:8081", b)

	g, ctx := errgroup.WithContext(context.Background())

	g.Go(func() error {
		fmt.Println("start server1")
		return server1.start()
	})

	g.Go(func() error {
		fmt.Println("start server2")
		return server2.start()
	})

	g.Go(func() error {
		for sig := range sigs {
			switch sig {
			case syscall.SIGINT, syscall.SIGTERM:
				_ = server1.stop()
				_ = server2.stop()
				return errors.New("all stop")
			default:
				return errors.New("other quit")
			}
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		fmt.Println(err)
	}
	fmt.Println(ctx.Err())
}
