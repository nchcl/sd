package chat

import (
    "os"
    "log"
    "encoding/csv"
    "fmt"
    "time"
    "strconv"

	"golang.org/x/net/context"
)

type Server struct {
}

var seguimiento = 130031
var ide = 1
var retail [][]string
var prioritario [][]string
var normal [][]string
var registro [][]string

func (s *Server) OrderRetail(ctx context.Context, in *Retail) (*Confirmation, error) {
	//log.Printf("Id: %s", in.Id)
    //log.Printf("Producto: %s", in.Producto)
    //log.Printf("Valor: %s", in.Valor)
    //log.Printf("Tienda: %s", in.Tienda)
    //log.Printf("Destino: %s", in.Destino)
    //log.Printf("")
    
    var data [][]string
    t := time.Now()
    data = append(data, []string{t.Format("2006-01-02 15:04:05"),fmt.Sprintf("%d",      ide),"retail",in.Producto,in.Valor,in.Tienda,in.Destino,"0"})    
    appendCSV(data)
    
    retail = append(retail, []string{fmt.Sprintf("%d",ide),"0","retail",in.Valor,"0","En bodega",in.Tienda,in.Destino})

    ide += 1    
	return &Confirmation{Body: "Orden confirmada"}, nil
}
 
func (s *Server) OrderPyme(ctx context.Context, in *Pyme) (*Confirmation, error) {
	//log.Printf("Id: %s", in.Id)
    //log.Printf("Producto: %s", in.Producto)
    //log.Printf("Valor: %s", in.Valor)
    //log.Printf("Tienda: %s", in.Tienda)
    //log.Printf("Destino: %s", in.Destino)
    //log.Printf("Prioritario: %s", in.Prioritario)
    //log.Printf("")
    
    prio, err := strconv.Atoi(in.Prioritario)
    if err != nil {           
        log.Fatal(err)                                                                                                                                                        
    }    
    var tipo string
    if prio > 0 {
        tipo = "prioritario"
        prioritario = append(prioritario, []string{fmt.Sprintf("%d",ide),fmt.Sprintf("%d", seguimiento),tipo,in.Valor,"0","En bodega",in.Tienda,in.Destino})
        //fmt.Println(prioritario[ide-1])
    } else {
        tipo = "normal"
        normal = append(normal, []string{fmt.Sprintf("%d",ide),fmt.Sprintf("%d", seguimiento),tipo,in.Valor,"0","En bodega",in.Tienda,in.Destino})
        //fmt.Println(normal[ide-1])
    }
    
    var data [][]string
    t := time.Now()
    data = append(data, []string{t.Format("2006-01-02 15:04:05"),fmt.Sprintf("%d", ide),tipo,in.Producto,in.Valor,in.Tienda,in.Destino,fmt.Sprintf("%d", seguimiento)})    
    appendCSV(data)
    
    ide += 1
    seguimiento += 1
    
	return &Confirmation{Body: fmt.Sprintf("%d", seguimiento)}, nil
}

func (s *Server) Seguimiento(ctx context.Context, in *Confirmation) (*Confirmation, error) {
    //var codigo = in.Body
    
    
    
	return &Confirmation{Body: "EN CONSTRUCCION"}, nil
}

func (s *Server) Camion(ctx context.Context, in *Tipo) (*Paquete, error) {
    var tipo_camion = in.Tipo
    var paquete_camion []string
    
    if tipo_camion == 1 {
        if len(retail) == 0 {
            return &Paquete{}, nil
        }
        
        paquete_camion, retail = retail[0], retail[1:]

    } else {
        if len(normal) == 0 && len(prioritario) == 0{
            return &Paquete{}, nil
        } else if len(prioritario) == 0 {
            paquete_camion, normal = normal[0], normal[1:]
        } else {
            paquete_camion, prioritario = prioritario[0], prioritario[1:]
        }
    }
    
    paquete_camion[5] = "En camino"
    registro = append(registro, paquete_camion)
    fmt.Println(registro)    
    
	return &Paquete{Idpaquete: paquete_camion[0],Seguimiento: paquete_camion[1],Tipo: paquete_camion[2],Valor: paquete_camion[3],Intentos: paquete_camion[4],Estado: paquete_camion[5], Origen: paquete_camion[6], Destino: paquete_camion[7]}, nil
}

func appendCSV(data [][]string) {
    var path = "result.csv"                                                                                                                                                   
    f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, os.ModeAppend)     
    
    if err != nil {           
        fmt.Println("hello")
        log.Fatal(err)                                                                                                                                                        
    }  
    
    defer f.Close()                                                                                                                                                           
                                                                                                                                                                                                                                                                                                                                                                                                                        
    w := csv.NewWriter(f)                                                                                                                                                     
    w.WriteAll(data)                                                                                                                                                          
                                                                                                                                                                                
    if err := w.Error(); err != nil {                                                                                                                                         
        log.Fatal(err)                                                                                                                                                        
    }
}
