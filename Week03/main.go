package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
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

func (s *server) start(ctx context.Context) error {
	go func() {
		<-ctx.Done()
		_ = s.Shutdown(ctx)
	}()
	return s.ListenAndServe()
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
	
	server1 := newServer("127.0.0.1:8008", a)
	server2 := newServer("127.0.0.1:8009", b)
	
	g, ctx := errgroup.WithContext(context.Background())
	
	g.Go(func() error{
		time.Sleep(5*time.Second)
		return errors.New("manual exit")
	})
	
	g.Go(func() error {
		fmt.Println("start server1")
		return server1.start(ctx)
	})
	
	g.Go(func() error {
		fmt.Println("start server2")
		return server2.start(ctx)
	})
	
	g.Go(func() error{
		for {
			select {
			case <-ctx.Done():
				fmt.Println("haha")
				return ctx.Err()
			default:
				return errors.New("other quit")
			}
		}
	})
	
	if err := g.Wait(); err != nil {
		fmt.Println("wait error", err)
	}
}
