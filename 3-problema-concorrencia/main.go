package main

import (
	"fmt"
	"net/http"
	"time"
)

var numeroChamadas uint64 = 0

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		numeroChamadas++
		time.Sleep(300 * time.Millisecond)
		w.Write([]byte(fmt.Sprintf("Você é o visitante número %d", numeroChamadas)))
	})
	http.ListenAndServe(":3000", nil)
}
