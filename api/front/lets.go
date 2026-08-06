package handler

import (
	"fmt"
	"net/http"
)

// Handler is the serverless entrypoint Vercel uses
func Handler(w http.ResponseWriter, r *http.Request) {
	// Set the header so the browser knows we are sending HTML
	w.Header().Set("Content-Type", "text/html")

	// Send an HTML snippet back to the HTMX frontend
	fmt.Fprintf(w, "<div class='post'><h3>Loaded dynamically from Go!</h3><p>This was sent from a Vercel Serverless function.</p></div>")
}
