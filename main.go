package main//GUIsocket

import (
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
	"config"
	"api/api"
)

func main() {
	defer run()
}
func run() {



	router := chi.NewRouter()


	http.ListenAndServe(":8080", router)

}