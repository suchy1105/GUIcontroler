package main//GUIsocket

import (
	"fmt"
	//"github.com/go-chi/chi"
	"net/http"
	"github.com/suchy1105/GUIcontroler/api"
//	"github.com/suchy1105/GUIcontroler/config"
)

func main() {
	defer run()
}
func run() {
	fmt.Println("karil")
	//router:=




	http.ListenAndServe(":8080", api.ControlAPI() )

}