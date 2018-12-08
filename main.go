package main

import (
    "github.com/gorilla/mux"
    "os"
    "fmt"
    "net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("hello!"))
}

func main() {
    router := mux.NewRouter()
    router.HandleFunc("/", hello)

    port := os.Getenv("PORT")
    if port == "" {
        port = "5000"
    }
    
    fmt.Print("Running server in port: " + port)
    if err := http.ListenAndServe(":" + port, router); err != nil {
        fmt.Print(err)
    }

}