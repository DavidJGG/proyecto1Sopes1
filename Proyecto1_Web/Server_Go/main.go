package main

import (
	"./procesos"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func main(){
	router := mux.NewRouter()
	// Get para los datos de la memoria
	router.HandleFunc("/memoria", procesos.RAMInfo).Methods("GET")
	// Get para los datos del CPU
	router.HandleFunc("/cpu", procesos.CPUInfo).Methods("GET")
	// Get para los datos de los procesos
	router.HandleFunc("/process", procesos.ProcessData).Methods("GET")
	// Get para los totales de los procesos
	router.HandleFunc("/process/total", procesos.ProcessTotal).Methods("GET")
	// Delete para el KILL del proceso
	router.HandleFunc("/process/{id}", procesos.KillProcess).Methods("GET")
	// Get para los totales de los procesos
	router.HandleFunc("/process/child/{id}", procesos.ChildProcesses).Methods("GET")
	

	err := http.ListenAndServe(":3000", router)
	if err != nil {
		fmt.Println(err)
	}
	
}

