package urls

import (
	"embed"
	_ "embed"
	"github.com/CarsonSlovoka/dovego/app"
	"github.com/CarsonSlovoka/dovego/app/server"
	http2 "github.com/CarsonSlovoka/dovego/pkg/net/http"
	"github.com/CarsonSlovoka/dovego/pkg/utils"
	"os"
	"path/filepath"
	"strings"

	"net/http"
)

//go:embed static/* tmpl/*
var systemResourceFS embed.FS

//go:embed tmpl/index.html
var rootHTMLFS embed.FS

func initSystemStaticResource() {
	server.Mux.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		startURL := app.Config.Server.StartURL

		urlPath := r.URL.Path
		if urlPath == "/" && startURL != "" {
			// log2.Trace.Println(filepath.Abs("./" + startURL))
			if _, err := os.Stat(startURL); os.IsNotExist(err) || strings.ToLower(filepath.Ext(startURL)) != ".html" {
				http2.ErrorWithHTML(w, err.Error(), http.StatusNotFound)
				return
			}

			http.ServeFile(w, r, app.Config.Server.StartURL)

			/*
				bytes, err := rootHTMLFS.ReadFile(app.Config.Server.StartURL)
				if err != nil {
					http2.ErrorWithHTML(w, err.Error(), http.StatusForbidden)
					return
				}
				_, _ = w.Write(bytes)
			*/
			return
		}

		// PWA Support: <link rel="manifest" href="/.webmanifest">
		if utils.ContainsAny([]string{"/manifest.json", "/.webmanifest", "/sw.js"}, urlPath) {
			http.ServeFile(w, r, "."+urlPath)
			return
		}

		http2.FileServer(http.FS(systemResourceFS)).ServeHTTP(w, r)
	})
}
