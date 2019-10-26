package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, SecureHeaders)
	dynamicMiddleware := alice.New(app.session.Enable)

	mux := pat.New()
	mux.Get("/", dynamicMiddleware.ThenFunc(app.home))
	mux.Get("/snippet/create", dynamicMiddleware.ThenFunc(app.createSnippetForm))
	mux.Post("/snippet/create", dynamicMiddleware.ThenFunc(app.createSnippet))
	mux.Get("/snippet/:id", dynamicMiddleware.ThenFunc(app.showSnippet))

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	return standardMiddleware.Then(mux)
	// r := mux.NewRouter()
	// r.HandleFunc("/", app.home).Methods("GET")
	// r.HandleFunc("/snippet/create", app.createSnippetForm).Methods("GET")
	// r.HandleFunc("/snippet/create", app.createSnippet).Methods("POST")
	// r.HandleFunc("/snippet/{id}", app.showSnippet).Methods("GET")

	// fileServer := http.FileServer(http.Dir("./ui/static/"))
	// r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fileServer))

	// return standardMiddleware.Then(r)
}

func (app *application) hello(w http.ResponseWriter, r *http.Request) {
	app.infoLog.Print("Hello")
}
