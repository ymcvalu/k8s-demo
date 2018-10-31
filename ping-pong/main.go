package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
)

var port *int

func main() {

	port = flag.Int("port", 8080, "port to listen on")
	flag.Parse()

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		log.Printf("request %s from %s", req.URL.String(), req.RemoteAddr)
		resp.Header().Add("Content-Type", "text/plain")
		resp.WriteHeader(200)
		resp.Write([]byte{'p', 'o', 'n', 'g', '!'})
	})

	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("server listening on %d", *port)
	http.Serve(ln, mux)
}
