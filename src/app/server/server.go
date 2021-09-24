package server

import (
    "fmt"
    "github.com/gorilla/mux"
    "net/http"
)

var (
    Mux    *mux.Router // http.HandleFunc
    Server http.Server
)

func init() {
    Mux = mux.NewRouter()
}

func ListenAndServe(portNumber int) error {
    Server = http.Server{Addr: fmt.Sprintf(":%d", portNumber), Handler: Mux}
    err := Server.ListenAndServe()
    return err
}
