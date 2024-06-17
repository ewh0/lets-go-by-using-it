package ch13

import (
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"
)

func ServeWithLogging() {
	mux := http.NewServeMux()

	// logging
	options := &slog.HandlerOptions{Level: slog.LevelDebug}
	f, err := os.Create("access.log")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	handler := slog.NewJSONHandler(f, options)
	logger := slog.New(handler)

	// mux
	mux.HandleFunc("GET /system/time/current", func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		s := t.Format("2006-01-02T15:04:05-07:00")
		w.Write([]byte(s))
	})
	// logging
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Debug("New client request received", "source", r.RemoteAddr)
		mux.ServeHTTP(w, r)
	})

	s := http.Server{
		Addr:         "0.0.0.0:8080",
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 90 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      h,
	}
	err = s.ListenAndServe()
	if err != nil {
		if err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}
}
