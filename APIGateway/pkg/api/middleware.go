package api

import (
	"context"
	"log"
	"net/http"
	"time"
)

const requestIdHeader = "request_id"

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewLoggingResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
	return &loggingResponseWriter{w, http.StatusOK}
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

// присвоение номера запросу
func requestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestId := r.Header.Get(requestIdHeader)
		if requestId == "" {
			log.Println("empty request_id")
			requestId = "empty request_id"
			// requestId = fmt.Sprint(rand.Intn(1_000_000))
		}
		ctx := context.WithValue(r.Context(), requestIdHeader, requestId)
		req := r.WithContext(ctx)

		// b, _ := httputil.DumpRequest(req, true)
		// fmt.Printf("middlaware req: %+v\n", string(b))
		next.ServeHTTP(w, req)
	})
}

// логирование запросов
func logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		requestId := r.Context().Value(requestIdHeader)
		// requestId := r.Header.Get(requestIdHeader)

		lrw := NewLoggingResponseWriter(w)
		next.ServeHTTP(lrw, r)
		status := lrw.statusCode
		log.Println(t.Local(), r.Method, status, http.StatusText(status), r.RemoteAddr, r.RequestURI, requestId)
	})
}
