Core Functions — net/http
 
Function / Type
What it does
http.NewServeMux()
Creates a new, independent request multiplexer (router). Pass it to ListenAndServe instead of nil to use it as the main router.
http.StripPrefix(prefix, handler)
Returns a handler that strips the given prefix from the request URL before passing it to the next handler. Used to mount a submux.
http.ListenAndServe(addr, handler)
Starts the server. Pass a *ServeMux or nil (uses DefaultServeMux). Never returns unless the server crashes.
w.Header().Set(key, value)
Sets a response header. Must be called BEFORE w.WriteHeader() or any w.Write() call.
w.WriteHeader(code)
Sends the HTTP status code. Must be called BEFORE w.Write(). Calling it after w.Write() has no effect.
http.StatusText(code)
Returns the official status text for a code. e.g. http.StatusText(404) returns "Not Found".
template.Must(t, err)
Wraps a template parse call. Panics if err is not nil — use only at startup, never inside a handler.
template.New(name).Parse(str)
Parses an HTML template from a string. Returns (*Template, error).
tmpl.Execute(w, data)
Renders the template with data and writes the result to w. Returns an error if rendering fails.


The ServeMux Subtree — Pattern 
func main() {
    // The main mux — receives all requests
    mainMux := http.NewServeMux()

    // A sub-mux — handles only /api/* routes
    apiMux := http.NewServeMux()
    apiMux.HandleFunc("/v1/ping", pingHandler)
    apiMux.HandleFunc("/v1/greet", greetHandler)

    // Mount apiMux under /api/ — StripPrefix removes "/api"
    // so apiMux sees /v1/ping instead of /api/v1/ping
    mainMux.Handle("/api/", http.StripPrefix("/api", apiMux))

    // Register other top-level routes on mainMux
    mainMux.HandleFunc("/ping", pingHandler)

    http.ListenAndServe(":8080", mainMux)
}

Response Header Order — The Rule
 
Step
Call
Notes
1
w.Header().Set("Content-Type", "text/html")
Set ALL headers here. Order matters — before everything.
2
w.WriteHeader(http.StatusCreated)
Set the status code if it is not 200. Only call once.
3
fmt.Fprintf(w, "...") or tmpl.Execute(w, data)
Write the body last. Writing the body locks headers and status.

 

Request Fields at a Glance
 
Field / Method
Purpose
Returns when missing
r.Method
The HTTP method — GET, POST, etc.
Never missing — always set
r.URL.Query().Get(key)
A query parameter value
Empty string ""
r.Header.Get(key)
A request header value
Empty string ""
r.FormValue(key)
A form field from POST body or query string
Empty string ""
io.ReadAll(r.Body)
The raw request body as []byte
Empty slice, nil error
r.Body.Close()
Frees the connection — always defer this
—


