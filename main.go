// main.go uses the Google Cloud
// App Engine to host the
// web browsing task.
package main

import (
	"net/http"

	"google.golang.org/appengine"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})
	appengine.Main() // Starts the server to receive requests.
}
