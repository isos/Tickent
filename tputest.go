package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"net/http"
	"net/url"
)

type CompraItem struct {
	Articulo string
	Precio   float32
	Cantidad int
}
type TicketJson struct {
	IDNFC    string
	IDTienda string
	Items    []CompraItem
}
type ClientJson struct {
	Tienda string
	Logo   string
	Color  int
	Total  float32
	Items  CompraItem
}

func main() {
	var tpu TicketJson

	tpu.Items = make([]CompraItem, 2)
	//tpu.IDNFC = `#79c"AXE?bc%>@rS8{G??z6"QJMk&>/By8:u5(Sdpz<LQ%a5LV2w;x#K/4>tMwm8:MnUY=E,[.WVh"pjbSQrL&k@_EGe6y4&8tnNdh&[+U(YFTp?.bHnKGj.gEx#!7'r`
	tpu.IDNFC = "pepe"
	tpu.IDTienda = "1234"
	tpu.Items[0].Articulo = "pan"
	tpu.Items[0].Precio = 0.69
	tpu.Items[0].Cantidad = 20
	tpu.Items[1].Articulo = "condones"
	tpu.Items[1].Precio = 19.65
	tpu.Items[1].Cantidad = 69

	spew.Dump(tpu)

	resp, _ := json.Marshal(tpu)
	spew.Dump(resp)
	fmt.Println(string(resp))
	w := new(bytes.Buffer)
	json.NewEncoder(w).Encode(resp)

	form := url.Values{}
	form.Add("json", string(resp))
	http.PostForm("http://tickent.tk:8080/tpuconnect", form)
}
