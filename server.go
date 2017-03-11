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

func ClientConn(w http.ResponseWriter, r *http.Request) {
  r.ParseForm ()

  if r.Method == "GET" {
    http.Error ( w, "Error: Use POST Method. GET Method is not secure", 406 )
    return
  }

  iduser := r.PostFormValue ( "iduser" );

	db, err := sql.Open("mysql", "ticket:X2L1aLOJ@/tickent")
	if err != nil {
		fmt.Println ("Error: Failed to open SQL connection")
		db.Close  ()
		return
	}

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
  var resp ClientJson

  query = "SELECT t.empresa, t.color, t.logo, cs.total, cs.fecha FROM tienda t, compras cs WHERE cs.idnfc = ?"

  for p := 0; p < rows.Next (); p++ {
    rows, err = db.Query ( query, nfcusers[p] );

    db.Prepare ("SELECT COUNT(idl) FROM compra WHERE idc = ?")
    db.Exec ( idc )

    for i := 0; 

  resp.Items = make ( []CompraItem, 

	json.NewEncoder(w).Encode(resp)
}

func main() {
  fmt.Println ("Server Started")

	http.HandleFunc("/client", ClientConn)

	http.ListenAndServe(":8080", nil)
}
