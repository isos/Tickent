package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"time"
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

type NFCInfo struct {
	IdNFC string
	Model string
}

type ClientJson struct {
	Tienda string
	Logo   string
	Model  string
	Color  int
	Total  float32
	Items  []CompraItem
}

<<<<<<< HEAD
=======
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

	err := json.Unmarshal([]byte(r.FormValue("json")), &tpuJson)
	if err != nil {
		fmt.Println(err)
		return
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

	for i := 0; i < len(tpuJson.Items); i++ { //insert all Items in json
		stmt, err := db.Prepare("INSERT INTO compra(idc, articulo, precio, cantidad) VALUES (?, ?, ?, ?)")
		if err != nil {
			fmt.Println(err)
		}
		_, err := stmt.Exec(lastId, tpuJson.Items[i].articulo, tpuJson.Items[i].precio, tpuJson.Items[i].cantidad)
		if err != nil {
			fmt.Println(err)
		}
	}
	http.Error(w, "OK", 200)
}

>>>>>>> d3fee41a11ada9aa509cfeffbf6d9fbfc27ebfce
func ClientConn(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	if r.Method == "GET" {
		http.Error(w, "Error: Use POST Method. GET Method is not secure", 406)
		return
	}

	iduser := r.PostFormValue("userid")
	if userid == "" {
		http.Error(w, "Forbidden", 400)
		return
	}

	db, err := sql.Open("mysql", "ticket:X2L1aLOJ@/tickent")
	if err != nil {
		fmt.Println("Error: Failed to open SQL connection")
		db.Close()
		return
	}

<<<<<<< HEAD
  query := "SELECT idnfc, model FROM nfc WHERE idu = ?"

  rows, err = db.Query ( query )
  if err != nil {
    http.Error ( w, "Error: Failed executing query", 503 )
    rows.Close ()
    db.Close ()
    return
  }

  var nfcusers []NFCInfo

  for p := 0; rows.Next (); p++ {
    if rows.Scan ( &nfcusers[p].IdNFC, &nfcusers[p].Model ) != nil {
      http.Error ( w, "Error: Failed getting data", 500 )
      rows.Close ()
      db.Close ()
      return
    }
  }
  /*
type TicketJson struct {
	IdNfc    string
	IdTienda string
	Items    []CompraItem
}
=======
	rows, err = db.Query("SELECT idnfc,model FROM nfc WHERE idu = ?", iduser)
	if err != nil {
		http.Error(w, "Error: Failed executing query", 503)
		rows.Close()
		db.Close()
		return
	}
>>>>>>> d3fee41a11ada9aa509cfeffbf6d9fbfc27ebfce

	var nfcusers []NFCInfo

	for p := 0; rows.Next(); p++ {
		if rows.Scan(&nfcusers[p].IdNFC, &nfcusers[p].Model) != nil {
			http.Error(w, "Error: Failed getting data", 500)
			rows.Close()
			db.Close()
			return
		}
	}
	/*
		type TicketJson struct {
			IdNfc    string
			IdTienda string
			Items    []CompraItem
		}

		type NFCInfo struct {
		  IdNFC string
		  Model string
		}

		type ClientJson struct {
			Tienda  string
			Logo    string
		  Model   string
			Color   int
			Total   float32
			Items   []CompraItem
		}
	*/

	//return json to client
	var client []ClientJson
	resp, _ := json.Marshal(client)
	fmt.Fprintf(w, string(resp))

<<<<<<< HEAD
  query = "SELECT t.empresa, t.color, t.logo, cs.total, cs.fecha FROM tienda t, compras cs WHERE cs.idnfc = ?"

  for p := 0; p < rows.Next (); p++ {
    rows, err = db.Query ( query, nfcusers[p] );

    db.Prepare ("SELECT COUNT(idl) FROM compra WHERE idc = ?")
    db.Exec ( idc )

    for i := 0; 

  resp.Items = make ( []CompraItem, 

	json.NewEncoder(w).Encode(resp)
=======
>>>>>>> d3fee41a11ada9aa509cfeffbf6d9fbfc27ebfce
}

func main() {
	fmt.Println("Server Started")

	http.HandleFunc("/client", ClientConn)

	http.ListenAndServe(":8080", nil)
}
