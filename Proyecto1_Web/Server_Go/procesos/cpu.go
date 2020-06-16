package procesos

import (
	"net/http"
	"encoding/json"
	"os/exec"	
	"strings"
	"fmt"
	"strconv"
)

func CPUInfo(w http.ResponseWriter, r *http.Request){

	//Comando: ps -eo %cpu
	app := "ps"
    arg0 := "-eo"
    arg1 := "%cpu"

	//Ejecucion del comando 
    cmd := exec.Command(app, arg0, arg1)
    stringPorcentajes, err := cmd.Output()
    if err != nil {
        println("Error exec command CPU: "+err.Error())
        return
    }
	arreglo := strings.Fields(string(stringPorcentajes))	//Split por saltos de linea
	//fmt.Println(arreglo)

	var resp float64  = 0.0	
	for i := 1; i < len(arreglo); i++ {
		var porcentaje float64 = 0.0
		porcentaje, err := strconv.ParseFloat(arreglo[i],64)
		if(err != nil){
			fmt.Print("Error Conversion en CPU: ")
			fmt.Println(0.0)
			return
		}
		resp = resp + porcentaje;
	}

	jsonResponse, errorjson := json.Marshal(resp)
	if errorjson != nil {
		http.Error(w, errorjson.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
	//fmt.Println(resp)

}

