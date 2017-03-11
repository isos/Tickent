package main

import (
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/skip2/go-qrcode"
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
	Tienda string
	Logo   string
	Model  string
	Fecha  string
	Color  int
	Total  float32
	Items  []CompraItem
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola")
}

func ClientConn(w http.ResponseWriter, r *http.Request) {
	var iduser string

	r.ParseForm()

	if r.Method == "GET" {
		http.Error(w, "Error: Use POST Method. GET Method is not secure", 406)
		return
	}

	iduser = r.FormValue("userid")
	if len(iduser) == 0 {
		http.Error(w, "Forbidden", 400)
		return
	}

	db, err := sql.Open("mysql", "ticket:X2L1aLOJ@/tickent")
	if err != nil {
		fmt.Println("Error: Failed to open SQL connection")
		db.Close()
		return
	}

	if db.Ping() != nil {
		fmt.Println("Error: Connecting to database")
		fmt.Fprintf(w, "Error: Connecting to MySQL Database", 500)
		return
	}

	query := "SELECT idnfc, model FROM nfc WHERE idu = ?"

	var nfcusers []NFCInfo

	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, "Error: Failed executing query", 503)
		db.Close()
		return
	}

	var p int = 0
	for rows.Next() {
		err := rows.Scan(&nfcusers[p].IdNFC, &nfcusers[p].Model)
		if err != nil {
			http.Error(w, "Error: Failed getting data", 500)
			rows.Close()
			db.Close()
			return
		}

		p++
	}

	rows.Close()

	var client []ClientJson

	var i int = 0
	for i < len(nfcusers) {
		query = "SELECT cs.idc, t.empresa, t.color, t.logo, cs.total, cs.fecha FROM tienda t, compras cs WHERE cs.idnfc = ?"
		rows, err := db.Query(query, nfcusers[i].IdNFC)
		if err != nil {
			fmt.Fprintf(w, "Error: Executing SQL Query", 500)
			db.Close()
			return
		}

		var idc int = 0
		p = 0
		for rows.Next() {
			rows.Scan(&idc, &client[p].Tienda, &client[p].Color, &client[p].Logo, &client[p].Total, &client[p].Fecha)

			cmQuery := "SELECT articulo, cantidad, precio FROM compra WHERE idc = ?"

			cRows, err := db.Query(cmQuery, idc)
			if err != nil {
				fmt.Fprintf(w, "Error: Executing query", 500)
				rows.Close()
				db.Close()
				return
			}

			var scItem CompraItem
			for c := 0; cRows.Next(); c++ {
				cRows.Scan(&scItem.Articulo, &scItem.Cantidad, &scItem.Precio)
				client[p].Items = append(client[p].Items, scItem)
			}

			p++
			cRows.Close()
		}

		i++
		rows.Close()
	}

	db.Close()

	resp, _ := json.Marshal(client)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PUT")
	w.Header().Set("Access-Control-Max-Age", "1000")
	w.Header().Set("Access-Control-Allow-Headers", `"x-requested-with, Content-Type, origin, authorization, accept, client-security-token"`)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(resp))
}

func TpuConnect(w http.ResponseWriter, r *http.Request) {
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

	var qr bool = false

	if tpuJson.IdNfc == "borja" {
		qr = true
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
	_, err = stmt.Exec(tpuJson.IdNfc, tpuJson.IdTienda)
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
		_, err = stmt.Exec(&lastId, &tpuJson.Items[i].Articulo, &tpuJson.Items[i].Precio, &tpuJson.Items[i].Cantidad)
		if err != nil {
			fmt.Println(err)
			fmt.Fprintf(w, "Error: Internal error", 500)
			return
		}
	}
	if qr {
		fmt.Fprintf(w, string(lastId))
	} else {
		fmt.Fprintf(w, "OK")
	}
}

func QrIf(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK")
}

func QrImg(w http.ResponseWriter, r *http.Request) {
	content := r.FormValue("idc")
	if content != "" {
		png, _ := qrcode.Encode(content, qrcode.Medium, 256)
		png_b64 := base64.StdEncoding.EncodeToString(png)
		fmt.Fprintf(w, `<html><body><img alt="QR Code: pepe" src="data:image/png;base64,%s"></body></html>`, png_b64)
	}
}

func main() {
	fmt.Println("Server Started")

	http.HandleFunc("/", Index)
	http.HandleFunc("/client", ClientConn)
	http.HandleFunc("/tpuconnect", TpuConnect)
	http.HandleFunc("/qrif", QrIf)
	http.HandleFunc("/qr", QrImg)

	http.ListenAndServe(":8080", nil)
}
