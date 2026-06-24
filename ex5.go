/*
Exercise 5: Status Code Factory
Goal
Build a /status endpoint that accepts a code query parameter containing any HTTP status code number. The server must respond using that exact status code. This forces you to think about how status codes are set — and when they cannot be changed.

Key Tasks
●     Read the code query parameter using r.URL.Query().Get("code").
●     If code is missing or empty — return 400 with the message: "code parameter is required".
●     Convert code to an integer using strconv.Atoi(). If conversion fails — return 400 with: "code must be a valid integer".
●     If the integer is not between 100 and 599 — return 400 with: "code must be a valid HTTP status code (100–599)".
●     If the code is valid — respond with that exact status code using w.WriteHeader(code) and a body message: "Responding with status [code]".

Critical rule —
You must call w.WriteHeader(code) BEFORE writing anything to w with w.Write()
or fmt.Fprintf(). If you write the body first, Go automatically sends a 200 header
and you cannot change it afterwards. The order is: WriteHeader → then Write.
Test this deliberately: call w.Write() first, then w.WriteHeader(404). What does
curl -v show you? Write your observation in a comment in your file.

Stretch — do this after the core task works
●     After calling w.WriteHeader(), append a descriptive name to the body message.
○     ?code=404 → "Responding with status 404 Not Found"
○     Use http.StatusText(code) to get the official status name.
*/
package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func statushandler(w http.ResponseWriter, r *http.Request) {
	querycode := r.URL.Query().Get("code")
	if querycode == "" {
		http.Error(w, "code parameter is required", http.StatusBadRequest)
		return
	}

	code, err := strconv.Atoi(querycode)
	if err != nil {
		http.Error(w, "code must be a valid integer", http.StatusBadRequest)
		return
	}

	if code < 100 || code > 599 {
		http.Error(w, "code must be a valid HTTP status code (100–599)", http.StatusBadRequest)
		return
	}

	w.WriteHeader(code)

	//	w.Write([]byte("Responding with status %v", code))

	fmt.Fprintf(w, "Responding with status %v %v", code, http.StatusText(code))

}

func main() {
	http.HandleFunc("/status", statushandler)
	fmt.Println("server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
