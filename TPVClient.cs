using System;
using System.Threading;
using System.IO.Ports;
using System.IO;
using System.Collections.Generic;
using System.ComponentModel;
using System.Linq;
using System.Text;
using System.Collections.Specialized;
using System.Net;



public class ArduinoNFC
{
    static byte jump = 0;
    
    public struct articulo
    {
        public string nombre;
        public int cantidad;
        public double precio;
    }
    
    public static List<articulo> articulos = new List<articulo>();
    
    public static void AnadirProductos()
    {
        string inpt = "";
        while (inpt != "0")
        {
            articulo tempArt;
            Console.Write("Introduce producto: ");
            inpt = Console.ReadLine();
            if (inpt != "0")
            {
                tempArt.nombre = inpt;
                
                Console.Write("Introduce cantidad: ");
                inpt = Console.ReadLine();
                if (inpt != "0")
                {
                    tempArt.cantidad = Convert.ToInt32(inpt);
                    
                    Console.Write("Introduce precio: ");
                    inpt = Console.ReadLine();
                    
                    if (inpt != "0")
                    {
                        tempArt.precio = Convert.ToDouble(inpt);
                        
                        articulos.Add(tempArt);
                    }
                    
                    
                }
            }
            
        }
        
    }
    
    public static string GetProductoJson()
    {
        string prodJson = "";
        
        foreach (articulo a in articulos)
        {
            prodJson += "{\"Articulo\":\"" + a.nombre +
             "\",\"Precio\":" + a.precio.ToString() +
              ",\"Cantidad\":" + a.cantidad.ToString() + "},";
        }
        
        return prodJson.Substring(0, prodJson.Length -1);
    }
    
    public static string GetID(SerialPort sp)
    {
        string hexString = sp.ReadExisting();
            
        if (jump == 0)
        {
            jump++;
            return "";
        }
        else
        {
            jump = 0;
            return ("AAAAAAAAAAAAAAAAAAAAAAAA" + hexString +
                ((hexString.Length > 7) ? "":"0"));
        }
    }
    
    public static void SendJson(string id)
    {
        string URL = "http://tickent.tk:8080/tpuconnect";
        WebClient webClient = new WebClient();

        NameValueCollection formData = new NameValueCollection();
        formData["json"] =
            "{\"IdNfc\":\"" +
             id + 
            "\",\"IdTienda\":\"1234\",\"Items\":[" + 
            GetProductoJson()
            + "]}";
        webClient.UploadValues(URL, "POST", formData);
        webClient.Dispose();
    }
    
    public static void Main()
    {
        bool quit = false;
        SerialPort sp = new SerialPort();
        sp.PortName = "/dev/ttyACM0";
        sp.Open();
        string id = "";
        
        while (!quit)
        {
            AnadirProductos();
            
            while (id.Length != 32)
            {
                id = GetID(sp);
                Thread.Sleep(200);
            }
            SendJson(id);
            id = "";
            
            articulos.Clear();
        }
        
        sp.Close();
    }
}
