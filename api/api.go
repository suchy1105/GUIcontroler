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

//GetMessages API  messages get provider
func GetMessages(w http.ResponseWriter, r *http.Request) {


/*	 guistate:=GuiState{
		 Ip:        "0.0.0.0",
		 Mac:       "ABCD:EGGH:1234:5678",
		 ConnState: true,
		 PlayStipa: false,
	 }*/

	fmt.Println("List allrequri: ", r.RequestURI)





			response, err := json.Marshal(guistate)
			if err != nil {
				log.Println("Can't marshal data: ", err)
			}
			fmt.Println("odpowiedz: ",response)
			w.Write(response)





	fmt.Println("done")

}
//PostMessage posts
func PostMessage(w http.ResponseWriter, r *http.Request) {
	var guistate GuiState
	json.NewDecoder(r.Body).Decode(&guistate)

	fmt.Println("mess: ",guistate," :mess")


	w.Write([]byte(`{"message": "post called"}`))

	//w.WriteHeader(http.StatusCreated)
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