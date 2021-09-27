package http

/*
content type的處理，先比對附檔名 (可以用mime.TypeByExtension(".html")即可)
如果副檔名比對不出來，就要利用嗅探器去找可能的content type，這時候就會比較複雜，如果要做可以參考
> go/src/net/http/sniff.go
*/

/* 以下的這部分全部靠系統API即可完成 mime.TypeByExtension(".html")
type ContentTypeStruct struct {
    html string
    text string
    css  string
    xml  string

    js   string
    json string

    png   string
    image string
    svg   string
    ico   string
    bmp   string
    gif   string

    mp4  string
    webm string

    ttf   string
    otf   string
    ttc   string
    woff  string
    woff2 string
}

var ContentType *ContentTypeStruct

func init() {
    ContentType := map[string]string{"text/html; charset=utf-8"} // such that you don't add ``<meta charset="utf-8" />`` on HTML.
        "text/plain; charset=utf-8",
        "text/css; charset=utf-8",
        "text/xml; charset=utf-8",
        "application/javascript; charset=utf-8",
        "application/json; charset=utf-8",
        "image/png",
        "image/jpeg",
        "image/svg+xml",
        "image/x-icon",
        "image/bmp",
        "image/gif",
        "video/mp4",
        "video/webm",
        "font/ttf",
        "font/otf",
        "font/collection",
        "font/woff",
        "font/woff2"}
}
*/
