package middleware

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"gitlab.ozon.dev/kavkazov/homework-8/internal/pkg/logger"
)

func KafkaLogging(sender logger.Sender) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				err := sender.SendMessage(
					logger.LogMessage{Path: r.URL.Path, Type: r.Method, TimeStamp: time.Now()},
				)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
				next.ServeHTTP(w, r)
			},
		)
	}
}
