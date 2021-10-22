package main

import (
    "fmt"
    "github.com/CarsonSlovoka/dovego/app"
    "github.com/CarsonSlovoka/dovego/app/config"
    "github.com/CarsonSlovoka/dovego/app/log"
    "github.com/CarsonSlovoka/dovego/app/server"
    "github.com/CarsonSlovoka/dovego/app/urls"
    "os/exec"
)

func main() {
    file := log.InitLog("dovego.temp.log")
    if file != nil {
        defer func() {
            log.Trace.Printf("Exit App.")
            _ = file.Close()
        }()
    }
    config.LoadConfig("manifest.dovego.json", &app.Config)

    quit := make(chan bool)
    log.Trace.Printf("%+v\n", app.Config)
    port := app.Config.Server.Port
    go func() {
        urls.InitURLs()
        if err := server.ListenAndServe(port); err != nil {
            log.Trace.Println(err)
        }
        quit <- true
    }()

    rootURL := fmt.Sprintf("http://localhost:%d", port)
    go func() {
        if err := exec.Command("rundll32", "url.dll,FileProtocolHandler",
            rootURL,
        ).Start(); err != nil {
            panic(err)
        }
    }()

    for {
        select {
        case <-quit:
            // log.Println("Close App.")
            return
        }
    }
}
