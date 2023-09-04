package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		product := r.URL.Query().Get("product")
		if id != "" {
			w.Write([]byte(id + " " + product))
		} else {
			w.Write([]byte("Digite um nome"))
		}

	})
	r.Get("/{productName}", func(w http.ResponseWriter, r *http.Request) {
		param := chi.URLParam(r, "productName")
		w.Write([]byte(param))
	})

	r.Get("/json", func(w http.ResponseWriter, r *http.Request) {

		obj := map[string]string{"message": "sucess"}
		render.JSON(w, r, obj)
	})
	http.ListenAndServe(":5000", r)

}
