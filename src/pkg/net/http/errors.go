package http

import (
    "errors"
    "fmt"
    "io/fs"
    "mime"
    "net/http"
)

func Error(w http.ResponseWriter, error string, code int, contentType string) {
    if contentType == "" {
        contentType = mime.TypeByExtension(".txt")
    }
    w.Header().Set("Content-Type", contentType)
    w.Header().Set("X-Content-Type-Options", "nosniff") // Avoid XSS attacks.
    w.WriteHeader(code)
    _, _ = fmt.Fprintln(w, error)
}

func ErrorWithHTML(w http.ResponseWriter, error string, code int) {
    error = fmt.Sprintf(`<h1 style="text-align: center">%d - %s</h1>`, code, error)
    Error(w, error, code, mime.TypeByExtension(".html"))
}

func ToHTTPError(err error) (msg string, httpStatus int) {
    if errors.Is(err, fs.ErrNotExist) {
        return http.StatusText(http.StatusNotFound), http.StatusNotFound
    }
    if errors.Is(err, fs.ErrPermission) {
        return http.StatusText(http.StatusForbidden), http.StatusForbidden
    }
    return http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError
}
