/*
Exercise 2: The Echo Chamber
Goal
Create an /echo endpoint that only accepts POST requests. When a client sends a POST with a body, read the entire body and send it straight back. The response must be exactly what was sent — nothing added, nothing removed.
 
Key Tasks
●     Reject any non-POST request with http.StatusMethodNotAllowed (405).
●     Read the full request body using io.ReadAll(r.Body).
●     Always defer r.Body.Close() immediately after checking for an error — not at the end of the function.
●     If the body is completely empty (zero bytes), return a 400 Bad Request with the message "body cannot be empty".
●     Write the body content back to w exactly as received.

Think about —
What does io.ReadAll return if the request has no body at all?
Is len(body) == 0 the same as body == nil? Try both and see.
What happens to r.Body if you read it twice without closing it? 

Stretch — do this after the core task works
●     Add a response header: Content-Type: text/plain before writing the body back.
○     Use w.Header().Set("Content-Type", "text/plain") — and call it before w.Write().
○     What happens if you call w.Header().Set() after w.Write()? Try it and explain what you observe.
*/