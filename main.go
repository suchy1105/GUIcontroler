package main//GUIsocket

import (
	"github.com/go-chi/chi"
	_ "github.com/rs/zerolog"
	"github.com/suchy1105/GUIcontroler/api"
	"github.com/suchy1105/GUIcontroler/config"
	"net/http"

	//	"github.com/suchy1105/GUIcontroler/config"
)

func main() {
	defer run()
}
func run() {
	var err error
	var conf config.Configuration
	conf.GetConf()


	router := chi.NewRouter()


	router.Route("/apiv1", func(router chi.Router) {

		router.Post("/message", api.PostMessage)


		router.Get("/message/", api.GetMessages)
		router.Get("/checkhealth", api.CheckHealth)
		router.NotFound(api.NotFound)
	})
	http.ListenAndServe(":1600", router)

}