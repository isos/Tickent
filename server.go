package main

import (
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
	resp := r.FormValue("tpu_quKW3zx")
	err = json.Unmarshal(resp, &tpuJson)
}

func main() {
	http.HandleFunc("/", Index)
	http.HandleFunc("/tpuconnect", TpuConnect)
	http.ListenAndServe(":8080", nil)
}
