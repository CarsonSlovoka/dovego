package urls

import (
    _ "embed"
    "github.com/CarsonSlovoka/dovego/app/server"
    "net/http"
)

//go:embed static/img/favicon.svg
var faviconBytes []byte

func initFavicon() {
    server.Mux.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Add("Content-Type", "image/svg+xml")
        _, _ = w.Write(faviconBytes)
    })
}
