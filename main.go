package main

import (
	"fmt"
	"net/http"
)

func main() {

	apiMux := http.NewServeMux()

	apiMux.HandleFunc("/v1/ping", pinghandler)
	apiMux.HandleFunc("/v1/greet", greethandler)

	mainMux := http.NewServeMux()

	mainMux.Handle("/api/", http.StripPrefix("/api", apiMux))

	mainMux.HandleFunc("/method-inspector", methodinspectorhandler)
	mainMux.HandleFunc("/echo", echohandler)
	mainMux.HandleFunc("/headers", headershandler)
	mainMux.HandleFunc("/form", formhandler)
	mainMux.HandleFunc("/status", statushandler)
	mainMux.HandleFunc("/render", renderhandler)

	fmt.Println("server running on http://localhost:8080")

	http.ListenAndServe(":8080", mainMux)
}
