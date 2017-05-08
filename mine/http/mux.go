
package main

import (
    "fmt"
    "net/http"
)

type myHandler struct {
    myvar string
}

func (mh *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, mh.myvar + "\n")
}

func main() {
    mymy := "My words"

    apiMux := http.NewServeMux()
    apiMux.Handle("/", &myHandler{myvar: mymy})

    http.ListenAndServe(":8888", apiMux)
}
