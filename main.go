package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"
)

// Article holds the data for your proposals
type Article struct {
	ID      string
	Title   string
	Content string
}

// Simulated database based on your website's content
var articles = map[string]Article{
	"eliminate-taxes": {
		ID:      "eliminate-taxes",
		Title:   "Eliminate Direct Taxes on Wages",
		Content: "hello ",
	},
	"eliminate-filings": {
		ID:      "eliminate-filings",
		Title:   "Eliminate 160 million tax filings",
		Content: "This proposal would eliminate tax filings for 160 million filers, which is the majority of US citizens, while keeping indirect taxes intact...",
	},
	"original-intent": {
		ID:      "original-intent",
		Title:   "Returns the income tax to its original intent",
		Content: "This is proposal would eliminate tax filings for 160 million filers, which is the majority of US citizens, while keeping indirect taxes intact...",
	},
	"reduce-intrusiveness": {
		ID:      "reduce-intrusive-affairs",
		Title:   "Reduce intrusiveness into personal affairs",
		Content: "This proposal would eliminate tax filings for 160 million filers, which is the majority of US citizens, while keeping indirect taxes intact...",
	},
	"administrative-costs": {
		ID:      "administrative-costs",
		Title:   "Significant reduction of administrative costs",
		Content: "This proposal would eliminate tax filings for 160 million filers, which is the majority of US citizens, while keeping indirect taxes intact...",
	},
	"record-keeping": {
		ID:      "record-keeping",
		Title:   "Frees up billions of hours of record keeping",
		Content: "This proposal would eliminate tax filings for 160 million filers, which is the majority of US citizens, while keeping indirect taxes intact...",
	},
	"economic-efficiency": {
		ID:      "economic-efficiency",
		Title:   "Promotes Economic Efficiency",
		Content: "This proposal would eliminate tax filings for 160 million filers, which is the majority of US citizens, while keeping indirect taxes intact...",
	},
	"sham-court": {
		ID:      "sham-court",
		Title:   "Hyton Court Sham",
		Content: "This proposal would eliminate tax filings for 160 million filers, which is the majority of US citizens, while keeping indirect taxes intact...",
	},
	"constitutional-intent": {
		ID:      "constitutional-intent",
		Title:   "Return to Constitutional intect on taxation",
		Content: "This proposal would eliminate tax filings for 160 million filers, which is the majority of US citizens, while keeping indirect taxes intact...",
	},
}

func main() {
	// 1. Parse your HTML templates
	tmpl := template.Must(template.ParseFiles("index.html", "articles.html"))

	// 2. Route: Full Page Load (e.g., when a user first visits the site)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Pass a default article to display on first load
		data := struct {
			DefaultArticle Article
		}{
			DefaultArticle: articles["eliminate-taxes"],
		}

		err := tmpl.ExecuteTemplate(w, "index.html", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	// 3. Route: HTMX Article Requests
	http.HandleFunc("/article/", func(w http.ResponseWriter, r *http.Request) {
		// Extract the article ID from the URL (e.g., "/article/eliminate-filings")
		id := strings.TrimPrefix(r.URL.Path, "/article/")

		article, exists := articles[id]
		if !exists {
			http.NotFound(w, r)
			return
		}

		// Execute ONLY the "article" template fragment, not the whole page
		err := tmpl.ExecuteTemplate(w, "article", article)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
