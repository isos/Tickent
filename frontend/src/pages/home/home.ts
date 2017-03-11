import { Component, OnInit } from '@angular/core';
import { NavController } from 'ionic-angular';
import { HttpData } from '../../providers/http-data';

@Component({
  selector: 'page-home',
  templateUrl: 'home.html'
})

export class HomePage implements OnInit {
  tickets : Tickent[];
  //tickets : JSON;
  aham : any;
  texto : String;
  items : any;
  todosTickets : any[] = [];
  constructor(public navCtrl: NavController, public serviceData: HttpData) {    
      
  }


  

  ngOnInit()
  {
    this.serviceData.load().then(data => {
        this.items = data;
        console.log(this.items);
        
        //this.aham = JSON.stringify(this.items);
        /*for (var i = 0; i < this.items.length; i++) {
            var element = this.items[i];
            let tmp : Tickent;
            //tmp.logo = element.Logo;
            //tmp.color = element.Color;
            //tmp.total = element.Total;
            //tmp.fecha = element.Fecha;
            console.log(element.Color);

            //this.tickets.push(tmp);

        }*/
    });
    //this.serviceData.load();
      //this.tickets = this.serviceData.getData();
        //var contenido = JSON.parse( 'tienda":"Aldi","logo":"https://www.aldi.es/wp-content/uploads/2016/10/logo_aldi.png","color":"#232E7D","total":22.3,"fecha": "3 de febrero", "productos":[{"nombreProducto":"pan","precioTotal":0.69,"Cantidad":20},{"nombreProducto":"condones","precioTotal":19.65,"Cantidad":69}]}');
      //console.log( contenido );


      /*this.tickets = [
        {
            "tienda": "Carrefour Market",
            "logo": "http://www.animationsvin.fr/sites/animvin/images/vin/logos/awi_carrefour.gif",
            "color": "#00529C",
            "total": 1.45,
            "productos" : [{"nombreProducto": "Galletas","precioTotal":1.02,"cantidad": 1},
                            {"nombreProducto": "Agua","precioTotal":0.43,"cantidad": 2} 
            ],
            "fecha" : "Justo ahora"
        },
        {
            "tienda": "Lidl",
            "logo": "https://upload.wikimedia.org/wikipedia/commons/thumb/9/91/Lidl-Logo.svg/2000px-Lidl-Logo.svg.png",
            "color": "#ffcc00",
            "total": 10.02,
            "productos" : [{"nombreProducto": "Galletas","precioTotal":1.02,"cantidad": 1},
                            {"nombreProducto": "Agua","precioTotal":0.43,"cantidad": 1} 
            ],
            "fecha" : "Hace 3 dias"
        },
        {
            "tienda": "Mercadona",
            "logo": "https://quienmanda.s3.amazonaws.com/uploads/avatar/9277/logo.png",
            "color": "#00754D",
            "total": 10.27,
            "productos" : [
                {"nombreProducto": "Pack x6 Leche","precioTotal":6.82, "cantidad": 1}, 
                {"nombreProducto": "Chocolate","precioTotal":0.7,"cantidad": 1},
                {"nombreProducto": "Baguette","precioTotal":0.3, "cantidad": 1},
                {"nombreProducto": "Chorizo","precioTotal":1.1,"cantidad": 2},
                {"nombreProducto": "Agua","precioTotal":0.25, "cantidad": 1}
            ],
            "fecha": "Ayer"
        },
        
        {
            "tienda":"Aldi",
            "logo":"https://www.aldi.es/wp-content/uploads/2016/10/logo_aldi.png",
            "color":"#232E7D",
            "total":22.3,
            "fecha": "3 de febrero",
            "productos":[
                {"nombreProducto":"pan","precioTotal":0.69, "cantidad": 3},
                {"nombreProducto":"Ron Almirante","precioTotal":19.65, "cantidad": 2}
            ]  
        }
    ];
      //this.tickets.push(contenido); */
  }  
}





export class Tickent {

    tienda: String;
    logo: String;
    color: String;
    total: number;
    productos: linPed[];
    fecha: string;
    constructor(tienda: string,logo: string,color: string,total: number, fecha: string) {
        this.tienda = tienda;
        this.logo = logo;
        this.color = color;
        this.total = total;
        this.fecha = fecha;
    }
 }

export class linPed {
    nombreProducto: String;
    precioTotal: number;
    cantidad: number;
}

