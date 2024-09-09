package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
	"time"
)

var numeroChamadas uint64 = 0

func main() {
	// m := sync.Mutex{}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// m.Lock()
		// numeroChamadas++
		// m.Unlock()

		atomic.AddUint64(&numeroChamadas, 1)
		time.Sleep(300 * time.Millisecond)
		w.Write([]byte(fmt.Sprintf("Você é o visitante número %d", numeroChamadas)))
	})
	http.ListenAndServe(":3000", nil)
}
