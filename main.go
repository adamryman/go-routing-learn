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

//One off that we can pass things to!
func timeHandlerClosure(format string) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		tm := time.Now().Format(format)
		w.Write([]byte("The time is " + tm + " | CLOSURE!"))
	}
	return http.HandlerFunc(fn)
}

//different closure
func timeHandlerDifferentClosure(format string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tm := time.Now().Format(format)
		w.Write([]byte("The time is: " + tm))
	})
}

//Third closure
func timeHandlerThirdClosure(format string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tm := time.Now().Format(format)
		w.Write([]byte("The time is: " + tm))
	}
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

	//even simpler
	mux.HandleFunc("/time/simpler", timeHandlerOneOff)

	//passing through closure
	thClosure := timeHandlerClosure(time.RFC1123)
	mux.Handle("/time/closure", thClosure)

	//second closure
	thDifferent := timeHandlerDifferentClosure(time.RFC1123)
	mux.Handle("/time/different", thDifferent)

	//third closure
	thThirdClosure := timeHandlerThirdClosure(time.RFC1123)
	mux.Handle("/time/third", thThirdClosure)

	log.Println("Listening...")
	http.ListenAndServe(":3000", mux)
}
