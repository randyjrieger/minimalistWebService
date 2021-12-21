package main

import (
	"log"
	"net/http"
)

func main() {
	//  args root path, function to call
	// the function is going to be an anonyous function defined in-place
	//   http.ResponseWrite - what we will be writing our response to
	//   http.Request - get URL, Header, Body, etc.)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	// Start up the web server
	// http.ListenAndServe - exactly how it sounds...
	//   the port to listen on
	//   handler to use - nil says to use the default handler
	err := http.ListenAndServe(":3001", nil)

	if err != nil {
		log.Fatal(err)
	}
}
