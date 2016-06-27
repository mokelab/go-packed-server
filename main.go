package main

import (
	"./data"
	"flag"
	"fmt"
	"net/http"
)

func makeHandler(c *data.Cache) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		e, ok := c.Entry[r.RequestURI]
		if !ok {
			w.WriteHeader(404)
			fmt.Fprint(w, "Not found")
			return
		}
		w.WriteHeader(200)
		w.Header().Set("Content-Type", e.ContentType)
		w.Write(e.Get())
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, r.RequestURI)
}

func main() {
	var port int
	flag.IntVar(&port, "p", 3000, "Port number")
	flag.Parse()
	fmt.Printf("Running server on port %d\n", port)
	fmt.Printf("Open http://localhost:%d/index.html on your browser\n", port)

	c := data.Create()
	http.HandleFunc("/", makeHandler(c))
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
