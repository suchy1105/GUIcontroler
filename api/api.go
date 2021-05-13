package api

import (
	"encoding/json"
	"fmt"
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
//GetMessages API  messages get provider
func (g *GuiState)GetMessages(w http.ResponseWriter, r *http.Request) {

	fmt.Println("List allrequri: ", r.RequestURI)
	fmt.Println("GG: ", g)

			response, err := json.Marshal(g)
			if err != nil {
				log.Println("Can't marshal data: ", err)
			}
			fmt.Println("odpowiedz: ",response)
			w.Write(response)

w.WriteHeader(http.StatusOK)
}
//PostMessage posts
func (g *GuiState)PostMessage(w http.ResponseWriter, r *http.Request) {
	guistate:=NewGuiState()
	json.NewDecoder(r.Body).Decode(guistate)

	g.Ip=guistate.Ip
	g.Mac=guistate.Mac
	g.ConnState=guistate.ConnState
	g.PlayStipa=guistate.PlayStipa


	fmt.Println(g)
	fmt.Println(g.Ip)
	fmt.Println(guistate.Ip)

	w.WriteHeader(http.StatusCreated)
}

//NotFound 404
func NotFound(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusNotFound)


}

func  CheckHealth(w http.ResponseWriter, r *http.Request) {

		w.WriteHeader(http.StatusOK)

}
//Get API
/*
func getData() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		reqMessage := strings.Split(r.RequestURI, "/v1/message/")
		fmt.Println(reqMessage[1])
		w.WriteHeader(http.StatusTeapot)
	}

}
//Post API
func postData() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var s State
		json.NewDecoder(r.Body).Decode(&s)
		fmt.Println("messageold:", &s)
		fmt.Println("messageold:", s)
		state=&s
		fmt.Println("message:",s)
		//TODO
		w.Write([]byte(`{"message": "post called"}`))

		w.WriteHeader(http.StatusCreated)
	}

}*/