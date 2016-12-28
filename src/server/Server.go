package main 

import (
    "net/http"
    "fmt"
    "strings"
    "log"
    "encoding/json"
	"net/url"
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

func verifyNumber(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
    fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	    for k, v := range r.Form {
        fmt.Println("key:", k)
        fmt.Println("val:", strings.Join(v, ""))
    }
	fmt.Fprintf(w, "Hello verifyNumber!") 
}

func verifyPhoneNumber(w http.ResponseWriter, r *http.Request) {
	phone := "14158586273"
	// QueryEscape escapes the phone string so
	// it can be safely placed inside a URL query
	safePhone := url.QueryEscape(phone)

	url := fmt.Sprintf("http://apilayer.net/api/validate?access_key=b912c69a00eaffdeead23546778cc785&number=%s", safePhone)

	// Build the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}

	// For control over HTTP client headers,
	// redirect policy, and other settings,
	// create a Client
	// A Client is an HTTP client
	client := &http.Client{}

	// Send the request via a client
	// Do sends an HTTP request and
	// returns an HTTP response
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return
	}

	// Callers should close resp.Body
	// when done reading from it
	// Defer the closing of the body
	defer resp.Body.Close()

	// Fill the record with the data from the JSON
	var record Numverify

	// Use json.Decode for reading streams of JSON data
	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Println(err)
	}

	fmt.Println("Phone No. = ", record.InternationalFormat)
	fmt.Println("Country   = ", record.CountryName)
	fmt.Println("Location  = ", record.Location)
	fmt.Println("Carrier   = ", record.Carrier)
	fmt.Println("LineType  = ", record.LineType)
	
	fmt.Fprintf(w, "Phone No. = "+ record.InternationalFormat + "\n")
	fmt.Fprintf(w, "Country   = "+ record.CountryName + "\n")
	fmt.Fprintf(w, "Location  = "+ record.Location + "\n")
	fmt.Fprintf(w, "Carrier   = "+ record.Carrier + "\n")
	fmt.Fprintf(w, "LineType  = "+ record.LineType + "\n")
	
}
//func main() {
//    http.HandleFunc("/defaultHandler", defaultHandler) // set router
//    http.HandleFunc("/verifyNumber&number=$number", verifyNumber) // set router
//    err := http.ListenAndServe(":8080", nil) // set listen port
//    if err != nil {
//        log.Fatal("ListenAndServe: ", err)
//    }
//}

type Numverify struct {
	Valid               bool   `json:"valid"`
	Number              string `json:"number"`
	LocalFormat         string `json:"local_format"`
	InternationalFormat string `json:"international_format"`
	CountryPrefix       string `json:"country_prefix"`
	CountryCode         string `json:"country_code"`
	CountryName         string `json:"country_name"`
	Location            string `json:"location"`
	Carrier             string `json:"carrier"`
	LineType            string `json:"line_type"`
}

func main() {
	// Test can it call methods from other go files.
	http.HandleFunc("/defaultHandler", defaultHandler) // set router
    http.HandleFunc("/verifyNumber&number=$number", verifyNumber) // set router
    http.HandleFunc("/verifyPhoneNumber", verifyPhoneNumber)
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