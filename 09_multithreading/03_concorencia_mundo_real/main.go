package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
	"time"
)

var number uint64 = 0

func main() {
	// m := sync.Mutex{}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// m.Lock()
		// number++
		// m.Unlock()
		//soma atomica: ele altera o valor de number e garante que não haverá concorrência
		atomic.AddUint64(&number, 1)
		time.Sleep(300 * time.Millisecond)
		w.Write([]byte("Você teve tantas requisições: " + fmt.Sprint(number)))
	})

	//para cada requisição o servidor cria uma thread
	http.ListenAndServe(":3000", nil)
}
