package main //GUIsocket

import (
	//"./gui"
	"fmt"
	"github.com/go-chi/chi"
	_ "github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"sync"
	"time"
	//"./api"
	//"github.com/suchy1105/GUIcontroler/config"
	"net/http"

	//	"github.com/suchy1105/GUIcontroler/config"
)
//var g *api.GuiState
func main() {
	defer run()
}
func run() {


	var frontendRouter chi.Router = chi.NewRouter()
	var wg sync.WaitGroup
	//data:=api.NewGuiState()

	apiServer := http.Server{
		Addr:           ":8080",
		Handler:        frontendRouter,
		ReadTimeout:    360 * time.Second,
		WriteTimeout:   360 * time.Second,
		MaxHeaderBytes: 1 << 20,
		// discard error logs
		//ErrorLog: golog.New(ioutil.Discard, "", 0),
	}
	wg.Add(1)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		log.Log().Msgf("starting backend api server on %s", apiServer.Addr)
		err := apiServer.ListenAndServe()
		if err != nil {
			if err != http.ErrServerClosed {
				log.Warn().Err(err).Caller().Msg("error while closing api server")
			}
		}
	}(&wg)
	//backendRouter.Route("/api/control", api.ControlAPI())
	//frontendRouter.Route("/frontend", api.FrontendAPI()	)
	frontendRouter.Get("/", get)

//	var err error
//	var conf config.Configuration
//	conf.GetConf()
//
 timer()
 //gui.GUI()




	fmt.Println("lisetner")


}
func  timer() {
	for {
		fmt.Println("work in progress")
		time.Sleep(1 * time.Second)
	}
}
func get(w http.ResponseWriter, r *http.Request) {
	fmt.Println("404x")
	w.WriteHeader(http.StatusOK)


}