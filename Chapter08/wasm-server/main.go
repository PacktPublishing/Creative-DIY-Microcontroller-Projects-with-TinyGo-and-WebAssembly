package main

import (
	"log"
	"net/http"
	"strings"
)

const dir = "Chapter08/html"

var fs = http.FileServer(http.Dir(dir))

func main() {

	log.Print("Serving " + dir + " on http://localhost:8080")
	http.ListenAndServe(":8080", http.HandlerFunc(handleRequest))

}

func handleRequest(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Add("Cache-Control", "no-cache")
	if strings.HasSuffix(req.URL.Path, ".wasm") {
		resp.Header().Set("content-type", "application/wasm")
	}

	requestURI := req.URL.RequestURI()
	if strings.Contains(requestURI, "dashboard") ||
		strings.Contains(requestURI, "login") {
		http.Redirect(resp, req, "http://localhost:8080", http.StatusMovedPermanently)
		return
	}

	fs.ServeHTTP(resp, req)
}
