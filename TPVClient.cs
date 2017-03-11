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
using System.Diagnostics;


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
    
    public static void StartQR()
    {
        string idc = GetQR();
        OpenQR(idc);
    }
    public static string GetQR()
    {
        WebClient webClient = new WebClient();
        string URL = "http://tickent.tk:8080/tpuconnect";
        NameValueCollection formData = new NameValueCollection();
        string id = "borja";
        string idc = "";
        
        formData["json"] =
            "{\"IdNfc\":\"" +
             id + 
            "\",\"IdTienda\":\"1234\",\"Items\":[" + 
            GetProductoJson()
            + "]}";
        
        byte[] arrByte = webClient.UploadValues(URL, "POST", formData);
        foreach (byte b in arrByte)
        {
            idc += Convert.ToChar(b);
        }
        
        return idc;
    }
    public static void OpenQR(string idc)
    {
        /*string filename
            = string.Format(@"{0}\{1}",
            System.IO.Path.GetTempPath(),
            "testhtm.htm");
            
        File.WriteAllText(filename, "<h1>Hello</h1>");
        Process.Start(filename);*/
        
        //System.Diagnostics.Process.Start("data:image/png;base64," + QR);
        //Process.Start("data:image/png;base64," + QR);
        Process.Start("http://Tickent.tk:8080/qr?idc=" + idc);
    }
    
    public static void AnadirProductos()
    {
        string inpt = "";
        while (inpt != "0")
        {
            articulo tempArt;
            Console.Write("Introduce producto: ");
            inpt = Console.ReadLine();
            
            if (inpt == "QR" || inpt == "qr" || inpt == "Qr" || inpt == "qR")
            {
                StartQR();
            }
            else if (inpt != "0")
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
        if (prodJson.Length < 4)
        {
            return "NoProds";
        }
        else
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
