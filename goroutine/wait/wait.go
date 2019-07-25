package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/pprof"
	"runtime"
	"sync"
	"time"
)

func x(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("X"))
	n := runtime.NumCPU()

	if n < 2 {
		log.Fatal("At least 2 CPUs needed")
	}

	// Expose CPUs to scheduler
	runtime.GOMAXPROCS(n)

	wg := &sync.WaitGroup{}

	// Use CPUs/2 to give the Go scheduler something to work with
	l := n / 2

	wg.Add(l)

	for i := 0; i < l; i++ {
		go func(wg *sync.WaitGroup, v int) {
			defer wg.Done()

			runtime.LockOSThread()

			defer runtime.UnlockOSThread()

			w.Write([]byte("Hi"))
			fmt.Printf("Hello from %d", v)

			time.Sleep(10 * time.Second)
		}(wg, i)
	}

	wg.Wait()
	log.Println("Finished!")
}

func hiHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hi"))
	c := make(chan string, 20)

	num_cores := 100
	for i := 0; i < num_cores; i++ {
		go func(c chan string) {
			w.Write([]byte(<-c))
			// fmt.Println(<-c)
		}(c)
	}

	message := "x"

	for i := 0; i < 100; i++ {
		c <- message
	}
}

func main() {
	r := http.NewServeMux()
	r.HandleFunc("/", hiHandler)
	r.HandleFunc("/x", x)

	// Register pprof handlers
	r.HandleFunc("/debug/pprof/", pprof.Index)
	r.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	r.HandleFunc("/debug/pprof/profile", pprof.Profile)
	r.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	r.HandleFunc("/debug/pprof/trace", pprof.Trace)

	http.ListenAndServe(":8080", r)
}
