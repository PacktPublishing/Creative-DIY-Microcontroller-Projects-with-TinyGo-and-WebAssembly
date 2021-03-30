package main

import (
	"log"
	"net/http"
	"strings"
)

const dir = "Chapter07/html"

func main() {
	fs := http.FileServer(http.Dir(dir))
	log.Print("Serving " + dir + " on http://localhost:8080")
	err := http.ListenAndServe(":8080", http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		resp.Header().Add("Cache-Control", "no-cache")
		if strings.HasSuffix(req.URL.Path, ".wasm") {
			resp.Header().Set("content-type", "application/wasm")
		}

		if strings.HasSuffix(req.URL.Path, ".css") {
			resp.Header().Set("content-type", "text/css")
		}

		if strings.HasSuffix(req.URL.Path, ".jpg") {
			resp.Header().Set("content-type", "image/jpeg")
		}

		fs.ServeHTTP(resp, req)
	}))

	if err != nil {
		println("failed to server http requests:", err.Error())
	}
}
