/*
Exercise 6: The API Subtree
Goal
Build a mini API under the /api/v1/ path prefix using a separate http.ServeMux —
not the default mux. This mux handles only /api/v1/ routes.
Register two handlers inside it: /api/v1/ping and /api/v1/greet.
Mount the whole submux onto the main server at /api/.

What a ServeMux subtree is
Go's http.ServeMux uses a simple rule:
a pattern that ends in / matches any path that starts with it.
This makes it possible to create a sub-router —
a separate ServeMux that handles a subtree of routes —
and mount it onto the main mux with http.StripPrefix.
The main mux passes all /api/ requests to the submux,
which handles them as if the /api prefix did not exist.

// The pattern for a subtree — note the trailing slash
mainMux.Handle("/api/", http.StripPrefix("/api", apiMux))

// Inside apiMux, routes are registered WITHOUT /api
apiMux.HandleFunc("/v1/ping", pingHandler)
apiMux.HandleFunc("/v1/greet", greetHandler)

// A request to /api/v1/ping:
// mainMux strips "/api" → apiMux sees "/v1/ping" → routes to pingHandler

Key Tasks
●     Create a new mux: apiMux := http.NewServeMux()
●     Register /v1/ping on apiMux — responds with "pong" in plain text.
●     Register /v1/greet on apiMux — reads a name query parameter
and responds with "Greetings, [name]!" or "Greetings, Stranger!" if name is missing.
●     Mount apiMux onto the main mux:
○     mainMux := http.NewServeMux()
○     mainMux.Handle("/api/", http.StripPrefix("/api", apiMux))
●     Start the server using http.ListenAndServe(":8080", mainMux) — not nil.
●     Visiting /api/v1/ping must return "pong".
Visiting /api/v1/greet?name=Zion must return "Greetings, Zion!".

	Why this matters —

ascii-art-web mounts handlers at / and /ascii-art. A real application groups
related routes under a prefix — /api/, /admin/, /v2/. ServeMux subtrees are
Go's standard way to do this without an external router. Understanding
StripPrefix is the foundation of any multi-route Go server.

Think about — write your answers in comments at the top of your file
●     What happens if you use mainMux.Handle("/api", ...)
without the trailing slash? Try it.
●     What does http.StripPrefix do —
what would apiMux receive if you did NOT use StripPrefix?
●     What does it mean that http.ListenAndServe takes a Handler interface —
and how does a *http.ServeMux satisfy that interface?
*/
package main

import (
	"fmt"
	"net/http"
)

func pinghandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "pong")
}

func greethandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "Stranger"
	}
	fmt.Fprintf(w, "Greetings, %s!", name)
}
