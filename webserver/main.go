//This program starts a simple webserver on port 80 which replies with hello world
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})
	//output an erro if we can't start the server
	if err := http.ListenAndServe(":80", nil); err != nil {
		panic(err)
	}

}
