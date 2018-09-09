// main.go uses the Google Cloud
// App Engine to host the
// web browsing task.
package main

import (
	"html/template"
	"net/http"

	"google.golang.org/appengine"
)

func main() {
	http.HandleFunc("/", handler)
	appengine.Main() // Starts the server to receive requests.
}

func handler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("index.html"))
	t.Execute(w, nil)
}
