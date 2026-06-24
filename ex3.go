/*
Exercise 3: Header Detective
Goal
Create a /headers endpoint that inspects two specific request headers: X-Custom-Token and Content-Type. The handler reads both, reports what it found, and enforces a rule about one of them.

Key Tasks
●     Read X-Custom-Token using r.Header.Get("X-Custom-Token").
●     If X-Custom-Token is missing or empty — respond with 400 Bad Request and the message: "X-Custom-Token header is missing".
●     If X-Custom-Token is present — respond with a message that includes its value. Example: "Token received: abc123".
●     Also read Content-Type and append it to the response. If it is missing, append "Content-Type not provided".
●     The full response for a valid request must look like this:
○     Token received: abc123
○     Content-Type: application/json

Why this matters —
ascii-art-web reads r.Header indirectly through template and form handling.
Understanding how headers work — and what happens when they are absent —
prepares you for writing handlers that behave correctly under any input.

Stretch — do this after the core task works
●     What does r.Header.Get() return for a header key that was never sent? Write a one-sentence answer in a comment at the top of your file.
●     Is r.Header.Get("x-custom-token") the same as r.Header.Get("X-Custom-Token")? Find out.
*/
package main

import (
	"fmt"
	"net/http"
)

func headershandler(w http.ResponseWriter, r *http.Request) {
	customToken := r.Header.Get("X-Custom-Token")

	if customToken == "" {
		http.Error(w, "X-Custom-Token header is missing", http.StatusBadRequest)
		return
	} else {
		fmt.Fprintf(w, "Token received: %v", customToken)
	}

	contentType := r.Header.Get("Content-Type")

	if contentType == "" {
		http.Error(w, "Content-Type not provided", http.StatusBadRequest)
		return
	} else {
		fmt.Fprintf(w, "Content-Type: %v", contentType)
		return
	}

}
