package main

import (
	"fmt"
	"net/http"
	"time"

	_ "sysinfo"

	Mux "mux"
)

func main() {

	fmt.Println("hello")

	fmt.Println("hello two")
	go http.ListenAndServe(":9527", Mux.GetMuxSys())
	for true {
		time.Sleep(time.Second * 10)
	}

}
