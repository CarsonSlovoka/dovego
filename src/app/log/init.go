package log

import (
    log2 "github.com/CarsonSlovoka/dovego/pkg/log"
    "log"
    "os"
)

var (
    Trace *log.Logger
    Error *log.Logger
)

// log
func InitLog(filepath string) *os.File {
    file, err := os.OpenFile(
        filepath,
        os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666,
    )
    if err != nil {
        Trace, Error = log2.New(os.Stdout, os.Stderr, false)
        Error.Printf("Unable to create and/or open log file. Will log to Stdout and Stderr. Error: %v", err)
        return nil
    }
    Trace, Error = log2.New(file, file, false)
    return file
}
