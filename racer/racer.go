package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	blackhole := make([]byte, 100)
	wg := new(sync.WaitGroup)
	for i := 0; i < 10; i++ {
		go func(i int) {
			log.Printf("%d starting", i)
			for ii := range blackhole {
				blackhole[ii] = byte(i)
				time.Sleep(time.Nanosecond)
			}
			log.Printf("%d done", i)
		}(i)
	}

	log.Printf("> waiting")
	wg.Wait()
	log.Printf("> done")
}
