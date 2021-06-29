package api

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"log"
	"net/http"
)

type GuiState struct {
	Ip       string `json:"ip"`
	Mac       string `json:"mac"`
	ConnState bool `json:"cstate"`
	PlayStipa bool  `json:"stipa"`
}
func NewGuiState()*GuiState{
	 g:= GuiState{
		Ip:        "xx",
		Mac:       "xx",
		ConnState: false,
		PlayStipa: false,
	}
	return &g
}

//FrontendAPI xxxxd
func FrontendAPI(s *GuiState) func(router chi.Router) {
	return func(router chi.Router) {
		router.Get("/get", getMessagesHandler(s))
		router.Post("/post", postMessageHandler(s))
	}
}

//GetMessages API  messages get provider
func getMessagesHandler(s *GuiState) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("List allrequri: ", r.RequestURI)
		fmt.Println("GG: ", s)

		response, err := json.Marshal(s)
		if err != nil {
			log.Println("Can't marshal data: ", err)
		}
		fmt.Println("odpowiedz: ", response)
		w.Write(response)

		w.WriteHeader(http.StatusOK)
	}
}
//PostMessage posts
func postMessageHandler(s *GuiState) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("pst")
		guistate := NewGuiState()
		json.NewDecoder(r.Body).Decode(guistate)

		s.Ip = guistate.Ip
		s.Mac = guistate.Mac
		s.ConnState = guistate.ConnState
		s.PlayStipa = guistate.PlayStipa

		fmt.Println(s)
		fmt.Println(s.Ip)
		fmt.Println(guistate.Ip)

		w.WriteHeader(http.StatusCreated)
	}
}
//NotFound 404
func NotFound(w http.ResponseWriter, r *http.Request) {
fmt.Println("404x")
	w.WriteHeader(http.StatusNotFound)
}

func  CheckHealth(w http.ResponseWriter, r *http.Request) {

		w.WriteHeader(http.StatusOK)

}
