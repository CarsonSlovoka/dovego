package http

import (
	"io"
	"io/ioutil"
	"log"
	"mime"
	"net/http"
	"path"
	"path/filepath"
	"strings"
	"time"
)

type fileHandler struct {
	fs http.FileSystem
}

func (fh *fileHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	urlPath := r.URL.Path
	if !strings.HasPrefix(urlPath, "/") {
		// r.URL.Path always begins with /
		urlPath = "/" + urlPath
		r.URL.Path = urlPath
	}
	serveFile(w, r, fh.fs, path.Clean(urlPath), true)
}

func serveFile(w http.ResponseWriter, r *http.Request, fs http.FileSystem, filepath string, redirect bool) {
	const indexPage = "/index.html"

	if strings.HasSuffix(r.URL.Path, indexPage) {
		localRedirect(w, r, "./") // 導回相對路徑的目錄
		return
	}

	file, err := fs.Open(filepath)
	if err != nil {
		msg, statusCode := ToHTTPError(err)
		ErrorWithHTML(w, msg, statusCode)
		return
	}
	defer func() {
		_ = file.Close()
	}()

	fileState, err := file.Stat()
	if err != nil {
		msg, statusCode := ToHTTPError(err)
		ErrorWithHTML(w, msg, statusCode)
		return
	}

	if redirect {
		url := r.URL.Path
		if fileState.IsDir() {
			if url[len(url)-1] != '/' {
				localRedirect(w, r, path.Base(url)+"/")
				return
			}
		} else if url[len(url)-1] == '/' {
			localRedirect(w, r, "../"+path.Base(url))
			return
		}
	}

	if fileState.IsDir() {
		url := r.URL.Path
		if url == "" || url[len(url)-1] != '/' {
			localRedirect(w, r, path.Base(url)+"/")
			return
		}

		// use contents of index.html for directory, if present
		indexPath := strings.TrimSuffix(filepath, "/") + indexPage
		indexFile, err := fs.Open(indexPath)
		if err == nil {
			defer func() {
				_ = indexFile.Close()
			}()

			if indexFileState, err := indexFile.Stat(); err == nil {
				if indexFileState.IsDir() { // It is possible that "index.html" is a folder
					return
				}
				filepath = indexPath
				fileState = indexFileState
				file = indexFile
			}
		}
	}

	if DebugMode {
		log.Println(filepath)
	}
	sizeFunc := func() (int64, error) { return fileState.Size(), nil }
	serveContent(w, r, filepath, fileState.ModTime(), sizeFunc, file)
}

// localRedirect gives a Moved Permanently response.
// It does not convert relative paths to absolute paths like Redirect does.
func localRedirect(w http.ResponseWriter, r *http.Request, newPath string) {
	if q := r.URL.RawQuery; q != "" {
		newPath += "?" + q
	}
	w.Header().Set("Location", newPath)
	w.WriteHeader(http.StatusMovedPermanently)
}

func serveContent(w http.ResponseWriter, r *http.Request, name string, modtime time.Time, sizeFunc func() (int64, error), contentReadSeeker io.ReadSeeker) {
	cTypes, haveType := w.Header()["Content-Type"]
	// ↓ Decide Content-Type
	var cType string
	if !haveType {
		cType = mime.TypeByExtension(filepath.Ext(name))
		if cType == "" {
			// 這部分應該是要去讀取檔案，去sniff(嗅探)到底此檔案可能的ctype為何
			cType = mime.TypeByExtension(filepath.Ext(".txt")) // treat it as the Text file.
		}
		w.Header().Set("Content-Type", cType)
	} else if len(cTypes) > 0 {
		cType = cTypes[0]
	}

	// ↓ Handle Content-Range header (There may be a lot of data so that that system will pass it in batches. I do not consider this situation, directly read all the data.)
	_, err := sizeFunc()
	if err != nil {
		ErrorWithHTML(w, err.Error(), http.StatusInternalServerError)
	}

	byteData, err := ioutil.ReadAll(contentReadSeeker)

	if err != nil {
		ErrorWithHTML(w, "Read Error"+err.Error(), http.StatusForbidden)
		return
	}
	_, _ = w.Write(byteData)
	return
}
