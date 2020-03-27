package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	// Have you ever wondered what error is received by a Writer when they
	// write to a connection closed by their peer? Here's the answer + a
	// repro because it's OS dependent.
	errorPipe = &net.OpError{
		Op:     "write",
		Net:    "tcp",
		Source: &net.TCPAddr{},
		Addr:   &net.TCPAddr{},
		Err: &os.SyscallError{
			Syscall: "write",
			Err:     syscall.EPIPE,
		},
	}
)

func main() {
	l, err := net.Listen("tcp", "127.0.0.1:")
	if err != nil {
		log.Fatalf("listen err: %v", err)
	}

	done := make(chan struct{})
	go func() {
		defer close(done)
		for {
			conn, err := l.Accept()
			if err != nil {
				log.Printf("Accept err: %v", err)
				return
			}

			go handle(conn)
		}
	}()

	log.Printf("Listener addr: %s", l.Addr())

	conn, err := net.Dial("tcp", l.Addr().String())
	if err != nil {
		log.Fatalf("Dial err: %v", err)
	}

	log.Printf("Close err: %s", conn.Close())

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)
	select {
	case <-done:
		log.Printf("Listener done")
	case s := <-sigCh:
		log.Printf("Signal: %v", s)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	for {
		n, err := conn.Write([]byte{'x'})
		if err != nil {
			log.Printf("n=%d err(T)=%T err=%v", n, err, err)
			if operr, ok := err.(*net.OpError); ok {
				log.Printf("err#=%#v", operr)
				log.Printf("err(T)=%T err=%v", operr.Err, operr.Err)
				log.Printf("temp=%t timeout=%t", operr.Temporary(), operr.Timeout())
				if syserr, ok := operr.Err.(*os.SyscallError); ok {
					log.Printf("err#=%#v", syserr)
				}
			}
			return
		}

		time.Sleep(time.Second)
	}
}
