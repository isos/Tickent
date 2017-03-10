package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type TicketJson struct {
	IdU   string
	items []string
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola")
}

func TpuConnect(w http.ResponseWriter, r *http.Request) {
	var tpuJson TicketJson

	if r.Body == nil {
		http.Error(w, "Forbidden", 400)
		return
	}

	for key, _ := range r.Form {
		err := json.Unmarshal([]byte(key), &tpuJson)
		if err != nil {
			fmt.Println("Error: ", err)
			http.Error(w, "Forbidden", 400)
			return
		}
	}

	//put into db
}

func ClientConn(w http.ResponseWriter, r *http.Request) {
	userid := r.FormValue("userid")

	//get from db put in "resp" struct

	json.NewEncoder(w).Encode(resp)

}

func main() {
	fmt.Println("Server Started")
	http.HandleFunc("/", Index)
	http.HandleFunc("/tpuconnect", TpuConnect)
	http.HandleFunc("/client", ClientConn)
	http.ListenAndServe(":8080", nil)
}
