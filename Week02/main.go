package main

import "io"

type errWriter struct {
	io.Writer
	err error
}

func (e *errWriter) Write(buf []byte) (int, error) {
	if e.err != nil {
		return 0, nil
	}
	var n int
	n, e.err = e.Writer.Write(buf)
	return n, nil
}

func main() {

}
