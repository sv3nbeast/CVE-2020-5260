package main

import (
        "log"
        "net/http"
)

func h(w http.ResponseWriter, r *http.Request) {
        username, password, ok := r.BasicAuth()
        if ok {
                log.Printf("user: %v password: %v\n", username, password)
                w.WriteHeader(200)
                return
        }
        w.Header().Set("WWW-Authenticate", `Basic realm="foo"`)
        http.Error(w, "Not authorized", 401)
        return
}

func main() {
        http.HandleFunc("/", h)
        err := http.ListenAndServe(":8088", nil)
        if err != nil {
                log.Fatal("ListenAndServe: ", err)
        }
}