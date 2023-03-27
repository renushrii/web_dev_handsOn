/*
ListenAndServe on port ":8080" using the default ServeMux.

Use HandleFunc to add the following routes to the default ServeMux:

"/"
"/dog/"
"/me/

Add a func for each of the routes.

Have the "/me/" route print out your name.
*/

package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", kitty)
	http.HandleFunc("/panda/", doggy)
	http.HandleFunc("/me/", myName)
	http.ListenAndServe(":8080", nil)
}

func kitty(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "kitty ran")
}

func doggy(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "doggy ran")
}

func myName(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello Renu")
}
