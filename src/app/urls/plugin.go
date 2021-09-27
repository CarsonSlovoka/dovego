package urls

import (
    "github.com/CarsonSlovoka/dovego/app"
    "github.com/CarsonSlovoka/dovego/app/server"
    http2 "github.com/CarsonSlovoka/dovego/pkg/net/http"
    "net/http"
)

func initPlugin() {
    http2.DebugMode = app.Config.Debug.Enable
    server.Mux.PathPrefix("/plugin/").Handler(
        // http.StripPrefix("/plugin", someHandler),
        http2.FileServer(http.Dir("")), // http.Dir its type is a string, but it implements the "Open" method too.
    )
}
