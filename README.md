# SD T1
Integrantes:
    Bernabe Garcia, rol: 201773621-6
    Ignacio Figueroa, rol: 201773526-0
    
Consideraciones generales:
    * Los 4 archivos principales que forman el sistema son server.go (logistica), client.go (cliente), camion.go (camiones), finanzas.go (finanzas)
    * Todas las funciones de comunicacion de grpc asi como tambien el cerebro del sistema logistico se encuentran en el archivo chat_grpc.pb.go (dentro de la carpeta chat).
    * En ese mismo archivo, dentro de la funcion "EnvioTerminado" se encuentra la conexion de rabbitmq al sistema financiero.
    * El archivo chat.proto posee las definiciones de los protocol buffers.
    * Para un correcto funcionamiento, el orden de ejecucion deberia ser logistica, cliente, camiones, financiero.

Sistema logistico:
    Instruccions:
        * Iniciar el sistema corriendo el archivo makefile escribiendo "make" en la consola.
        * Iniciar el resto de los sistemas
        
    Consideraciones:
        * Crea el archivo "result.csv" para llevar el registro de las ordenes.

Cliente:
    Instruccions:
        * Iniciar el sistema corriendo el archivo makefile escribiendo "make" en la consola.
        * Ingresar el tiempo de espera entre paquetes en segundos (1,2,3,...)
        * Ingresar el tiempo que demora un envio en segundos (1,2,3,...)
    
    Consideraciones:
        * Existe un 50% de probabilidad que el cliente pida hacer seguimiento de algun paquete aleatorio
        
Camiones:
    Instruccions:
        * Iniciar el sistema corriendo el archivo makefile escribiendo "make" en la consola.
        * Ingresar el tipo de cliente, "1" para "Retail" y "2" para "Pyme"
        * Ingresar el tiempo entre pedidos en segundos (1,2,3,...)
        * Los archivos .csv que se leen son "pymes.csv" y "retail.csv" respectivamente
    
    Consideraciones:
        * Se crean tres archivos .csv para cada camion.
        
Financiero:
    Instruccions:
        * Iniciar el sistema corriendo el archivo makefile escribiendo "make" en la consola.
    
    Consideraciones:
        * Se va mostrando el registro contable cada vez que se actualiza.
    
