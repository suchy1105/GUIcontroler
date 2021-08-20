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
	AlsaVolumeM string `json:"alsa_volumem"`
	AlsaVolume1 string `json:"alsa_volume1"`
	AlsaVolume2 string `json:"alsa_volume2"`
	AlsaVolume3 string `json:"alsa_volume3"`
	AlsaVolume4 string `json:"alsa_volume4"`
	NumberOfCards string `json:"number_of_cards"`
	MuteMaster bool `json:"mute_master"`
	MuteCH1 bool`json:"mute_ch_1"`
	MuteCH2 bool`json:"mute_ch_2"`
	MuteCH3 bool`json:"mute_ch_3"`
	MuteCH4 bool`json:"mute_ch_4"`
	PlayVoice bool `json:"play_voice"`

}
func NewGuiState()*GuiState{
	 g:= GuiState{
		Ip:        "000.000.000.000",
		Mac:       "AAAA:BBBB:CCCC:DDDD",
		ConnState: true,
		PlayStipa: false,
		AlsaVolumeM: "80",
		AlsaVolume1: "70",
		AlsaVolume2: "60",
		AlsaVolume3: "50",
		AlsaVolume4: "40",
		NumberOfCards: "4",
		MuteMaster: false,
		MuteCH1: false,
		MuteCH2: false,
		MuteCH3: false,
		MuteCH4: false,
		PlayVoice: false,
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

		response, err := json.Marshal(s)
		if err != nil {
			log.Println("Can't marshal data: ", err)
		}

		w.Write(response)

		w.WriteHeader(http.StatusOK)
	}
}
//PostMessage posts
func postMessageHandler(s *GuiState) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		guistate := NewGuiState()
		json.NewDecoder(r.Body).Decode(guistate)

		s.Ip = guistate.Ip
		s.Mac = guistate.Mac
		s.ConnState = guistate.ConnState
		s.PlayStipa = guistate.PlayStipa
		s.AlsaVolumeM = guistate.AlsaVolumeM
		w.WriteHeader(http.StatusCreated)
		fmt.Println(s)
	}
}
//NotFound 404
func NotFound(w http.ResponseWriter, r *http.Request) {
fmt.Println("404")
	w.WriteHeader(http.StatusNotFound)
}

func  CheckHealth(w http.ResponseWriter, r *http.Request) {

		w.WriteHeader(http.StatusOK)

}
