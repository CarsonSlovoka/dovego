package http

import (
	"net/http"
)

// 其實我們只要給出Handler即可，之所以參數要放FileSystem(一個interface)，是因為這樣你就可以決定是要一般檔案或者用embed的檔案也可以
// Handler要實現ServeHTTP，所以這裡我們創建了一個新的type fileHandler來表示
// fileHandler我們最主要要運用的是他的Open和Read的方法，而他的對象主要來自r.URL.Path，把URL的路徑當作檔案路徑
func FileServer(fs http.FileSystem) http.Handler {
	return &fileHandler{fs}
}
