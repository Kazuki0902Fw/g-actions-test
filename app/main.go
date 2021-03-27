package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-2) + fibonacci(n-1)
}

func fibonacci2(n int, memo map[int]int) int {
	if n < 2 {
		return n
	}
	if _, ok := memo[n]; !ok {
		memo[n] = fibonacci2(n-2, memo) + fibonacci2(n-1, memo)
	}
	return memo[n]
}

func handler(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano())
	fibN := rand.Intn(11)
	fmt.Fprintf(w, "hello fibonacci(%d): %d", fibN, fibonacci(fibN))
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
