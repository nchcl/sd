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

func (s *Server) OrderRetail(ctx context.Context, in *Retail) (*Confirmation, error) {
	log.Printf("Id: %s", in.Id)
    log.Printf("Producto: %s", in.Producto)
    log.Printf("Valor: %s", in.Valor)
    log.Printf("Tienda: %s", in.Tienda)
    log.Printf("Destino: %s", in.Destino)
    log.Printf("")
    
    var data [][]string
    t := time.Now()
    data = append(data, []string{t.Format("2006-01-02 15:04:05"),fmt.Sprintf("%d",      ide),"retail",in.Producto,in.Valor,in.Tienda,in.Destino,"0"})    
    appendCSV(data)
    
    ide += 1
        
	return &Confirmation{Body: "Orden confirmada"}, nil
}
 
func (s *Server) OrderPyme(ctx context.Context, in *Pyme) (*Confirmation, error) {
	log.Printf("Id: %s", in.Id)
    log.Printf("Producto: %s", in.Producto)
    log.Printf("Valor: %s", in.Valor)
    log.Printf("Tienda: %s", in.Tienda)
    log.Printf("Destino: %s", in.Destino)
    log.Printf("Prioritario: %s", in.Prioritario)
    log.Printf("")
    
    prio, err := strconv.Atoi(in.Prioritario)
    if err != nil {           
        log.Fatal(err)                                                                                                                                                        
    }    
    var tipo = "normal"
    if prio > 0 {
        tipo = "prioritario"
    }
    
    var data [][]string
    t := time.Now()
    data = append(data, []string{t.Format("2006-01-02 15:04:05"),fmt.Sprintf("%d", ide),tipo,in.Producto,in.Valor,in.Tienda,in.Destino,fmt.Sprintf("%d", seguimiento)})    
    appendCSV(data)
    
    ide += 1
    seguimiento += 1
    
	return &Confirmation{Body: fmt.Sprintf("%d", seguimiento)}, nil
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
