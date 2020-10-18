package main

import (
	"context"
	"log"
    //"fmt"

	"github.com/nchcl/sd/chat"
	"google.golang.org/grpc"
)

func main() {
    var paquete [8]string
    var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

    c := chat.NewChatServiceClient(conn)
	
    response, err := c.Camion(context.Background(), &chat.Tipo{Tipo: 2})
    if err != nil {
        log.Fatalf("Error when calling SayHello: %s", err)
    }
    paquete[0] = response.Idpaquete
    paquete[1] = response.Seguimiento
    paquete[2] = response.Tipo
    paquete[3] = response.Valor
    paquete[4] = response.Intentos
    paquete[5] = response.Estado
    paquete[6] = response.Origen
    paquete[7] = response.Destino
    
    log.Printf("Response from server: %s", response)
    
    if len(response.Idpaquete) == 0 {
        log.Printf("henlo")
    }
    
    //fmt.Println(paquete)
    
}
