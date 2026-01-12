package singleflight

import (
	"fmt"
	"golang.org/x/sync/singleflight"
	"net/http"
	"time"
)

func callExternalApi() (int, error) {
	fmt.Println("Пример долгого обращения к внешнему сервису")
	time.Sleep(time.Second * 5)
	return 200, nil
}

func withoutSingleFlight() {
	http.HandleFunc("/request", func(w http.ResponseWriter, r *http.Request) {
		status, err := callExternalApi()
		if err != nil {
			fmt.Fprintf(w, "Error: %d", status)
		}
		fmt.Fprintf(w, "Status: %d", status)
	})
}

func withSingleflight() {
	var requestGroup singleflight.Group
	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		// Все запросы с ключом "api_status" объединяются
		v, err, shared := requestGroup.Do("api_status", func() (interface{}, error) {
			return callExternalApi()
		})
		if err != nil {
			fmt.Fprintf(w, "Error: %d", v)
		}
		fmt.Fprintf(w, "Status: %s (shared=%v)", v.(string), shared)
		// shared = true для всех, кроме самого первого вызова
	})
}
