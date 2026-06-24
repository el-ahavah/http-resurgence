/*
Exercise 1: The Method Inspector
Goal
Build a /method-inspector endpoint that reads the HTTP method of every incoming request and echoes it back in a descriptive sentence. No method should be rejected — this handler accepts everything and reports what it sees.

Key Tasks
●     Register a /method-inspector handler using http.HandleFunc.
●     Read the request method using r.Method.
●     Respond with a plain text message that includes the method name.
○     GET request → "You made a GET request."
○     POST request → "You made a POST request."
○     Any other method → "You made a [METHOD] request."
●     Do not hardcode each method with its own if/else branch — use the value of r.Method directly in your response string.

Why this matters —
In ascii-art-web your POST /ascii-art handler must distinguish incoming methods
before doing any work. This exercise builds the muscle of reading r.Method
cleanly and using it — not just checking it.
*/
package main

import (
	"fmt"
	"net/http"
)

func methodinspectorhandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You made a %v request", r.Method)
}

func main() {
	http.HandleFunc("/method-inspector", methodinspectorhandler)
	fmt.Println("server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
