package api

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
	"strings"
)

// Message struct
type Message struct {
	Volume       string `json:"volume"`
	Ip      string `json:"ip"`
	Content     string `json:"mac"`
	MagicNumber bool  `json:"connectionState"`
	PlayNTIStipa bool `json:"playstipa"`
}
func ControlAPI() func(router chi.Router) {
	return func(router chi.Router) {
			router.Post("v1/message/", PostData)

			router.Get("v1/message/", GetData)
			router.Get("v1/health", CheckHealth)
	}
}
func CheckHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
//Get API
func GetData(w http.ResponseWriter, r *http.Request) {

	reqEmail := strings.Split(r.RequestURI, "/api/message/")
	fmt.Println(reqEmail[1])

}
//Post API
func PostData(w http.ResponseWriter, r *http.Request) {
	var message Message
	json.NewDecoder(r.Body).Decode(&message)

//TODO
	w.Write([]byte(`{"message": "post called"}`))

	w.WriteHeader(http.StatusCreated)
}

//NotFound 404
func NotFound(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusNotFound)


}