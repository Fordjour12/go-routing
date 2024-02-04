package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	//	"github.com/julienschmidt/httprouter"
	"github.com/gorilla/mux"
)

func ArticleHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "ID: %v\n", vars["id"])
}

func ArticleCategoryHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category: %v\n", vars["category"])
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/articles/{category}/", ArticleCategoryHandler)
	r.HandleFunc("/articles/{category}/{id:[0-9]+}", ArticleHandler).Name("article")

	//HandleFunc chaning is possible (HandlerFunc and not HandleFunc)
	// r.HandleFunc("/articles/{cate}/").HandlerFunc(ArticleCategoryHandler)

	// Reverse mapping(Registered URLs)
	url, err := r.Get("article").URL("category", "books", "id", "123")
	fmt.Printf("%v, %v\n", url, err)

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// standard pratices to enforce timeout
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

}
