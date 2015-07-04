package main

import (
	"log"
	"net/http"
	"time"
)

//custom type
type timeHandler struct {
	format string
}

// ServeHTTP for custom time which makes it a handler
func (th *timeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tm := time.Now().Format(th.format)
	w.Write([]byte("The time is: " + tm))
}

//One off method handler
func timeHandlerOneOff(w http.ResponseWriter, r *http.Request) {
	tm := time.Now().Format(time.RFC1123)
	w.Write([]byte("The time is: " + tm))
}

func main() {
	mux := http.NewServeMux()

	rh := http.RedirectHandler("http://example.org", 307)
	mux.Handle("/foo", rh)

	//time
	th1123 := &timeHandler{format: time.RFC1123}
	mux.Handle("/time/rfc1123", th1123)

	//reuse!
	th3339 := &timeHandler{format: time.RFC3339}
	mux.Handle("/time/rfc3339", th3339)

	//one off method
	thOneOff := http.HandlerFunc(timeHandlerOneOff)
	mux.Handle("/time/oneoff", thOneOff)

	log.Println("Listening...")
	http.ListenAndServe(":3000", mux)
}
