package urls

import (
    "context"
    "github.com/CarsonSlovoka/dovego/app"
    "github.com/CarsonSlovoka/dovego/app/log"
    "github.com/CarsonSlovoka/dovego/app/server"
    "net/http"
)

func initSystemURL() {
    serverMux := server.Mux
    s := &server.Server
    if shutdownURL := app.Config.Server.Shutdown.URL; shutdownURL != "" {
        serverMux.HandleFunc(shutdownURL, func(w http.ResponseWriter, r *http.Request) {
            if err := s.Shutdown(context.Background()); err != nil {
                log.Trace.Printf("Can't close server: %v", err)
            }
        })
    }
}
