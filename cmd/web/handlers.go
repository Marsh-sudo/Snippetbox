package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	files := []string {
		"./ui/html/home.html",
		"./ui/html/base.html",
		"./ui/html/footer.html",
	}

	ts,err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w,err)
		return
	}

	err = ts.Execute(w,nil)
	if err != nil {
		app.serverError(w,err)
	}

	w.Write([]byte("Hello from Snippetbox"))
}

func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(r.URL.Query().Get("id"))
    if err != nil || id < 1 {
        app.notFound(w)
        return
}
    fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}


func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
        w.Header().Set("Allow", "POST")
        app.clientError(w, http.StatusMethodNotAllowed)
        return
}
    w.Write([]byte("Create a new snippet..."))
}