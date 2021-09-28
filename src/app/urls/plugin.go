package urls

import (
    "fmt"
    "github.com/CarsonSlovoka/dovego/app"
    "github.com/CarsonSlovoka/dovego/app/config"
    "github.com/CarsonSlovoka/dovego/app/server"
    http2 "github.com/CarsonSlovoka/dovego/pkg/net/http"
    "net/http"
    "strings"
)

func initPlugin() {
    http2.DebugMode = app.Config.Debug.Enable

    server.Mux.PathPrefix("/plugin/").Handler(
        // http.StripPrefix("/plugin", someHandler),
        http2.FileServer(http.Dir("")), // http.Dir its type is a string, but it implements the "Open" method too.
    )

    pluginList := app.Config.Plugins
    for _, plugin := range pluginList {
        setPluginURL(plugin) // for迴圈的變數會一直沿用，所以如果把這個實現直接寫到這邊，這只會抓到最後一個plugin而已，因此要使用函數來copy參數過去
    }
}

func setPluginURL(plugin config.Plugin) {
    server.Mux.PathPrefix(fmt.Sprintf("/%s/", plugin.Name)).
        HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            r.URL.Path = plugin.Path + strings.TrimPrefix(r.URL.Path, fmt.Sprintf("/%s/", plugin.Name))
            http2.FileServer(http.Dir("")).ServeHTTP(w, r)
        })
}
