package main

import (
    "os"
    "fmt"
    "net/http"
)

func main() {
    var port = os.Getenv("PORT_WWW")

    if len(port) == 0 {
        fmt.Println("$PORT_WWW not set defaulting to 8080", port)
        port = "8080"
    }
    fmt.Println("Using Port:", port)
    http.HandleFunc("/", handler)
    http.ListenAndServe(":"+port, nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello, GO world!")
}