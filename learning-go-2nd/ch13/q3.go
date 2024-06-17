package ch13

import (
	"encoding/json"
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"
)

type timeDto struct {
	DayOfWeek  string `json:"day_of_week"`
	DayOfMonth int    `json:"day_of_month"`
	Month      string `json:"month"`
	Year       int    `json:"year"`
	Hour       int    `json:"hour"`
	Minute     int    `json:"minute"`
	Second     int    `json:"second"`
}

func ServeWithLoggingAndReturnJson() {
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

		acceptType := r.Header.Get("Accept")
		if acceptType == "application/json" {
			var dto = &timeDto{
				DayOfWeek:  t.Weekday().String(),
				DayOfMonth: t.Day(),
				Month:      t.Month().String(),
				Year:       t.Year(),
				Hour:       t.Hour(),
				Minute:     t.Minute(),
				Second:     t.Second(),
			}
			b, err := json.Marshal(dto)
			if err != nil {
				logger.Error("Failed to process client request", "error", err.Error())
				w.WriteHeader(http.StatusInternalServerError)
			} else {
				w.Header().Set("Content-type", "application/json")
				w.WriteHeader(http.StatusOK)
				w.Write(b)
			}
		} else {
			w.Write([]byte(t.Format("2006-01-02T15:04:05-07:00")))
		}
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
