package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
	"test/lissajous"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/counter", counter)
	http.HandleFunc("/li", sin)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func sin(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	ca := r.Form.Get("cycles")
	cycles, _ := strconv.Atoi(ca)
	fmt.Println(cycles)
	if cycles == 0 {
		fmt.Println("not set")
		cycles = 5
	}
	fmt.Println("after", cycles)
	lissajous.Lissajous(w, float64(cycles))
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
