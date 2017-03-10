package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
  "time"
  "database/sql"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

type CompraItem struct {
	Articulo string
	Precio   float32
	Cantidad int
}

type TicketJson struct {
	IdNfc    string
	IdTienda string
	Items    []CompraItem
}

type ClientJson struct {
	Tienda string
	Logo   string
	Color  int
	Total  float32
	Items  []CompraItem
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola")
}

func TpuConnect(w http.ResponseWriter, r *http.Request) {
	//get Json from TPU
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

	//open SQL connection
	db, err := sql.Open("mysql", "ticket:X2L1aLOJ@/tickent")
	if err != nil {
		fmt.Println("Error: Failed to open SQL connection")
		db.Close()
		return
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("Error: Failed to ping SQL db")
	}

	//prepared statement from struct values
	//calculate total first?
	stmt, err := db.Prepare("INSERT INTO compras(idnfc, idt) VALUES (?, ?)")
	if err != nil {
		fmt.Println(err)
	}
	_, err := stmt.Exec(tpuJson.idNfc, tpuJson.idTienda)
	if err != nil {
		fmt.Println(err)
	}

	//get last idc -> lastId

	var lastId int

	for i := 0; i < len(tpuJson.items); i++ { //insert all items in json
		stmt, err := db.Prepare("INSERT INTO compra(idc, articulo, precio, cantidad) VALUES (?, ?, ?, ?)")
		if err != nil {
			fmt.Println(err)
		}
		_, err := stmt.Exec(lastId, tpuJson.items[i].articulo, tpuJson.items[i].precio, tpuJson.items[i].cantidad)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func ClientConn(w http.ResponseWriter, r *http.Request) {
  var buf bytes.Buffer

  logf := log.New (&buf, "logger: ", log.Lshortfile )

  r.ParseForm ()

  if r.Method == "GET" {
    http.Error ( w, "Error: Use POST Method. GET Method is not secure", 406 )
    return
  }

	db, err := sql.Open("mysql", "ticket:X2L1aLOJ@/tickent")
	if err != nil {
		fmt.Println ("Error: Failed to open SQL connection")
		db.Close  ()
		return
	}


	var resp TicketJson
  /*
	tienda string
	logo   string
	color  int
	total  float32
	items  []CompraItem
  */
  for p := 0; rows.Next (); p++ {
    ... 
  }

	json.NewEncoder(w).Encode(resp)
}

func main() {
  fmt.Println ("Server Started")

	http.HandleFunc("/", Index)
	http.HandleFunc("/tpuconnect", TpuConnect)
	http.HandleFunc("/client", ClientConn)

	http.ListenAndServe(":8080", nil)
}
