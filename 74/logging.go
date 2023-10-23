package main

import (
    "fmt"
    "bytes"
    "log"
    "os"
    "log/slog"
)

func main() {
    log.Println("standard logger")

    log.SetFlags(log.LstdFlags | log.Lmicroseconds) //Loggers can be configured with flags to set their output format
    log.Println("with micro")

    log.SetFlags(log.LstdFlags | log.Lshortfile)
    log.Println("with file/line")

    mylog := log.New(os.Stdout, "my:", log.LstdFlags)
    log.Println("from mylog")

    mylog.SetPrefix("ohmy:")
    mylog.Println("from mylog")

    var buf bytes.Buffer
    buflog := log.New(&buf, "buf:", log.LstdFlags)

    buflog.Println("hello")

    fmt.Print("from buflog:", buf.String())

    jsonHandler := slog.NewJSONHandler(os.Stderr, nil)
    myslog := slog.New(jsonHandler)
    myslog.Info("hi there")

    myslog.Info("hello again", "key", "val", "age", 25)
}
