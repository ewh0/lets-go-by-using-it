package ch13

import (
	"log"
	"net/http"
	"time"
)

func Serve() {
	mux := http.NewServeMux()
	s := http.Server{
		Addr:         "0.0.0.0:8080",
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 90 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      mux,
	}

	mux.HandleFunc("GET /system/time/current", func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		s := t.Format("2006-01-02T15:04:05-07:00")
		w.Write([]byte(s))
	})

	err := s.ListenAndServe()
	if err != nil {
		if err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}
}
