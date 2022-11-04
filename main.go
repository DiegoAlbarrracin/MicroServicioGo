package main

import (
	"fmt"

	"net/http"
	"log"
	"encoding/json"
	"io/ioutil"
	"time"
	
	
	
	"github.com/gorilla/mux"
)

var hora = time.Now();

var procesadoPorGoV = "Procesado por Go " + hora.Format("15:04:05") + "hs en /msGo";

//Defino la estructura de la respuesta en JSON
type user struct {
	IdCliente int32 `json:"idCliente"`
	Nombre string  `json:"nombre"`
	Dni string  `json:"dni"`
	Telefono string  `json:"telefono"`
	Email string  `json:"email"`
	ProcesadoPor string  `json:"procesadoPor"`
}

type allUsers []user

var users = allUsers {
	/*{
		IdCliente : "aaaaa",
		Nombre : "b",
		Dni : "c", 
		Telefono : "d", 
		Email : "e", 
		ProcesadoPor : procesadoPorGoV, 
	},*/
}


//routes
func indexRoute(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome to my API");
}

func getToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json"); //tipo de dato
    json.NewEncoder(w).Encode(users)
}



func msGo(w http.ResponseWriter, r *http.Request) {
	//Recibo la informacion que me envia Nodejs
	var infoNode user

	reqBody, err := ioutil.ReadAll(r.Body) //recibo los datos del cliente
	if err != nil {
		fmt.Fprintf(w, "Error al recibir datos en Go")
	}

	//Asigno dicha informacion a la variable infoNode
	json.Unmarshal(reqBody, &infoNode)

	//Al dato recibido de Node lo guardo en el array 
	users = append(users, infoNode)

	//Agrego un nuevo objeto reutilizando los datos pasados por Node
	//cambiando el campo "procesadoPor" y el idCliente
	infoNode.IdCliente =  int32(infoNode.IdCliente) + 1
	infoNode.ProcesadoPor = procesadoPorGoV
	users = append(users, infoNode)

	

	//Respondo al cliente con lo que acaba de crear/enviar desde Node
	w.Header().Set("Content-Type", "application/json"); //tipo de dato
	w.WriteHeader(http.StatusCreated) //si fue correcto
	json.NewEncoder(w).Encode(users) //le envio la info agregada

}







func main()  {

	fmt.Println(procesadoPorGoV);

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/user", getToken)
	router.HandleFunc("/msGo", msGo).Methods("POST")


	log.Fatal(http.ListenAndServe(":7000", router))
	
}