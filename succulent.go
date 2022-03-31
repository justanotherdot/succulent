package main

import (
	"fmt"
  "net/http"
  "log"
)

const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

func main() {
    http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "works")
    })

    fmt.Printf("Starting server at port 8080\n")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}

func Thing(x bool) bool {
	return !x
}

func Even(x int) int {
    return 2
    //if x % 2 == 1 {
    //    if x == MinInt {
    //        return x + 1
    //    }
    //    return x - 1
    //}
    //return x
}

func MaybePanic(x int) int {
    //if x % 3 == 0 {
    //    panic("uh oh")
    //}
    return x
}
