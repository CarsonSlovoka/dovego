package main

import (
	"flag"
	"fmt"
	"github.com/CarsonSlovoka/dovego/app"
	"github.com/CarsonSlovoka/dovego/app/config"
	log2 "github.com/CarsonSlovoka/dovego/app/log"
	"github.com/CarsonSlovoka/dovego/app/server"
	"github.com/CarsonSlovoka/dovego/app/urls"
	log "log"
	"os"
	"os/exec"
)

func init() {
	var workDir string
	flag.StringVar(&workDir, "wDir", ".", "working directory")
	flag.Parse()
	if err := os.Chdir(workDir); err != nil {
		log.Fatal(err)
	}
	workingDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(fmt.Sprintf("Working Directory:%s", workingDir))
}

func main() {
	file := log2.InitLog("dovego.temp.log")
	if file != nil {
		defer func() {
			log2.Trace.Printf("Exit App.")
			_ = file.Close()
		}()
	}
	config.LoadConfig("manifest.dovego.json", &app.Config)

	quit := make(chan bool)
	log2.Trace.Printf("%+v\n", app.Config)
	port := app.Config.Server.Port
	go func() {
		urls.InitURLs()
		if err := server.ListenAndServe(port); err != nil {
			log2.Trace.Println(err)
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
