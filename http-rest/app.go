package httprest

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type User struct {
	FirstName string    `json:"firstname"`
	LastName  string    `json:"lastname"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Go REST")
}

func getUserInfoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(w, "User id: ", vars["id"])
}

func RestHandler() http.Handler {
	mux := mux.NewRouter()

	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/users/{id:[0-9]+}", getUserInfoHandler)

	http.ListenAndServe(":8000", mux)
	return mux
}
