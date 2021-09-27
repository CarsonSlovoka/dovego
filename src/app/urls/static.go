package urls

import (
    "embed"
    _ "embed"
    "github.com/CarsonSlovoka/dovego/app/log"
    "github.com/CarsonSlovoka/dovego/app/server"
    http2 "github.com/CarsonSlovoka/dovego/pkg/net/http"

    "net/http"
)

//go:embed static/* tmpl/*
var systemResourceFS embed.FS

//go:embed tmpl/index.html
var rootHTMLFS embed.FS

func initSystemStaticResource() {
    server.Mux.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if r.URL.Path == "/" {
            bytes, err := rootHTMLFS.ReadFile("tmpl/index.html")
            if err != nil {
                log.Error.Printf(err.Error())
            }
            _, _ = w.Write(bytes)
            return
        }

        http2.FileServer(http.FS(systemResourceFS)).ServeHTTP(w, r)
    })
}
