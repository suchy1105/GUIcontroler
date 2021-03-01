package main//GUIsocket

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"io/ioutil"
	"os"
	"strings"
	"sync"
	"time"
	golog "log"

	_ "github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	//"github.com/go-chi/chi"
	"net/http"
	"github.com/suchy1105/GUIcontroler/api"
//	"github.com/suchy1105/GUIcontroler/config"
)

func main() {
	os.Exit(run())
}
func run() int {
	fmt.Println("karil")
	//router:=
	var guiRouter chi.Router = chi.NewRouter()

	guiRouter.Post("v1/message/", postData())

	guiRouter.Get("v1/message/", getData())
	guiRouter.Get("v1/health", checkHealth())

	var wg sync.WaitGroup

	apiServer := http.Server{
		Addr:           ":1009",
		Handler:        guiRouter,
		ReadTimeout:    360 * time.Second,
		WriteTimeout:   360 * time.Second,
		MaxHeaderBytes: 1 << 20,
		// discard error logs
		ErrorLog: golog.New(ioutil.Discard, "", 0),
	}
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		log.Log().Msgf("starting GUI server", apiServer.Addr)
		err := apiServer.ListenAndServe()
		if err != nil {
			if err != http.ErrServerClosed {
				log.Warn().Err(err).Caller().Msg("error while closing api server")
			}
		}
	}(&wg)

	return 0
}

func checkHealth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}
}
//Get API
func getData() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		reqEmail := strings.Split(r.RequestURI, "/api/message/")
		fmt.Println(reqEmail[1])
		w.WriteHeader(http.StatusTeapot)
	}

}
//Post API
func postData() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var message api.Message
		json.NewDecoder(r.Body).Decode(&message)

		//TODO
		w.Write([]byte(`{"message": "post called"}`))

		w.WriteHeader(http.StatusCreated)
	}
}

//NotFound 404
func NotFound(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusNotFound)


}