package main
//
//import (
//	"context"
//	"fmt"
//	"net/http"
//)
//
//type app struct {
//
//}
//
//type debug struct {
//
//}
//
//func (a *app) ServeHTTP (w http.ResponseWriter, r *http.Request){
//	fmt.Fprintln(w, "hello, app")
//}
//
//
//func (d *debug) ServeHTTP (w http.ResponseWriter, r *http.Request){
//	fmt.Fprintln(w, "hello, debug")
//}
//
//func serverApp(addr string, handler http.Handler, stop <- chan struct{}) error {
//	s := http.Server{
//		Addr: addr,
//		Handler: handler,
//	}
//	go func() {
//		<- stop
//		s.Shutdown(context.Background())
//	}()
//
//	return s.ListenAndServe()
//}
//
//func serverDubug(addr string, handler http.Handler, stop <- chan struct{}) error {
//	s := http.Server{
//		Addr: addr,
//		Handler: handler,
//	}
//	go func() {
//		<- stop
//		s.Shutdown(context.Background())
//	}()
//
//	return s.ListenAndServe()
//}
//
//
//
//func main() {
//	done := make(chan error, 2)
//	stop := make(chan struct{})
//	a := new(app)
//	d := new(debug)
//
//	func() {
//		done <- serverApp("127.0.0.1:8080", a, stop)
//	}()
//
//	go func() {
//		done <- serverDubug("127.0.0.1:8081", d, stop)
//	}()
//	var stopped bool
//	for i := 0; i < cap(done); i++ {
//		if err := <- done; err != nil{
//			fmt.Println("error: %v", err)
//		}
//		if ! stopped {
//			stopped = true
//			close(stop)
//		}
//	}
//}
