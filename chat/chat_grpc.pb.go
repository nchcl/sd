package chat

import (
    "os"
    "log"
    "encoding/csv"
    "fmt"
    "time"
    "strconv"
    "github.com/streadway/amqp"
    "encoding/json"

	"golang.org/x/net/context"
)

//Este archivo tiene todas las definiciones de las funciones para comunicacion grpc usando protocol buffers
//En la funcion "EnvioTerminado" se encuentra la conexion de rabbitmq al sistema financiero.

type Server struct {
}

//Struct para el JSON
type PaqFinanzas struct {
    Id int
    Tipo string
    Valor int
    Intentos int
    Fecha_entrega string
}

var seguimiento = 130031
var ide = 1

//Colas de paquetes
var retail [][]string
var prioritario [][]string
var normal [][]string

//Registro en memoria de los paquetes y sus estados
var registro [][]string

//Funcion para ordenes de tipo "retail".
//Envia un protocol buffer de tipo "Retail" y devuleve un mensaje simple de confirmacion "Confirmation"
func (s *Server) OrderRetail(ctx context.Context, in *Retail) (*Confirmation, error) {
    var data [][]string
    t := time.Now()
    data = append(data, []string{t.Format("2006-01-02 15:04:05"),fmt.Sprintf("%d",      ide),"retail",in.Producto,in.Valor,in.Tienda,in.Destino,"0"})    
    appendCSV(data)
    
    retail = append(retail, []string{fmt.Sprintf("%d",ide),"0","retail",in.Valor,"0","En bodega",in.Tienda,in.Destino})

    ide += 1    
	return &Confirmation{Body: "Orden confirmada"}, nil
}


//Funcion para ordenes de tipo "pyme".
//Envia un protocol buffer de tipo "Pyme" y devuleve un mensaje simple de confirmacion "Confirmation"
func (s *Server) OrderPyme(ctx context.Context, in *Pyme) (*Confirmation, error) {    
    prio, err := strconv.Atoi(in.Prioritario)
    if err != nil {           
        log.Fatal(err)                                                                                                                                                        
    }    
    var tipo string
    if prio > 0 {
        tipo = "prioritario"
        prioritario = append(prioritario, []string{fmt.Sprintf("%d",ide),fmt.Sprintf("%d", seguimiento),tipo,in.Valor,"En bodega","0",in.Tienda,in.Destino})
    } else {
        tipo = "normal"
        normal = append(normal, []string{fmt.Sprintf("%d",ide),fmt.Sprintf("%d", seguimiento),tipo,in.Valor,"En bodega","0",in.Tienda,in.Destino})
    }
    
    var data [][]string
    t := time.Now()
    data = append(data, []string{t.Format("2006-01-02 15:04:05"),fmt.Sprintf("%d", ide),tipo,in.Producto,in.Valor,in.Tienda,in.Destino,fmt.Sprintf("%d", seguimiento)})    
    appendCSV(data)
    
    var seguimiento_string = fmt.Sprintf("%d", seguimiento)
    ide += 1
    seguimiento += 1
    
	return &Confirmation{Body: seguimiento_string}, nil
}

//Funcion para hacer seguimiento de algun pedido tipo pyme.
//Envia un protocol buffer simple con el codigo de seguimiento (tipo Confirmation) y devuelve el estado de la orden
func (s *Server) Seguimiento(ctx context.Context, in *Confirmation) (*Confirmation, error) {
    var codigo = in.Body
    var estado string

    for _, index := range prioritario {
        if index[1] == codigo {
            estado = index[4]
            goto jump
        }
    }    
    
    for _, index := range normal {
        if index[1] == codigo {
            estado = index[4]
            goto jump
        }
    }
    
    for _, index := range registro {
        if index[1] == codigo {
            estado = index[4]
            break
        }
    }
    
    jump:
	return &Confirmation{Body: estado}, nil
}

//Funcion para asignar un pedido a un camion.
//Envia un protocol buffer "Tipo" con el tipo de camion y devuelve un paquete.
func (s *Server) Camion(ctx context.Context, in *Tipo) (*Paquete, error) {
    var tipo_camion = in.Tipo
    var paquete_camion []string
    
    if tipo_camion == 1 {
        if len(retail) == 0 {
            
            if len(prioritario) > 0 {
                paquete_camion, prioritario = prioritario[0], prioritario[1:]
                goto label
            }
            
            return &Paquete{Idpaquete: "-1"}, nil
        }
        paquete_camion, retail = retail[0], retail[1:]

    } else {
        if len(normal) == 0 && len(prioritario) == 0{
            return &Paquete{Idpaquete: "-1"}, nil
        } else if len(prioritario) == 0 {
            paquete_camion, normal = normal[0], normal[1:]
        } else {
            paquete_camion, prioritario = prioritario[0], prioritario[1:]
        }
    }
    
    label:
    paquete_camion[5] = "En camino"
    registro = append(registro, paquete_camion)
    
	return &Paquete{Idpaquete: paquete_camion[0],Seguimiento: paquete_camion[1],Tipo: paquete_camion[2],Valor: paquete_camion[3],Intentos: paquete_camion[4],Estado: paquete_camion[5], Origen: paquete_camion[6], Destino: paquete_camion[7]}, nil
}

//Funcion que actualiza la orden luego que un camion termine el envio.
//Envia un protocol buffer "Paquete" con los datos del paquete y devuelve un mensaje simple de confirmacion.
//Notar tambien que aqui es donde se utiliza rabbitmq para hacer el envio al sistema financiero.
func (s *Server) EnvioTerminado(ctx context.Context, in *Paquete) (*Confirmation, error) {
    var i int
    for i = 0; i < len(registro); i++ {
        if registro[i][0] == in.Idpaquete {
            registro[i][4] = in.Estado
            registro[i][5] = in.Intentos
            break
        }
    }
    
	conn, err := amqp.Dial("amqp://118:118@dist118:5672/")
	if err != nil {
		fmt.Println("Failed Initializing Broker Connection")
		panic(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"Finanzas",
		false,
		false,
		false,
		false,
		nil,
	)

	fmt.Println(q)
	if err != nil {
		fmt.Println(err)
	}
    
    var paquete_num = "0"
    
	if registro[i][4] == "Recibido" {
        paquete_num = "1"
    }
    
    
    paquete_id, err := strconv.Atoi(registro[i][0])
    if err != nil {           
        log.Fatal(err)                                                                                                                                                        
    }    
    
    paquete_valor, err := strconv.Atoi(registro[i][3])
    if err != nil {           
        log.Fatal(err)                                                                                                                                                        
    }
    
    paquete_intentos, err := strconv.Atoi(registro[i][5])
    if err != nil {           
        log.Fatal(err)                                                                                                                                                        
    }    
       
    //struct del JSON a enviar
    rabbit := &PaqFinanzas {
        Id: paquete_id,
        Tipo: registro[i][2],
        Valor: paquete_valor,
        Intentos: paquete_intentos,
        Fecha_entrega: paquete_num,
    }
	
	JSON, err := json.Marshal(rabbit)
    if err != nil {
        fmt.Println(err)
    }
    
	err = ch.Publish(
		"",
		"Finanzas",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(JSON),
		},
	)

	if err != nil {
		fmt.Println(err)
	}
    
	return &Confirmation{Body: "Envio actualizado"}, nil
}

//Funcion para escribir en un archivo .csv
func appendCSV(data [][]string) {
    var path = "result.csv"                                                                                                                                                   
    f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, os.ModeAppend)     
    
    if err != nil {           
        log.Fatal(err)                                                                                                                                                        
    }  
    
    defer f.Close()                                                                                                                                                           
                                                                                                                                                                                                                                                                                                                                                                                                                        
    w := csv.NewWriter(f)                                                                                                                                                     
    w.WriteAll(data)                                                                                                                                                          
                                                                                                                                                                                
    if err := w.Error(); err != nil {                                                                                                                                         
        log.Fatal(err)                                                                                                                                                        
    }
}
