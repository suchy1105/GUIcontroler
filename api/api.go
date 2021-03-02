package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"fmt"
	"time"
)

type Message struct {
	Email       string `json:"email"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	MagicNumber int32  `json:"magic_number"`
}

//GetMessages API  messages get provider
func GetMessages(w http.ResponseWriter, r *http.Request) {

	reqEmail := strings.Split(r.RequestURI, "/apiv1/messages/")
	fmt.Println(reqEmail[1])

	var message Message

	fmt.Println("List allrequri: ", r.RequestURI)





			response, err := json.Marshal(message)
			if err != nil {
				log.Println("Can't marshal data: ", err)
			}
			fmt.Println("odpowiedz: ",response)
			w.Write(response)





	fmt.Println("done")

}
//PostMessage posts
func PostMessage(w http.ResponseWriter, r *http.Request) {
	var message Message
	json.NewDecoder(r.Body).Decode(&message)

	zerotime := time.Now()
	fmt.Println(time.Now().Sub(zerotime))
	fmt.Println("mess: ",message," :mess")
	if err := Session.Query(`INSERT INTO mess (email, title, content, magic_number, timestamp) VALUES (?, ?, ?, ?, ?)`,
		&message.Email, &message.Title, &message.Content, &message.MagicNumber, zerotime.Format(time.UnixDate)).Exec(); err != nil {
		log.Println("Insertion error: ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Write([]byte(`{"message": "post called"}`))
	fmt.Println("messsend:  ", message.MagicNumber)
	w.WriteHeader(http.StatusCreated)
}

//NotFound 404
func NotFound(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusNotFound)


}


func CheckHealth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.WriteHeader(http.StatusOK)
	}
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