package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"strconv"
	"sync"
	"time"
)

func main() {
	s := http.Server{ConnState: HTTPConnStateFunc()}
	l, err := net.Listen("tcp", "0.0.0.0:0")
	if err != nil {
		log.Fatalf("error listeneing: %v", err)
	}
	go func() {
		if err := s.Serve(l); err != http.ErrServerClosed {
			log.Printf("error serving: %v", err)
		}
		log.Printf("http server closed")
	}()

	shutdown := func() {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		if err := s.Shutdown(ctx); err != nil && err != http.ErrServerClosed {
			log.Printf("error shutting down server: %v", err)
		}
		log.Printf("server shutdown")
	}
	defer shutdown()

	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			addr := "http://" + l.Addr().String() + "/" + strconv.Itoa(i)
			log.Printf("%d get  %s", i, addr)
			resp, err := http.Get(addr)
			if err != nil {
				log.Printf("%d error getting %s: %v", i, addr, err)
				return
			}

			log.Printf("%d dump %s", i, addr)
			buf, err := httputil.DumpResponse(resp, true)
			if err != nil {
				log.Printf("%d error reading response: %v", i, err)
				return
			}

			log.Printf("%d response:\n%s", i, string(buf))
		}(i)
	}

	log.Printf("waiting on gets")
	wg.Wait()
	log.Printf("gets done")
}

func HTTPConnStateFunc() func(net.Conn, http.ConnState) {
	message := "Your IP is issuing too many concurrent connections, please rate limit your calls\n"
	tooManyRequestsResponse := []byte(fmt.Sprintf("HTTP/1.1 429 Too Many Requests\r\n"+
		"Content-Type: text/plain\r\n"+
		"Content-Length: %d\r\n"+
		"Connection: close\r\n\r\n%s", len(message), message))
	// Buffer shared by all clients, we don't care about result
	//readBuffer := make([]byte, 0, 4096)      // lots of http client errors !!
	readBuffer := make([]byte, 4096) // no errors but seems like a data race !!
	return func(conn net.Conn, state http.ConnState) {
		switch state {
		case http.StateNew:
			log.Printf("state new")
			// We read to be sure not to block the client
			_, err := conn.Read(readBuffer)
			if err == nil {
				conn.Write(tooManyRequestsResponse)
			}
			conn.Close()
		case http.StateHijacked:
			log.Printf("state hijacked")
		case http.StateClosed:
			log.Printf("state closed")
		}
	}
}
