package main 

import (
    "net/http"
    "fmt"
     "strings"
    "log"
)

// Guides to refer : https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/03.2.html
// Default Request Handler
func defaultHandler(w http.ResponseWriter, r *http.Request) {
    //fmt.Fprintf(w, "<h1>Hello %s!</h1>", r.URL.Path[1:])
    r.ParseForm()  // parse arguments, you have to call this by yourself
    fmt.Println(r.Form)  // print form information in server side
    fmt.Println("path", r.URL.Path)
    fmt.Println("scheme", r.URL.Scheme)
    fmt.Println(r.Form["url_long"])
    for k, v := range r.Form {
        fmt.Println("key:", k)
        fmt.Println("val:", strings.Join(v, ""))
    }
    fmt.Fprintf(w, "Hello Ankush!") // send data to client side
}

func main() {
    http.HandleFunc("/", defaultHandler) // set router
    err := http.ListenAndServe(":8080", nil) // set listen port
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}


/// Write a script to kill the server afterwards
///netstat -anp tcp | grep 8080
///terminate(){
///  lsof -P | grep ':8080' | awk '{print $2}' | xargs kill -9 
///}