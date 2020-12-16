package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type server struct{
	*http.Server
}

type a struct {}

type b struct{}


func (a *a) ServeHTTP(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "a")
}


func (b *b) ServeHTTP(w http.ResponseWriter, r *http.Request){
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
		Addr: addr,
		Handler: handler,
	}}
}


func main() {
	a := new(a)
	b := new(b)
	
	
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGUSR1, syscall.SIGUSR2)
	
	server1 := newServer("127.0.0.1:8080", a)
	server2 := newServer("127.0.0.1:8081", b)
	
	g, _ := errgroup.WithContext(context.Background())
	
	g.Go(func() error {
		return server1.start()
	})
	
	g.Go(func() error {
		return server2.start()
	})
	
	g.Go(func() error {
		for {
			select {
			case sig := <-sigs:
				fmt.Println()
				fmt.Println(sig)
				server1.stop()
				server2.stop()
			}
		}
	})
	
	if err := g.Wait(); err != nil{
		fmt.Println(err)
	}
	//fmt.Println(ctx.Err())
}
