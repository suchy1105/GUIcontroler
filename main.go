package main//GUIsocket

import (
	"github.com/go-chi/chi"
	_ "github.com/rs/zerolog"
	"github.com/suchy1105/GUIcontroler/api"
	//"./api"
	//"github.com/suchy1105/GUIcontroler/config"
	"net/http"

	//	"github.com/suchy1105/GUIcontroler/config"
)
var Sstate *api.GuiState
func main() {
	defer run()
}
func run() {
//	var err error
//	var conf config.Configuration
//	conf.GetConf()
//

	router := chi.NewRouter()


	router.Route("/api/v1", func(router chi.Router) {

		router.Post("/message/", api.PostMessage)


		router.Get("/message/", api.GetMessages)
		router.Get("/checkhealth", api.CheckHealth)
		router.NotFound(api.NotFound)
	})
	http.ListenAndServe(":1600", router)

}