package main

import (
    "fmt"
    "net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) { //This function reads all the headers in the request and echoes them into the response body
    for name, headers := range req.Header {
        for _, h := range headers {
            fmt.Fprintf(w, "%v: %v\n", name, h)
        }
    }
}

func main() {
    // We register the handlers on server routes. It sets up the default router in the net/http package and takes a function as an arg (HandleFunc)
    http.HandleFunc("/hello", hello)
    http.HandleFunc("/headers", headers)

    http.ListenAndServe(":8090", nil) //Port and handler as args. nil tells it to use the default router we've just set up
}

// A fundamental concept in net/http servers is handlers. A handler is an object implementing the http.Handler interface. A common way to write a handler is by using the http.HandlerFunc adapter on functions with the appropriate signature
// Functions serving as handlers take a http.ResponseWriter and a http.Request as args. The response writer is used to fill the HTTP response
