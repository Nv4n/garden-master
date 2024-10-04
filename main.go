package main

import (
	"fmt"
	_ "github.com/Masterminds/sprig/v3"
	_ "github.com/gorilla/handlers"
	_ "github.com/gorilla/mux"
	_ "github.com/gorilla/sessions"
	_ "github.com/lib/pq"
)

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:min(len(primes), 10000)]
	fmt.Println(s)
}
