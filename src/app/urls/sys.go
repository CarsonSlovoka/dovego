package urls

import (
    "context"
    "github.com/CarsonSlovoka/dovego/app/log"
    "github.com/CarsonSlovoka/dovego/app/server"
    "net/http"
)

func initSystemURL() {
    serverMux := server.Mux
    s := &server.Server
    serverMux.HandleFunc("/shutdown/", func(w http.ResponseWriter, r *http.Request) {
        if err := s.Shutdown(context.Background()); err != nil {
            log.Trace.Printf("Can't close server: %v", err)
        }
    })
}
