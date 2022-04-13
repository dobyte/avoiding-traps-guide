/**
 * @Author: fuxiao
 * @Email: 576101059@qq.com
 * @Date: 2022/4/13 11:29 上午
 * @Desc: TODO
 */

package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"runtime/pprof"
	"time"
)

func callWithTimeout(f func() error, timeout time.Duration) (err error) {
	ch := make(chan error, 1)
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	go func() {
		ch <- f()
	}()

	select {
	case err = <-ch:
	case <-ctx.Done():
	}

	return
}

func servePProf() {
	addr := ":8080"

	fmt.Printf("pporf server started, listening on %s\n", addr)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		p := pprof.Lookup("goroutine")
		_ = p.WriteTo(w, 1)
	})
	_ = http.ListenAndServe(addr, nil)
}

func main() {
	go servePProf()

	_ = callWithTimeout(func() error {
		time.Sleep(5 * time.Second)
		return errors.New("remote call failed")
	}, time.Second)

	select {}
}
