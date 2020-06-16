package procesos

import (	
	"net/http"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"

	//"container/list"
	"strings"
    "io/ioutil"
	"log"
	"strconv"
	"os/exec"
)

type Process struct{
	PID 	string `json:"pid,omitempty"`
	Nombre 	string `json:"name,omitempty"`
	Usuario string `json:"user,omitempty"`
	Estado 	string `json:"state,omitempty"`
	RAM 	string `json:"ram,omitempty"`
}

var RunningProcess int = 0;
var SleepingProcess int = 0;
var StopedProcess int = 0;
var ZombieProcess int = 0;
var TotalProcess int = 0;


func ProcessData(w http.ResponseWriter, r *http.Request){
	var RunningP int = 0;
	var SleepingP int = 0;
	var StopedP int = 0;
	var ZombieP int = 0;
	var TotalP int = 0;
	//fmt.Println("GET Process")
	/*			REGISTRANDO TODA LA INFORMACION PARA MOSTRAR		*/
	archivos, err := ioutil.ReadDir("/proc")
    if err != nil {
		log.Fatal(err)
		return
	}

	dataProcess := make(map[string][]Process)		//map con la data a enviar
	var nProcesos int = 0;
	
	var PID string = ""
	var name string = ""
	var user string = ""
	var state string = ""
	var RAM string = ""

    for _, carpetaProceso := range archivos {
		if(carpetaProceso.IsDir()){
			
			numDir, err := strconv.Atoi(carpetaProceso.Name())

			if err == nil {
				numDir++;
				//le incremento para que no me de error por no usar la variable
				PID = carpetaProceso.Name()	//Almacenando el PID (name folder)
				b, err := ioutil.ReadFile("/proc/"+PID+"/status")
				if err != nil {
					return
				}
				str := string(b)
				listaInfo := strings.Split(string(str), "\n")
				
				/*	ARCHIVO STATUS
				Name:	systemd				name
				Umask:	0000		
				State:	S (sleeping)		state
				Tgid:	1
				Ngid:	0
				Pid:	1
				PPid:	0
				TracerPid:	0
				Uid:	0	0	0	0		user
				...
				...
				*/
					
				//Almacenando el Nombre (Name archivo status)
				sepName := strings.Fields(strings.Split(listaInfo[0],":")[1])
				name = sepName[0]
				//Almacenando el Estado (State archivo status)
				sepState := strings.Fields(strings.Split(listaInfo[1],":")[0])
				//fmt.Println(sepState[0])
				
				if(sepState[0] == "State"){
					sepState = strings.Fields(strings.Split(listaInfo[1],":")[1])
					
				}else {
					sepState = strings.Fields(strings.Split(listaInfo[2],":")[1])
					
				}
				state = sepState[0]
				/*contador de estados*/
				if(state == "R"){
					RunningP++
				} else if(state == "S"){
					SleepingP++
				} else if(state == "T"){
					StopedP++
				} else if(state == "Z"){
					ZombieP++
				}else{
					fmt.Print("state raro: ")
					fmt.Println(state)
					continue
				}
				TotalP++;


				//Almacenando el Usuario (Uid archivo status)
				sepUser := strings.Fields(strings.Split(listaInfo[8],":")[1])
				idUser := strings.Fields(sepUser[0])	//Almacenando el id del usuario
				cmd := exec.Command("id", "-un", ""+idUser[0])//alias del usuario: id -un 1000
				nameUser, err := cmd.Output()
				if err != nil {
					println("Error exec command Process: "+err.Error())
					return
				}
				user = string(nameUser)

				//Almacenando el porcentaje de la ram
				//Obteniendo los valores de la memoria
				//fmt.Println(name)
				
				m, err := ioutil.ReadFile("/proc/"+PID+"/statm")
				if err != nil { return }
				strm := string(m)
				listaM := strings.Split(string(strm), " ")
				memUsedKB := 0
				//listaMemUsedKB := strings.Split(listaM[i], " ")
				
				for i := 1; i < 7; i++ {	
					memUTemp, err1 := strconv.Atoi(listaM[i])
					if err1 == nil{
						memUsedKB = memUsedKB + memUTemp
					}else{
						break
					}
				}
				//fmt.Print("memUserKB = ")
				//fmt.Println(memUsedKB)
				//Obtenidendo el total de la ram
				t, err := ioutil.ReadFile("/proc/meminfo")
				if err != nil { return }
				strt := string(t)
				listaTot := strings.Split(string(strt), "\n")
				memoriaTotal := strings.Replace((listaTot[0])[10:24], " ", "", -1)
				ramTotalKB, err1 := strconv.Atoi(memoriaTotal)
				porcUso := 0.0
				//fmt.Println("%ram = "+ramEnUso+"")
				
				if err1 == nil {
					ramEnUso,err3 := strconv.ParseFloat(strconv.Itoa(memUsedKB), 64)
					ramTotal,err4 := strconv.ParseFloat(strconv.Itoa(ramTotalKB), 64)
					porcUso = (ramEnUso/ramTotal)*100.0
				
					if err3 == nil && err4 == nil {
						RAM = fmt.Sprintf("%.2f", porcUso)+"%"
						
					}else{
						RAM = "0%"
					}
				}
				//fmt.Println("%ram: "+RAM)
				

				dataProcess["key"] = append(dataProcess["key"], Process{PID, name ,user,state,RAM})
				nProcesos++;
			}else{
				//fmt.Println("aqui Dx")
				break;
			}
			
			//break
		}else{
			fmt.Println("aqui D:")
			break
			
		}
	}


	processResponse := dataProcess

	jsonResponse, errorjson := json.Marshal(processResponse)

	if errorjson != nil {
		http.Error(w, errorjson.Error(), http.StatusInternalServerError)
		fmt.Println("Send Error Process")
		return
	}
	RunningProcess = RunningP
	SleepingProcess = SleepingP
	StopedProcess = StopedP
	ZombieProcess = ZombieP
	TotalProcess = TotalP

	fmt.Println("Send Process")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}



type Totals struct{
	Running 	int `json:"running"`
	Sleeping 	int `json:"sleeping"`
	Stoped 		int `json:"stoped"`
	Zombie 		int `json:"zombie"` 
	Total 		int `json:"total"` 
}

func ProcessTotal(w http.ResponseWriter, r *http.Request){
	
	totalsResponse := Totals {RunningProcess, SleepingProcess, StopedProcess, ZombieProcess, TotalProcess}
	
	jsonResponse, errorjson := json.Marshal(totalsResponse)

	if errorjson != nil {
		http.Error(w, errorjson.Error(), http.StatusInternalServerError)
		fmt.Println("Send Error Process")
		return
	}

	//fmt.Print("Total: ")
	//fmt.Println(totalsResponse.total)
	fmt.Println("Send Total Process")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

type ResponseResult struct {
	Id string `json:"id"`
}

func KillProcess(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]
	respondWithJSON(response, http.StatusOK, ResponseResult{Id: id})

	fmt.Println("kill process: ", id)
	// EL CODIGO PARA MATAR EL PROCESO
	cmd := exec.Command("kill", "-9", ""+id)
	nameUser, err := cmd.Output()
	if err != nil {
		println("Error Killing Process: "+err.Error())
		return
	}
	print(nameUser)
	

}

type Child struct{
	PIDPadre 	string `json:"pidp,omitempty"`
	NombrePadre	string `json:"namep,omitempty"`
	PID 	string `json:"pid,omitempty"`
	Nombre 	string `json:"name,omitempty"`
}

func ChildProcesses(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Print("get: ")
	fmt.Println(id)

	archivos, err := ioutil.ReadDir("/proc")
    if err != nil {
		log.Fatal(err)
		return
	}
	dataProcess := make(map[string][]Child)		//map con la data a enviar
    for _, carpetaProceso := range archivos {
		if(carpetaProceso.IsDir()){
			
			if(carpetaProceso.Name() != id){
				continue
			}

			//	INFORMACION DEL PROCESO PADRE
			var PIDPadre string = carpetaProceso.Name()
			var namePadre string = ""
			b, err := ioutil.ReadFile("/proc/"+PIDPadre+"/status")
			if err != nil { return }
			
			str := string(b)
			listaInfo := strings.Split(string(str), "\n")
			/*	ARCHIVO STATUS
				Name:	systemd				name
			*/
			sepName := strings.Fields(strings.Split(listaInfo[0],":")[1])
			namePadre = sepName[0]
				
			//	OBTENIENDO LOS PROCESOS HIJOS
			archivosTask, err := ioutil.ReadDir("/proc/"+PIDPadre+"/task")
			if err != nil {
				log.Fatal(err)
				return
			}
			var PIDHijo string = ""
			var nameHijo string = ""
			for _, carpetaTask := range archivosTask {
				if(carpetaTask.IsDir()){
					PIDHijo = carpetaTask.Name()


					child, err := ioutil.ReadFile("/proc/"+PIDPadre+"/task/"+PIDHijo+"/status")
					if err != nil { return }
					
					strHijo := string(child)
					listaInfoHijo := strings.Split(string(strHijo), "\n")
					/*	ARCHIVO STATUS
						Name:	systemd				name
					*/
					sepName := strings.Fields(strings.Split(listaInfoHijo[0],":")[1])
					nameHijo = sepName[0]
					dataProcess["key"] = append(dataProcess["key"], Child{PIDPadre, namePadre, PIDHijo, nameHijo})
				}
			}
			break
		}
	}


	processResponse := dataProcess

	jsonResponse, errorjson := json.Marshal(processResponse)

	if errorjson != nil {
		http.Error(w, errorjson.Error(), http.StatusInternalServerError)
		fmt.Println("Send Error Process")
		return
	}

	fmt.Println("Send Process")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)

}



func respondWithJSON(response http.ResponseWriter, statusCode int, data interface{}) {
	//result, _ := json.Marshal(data) // Puedo retornan la misma data o un boolean
	result, _ := json.Marshal(true)
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(statusCode)
	response.Write(result)
	//fmt.Println("hola")

	
}









/*
package main

import (
    
)

func main() {
    archivos, err := ioutil.ReadDir("/proc")
    if err != nil {
        log.Fatal(err)
	}
	var nProcesos int = 0;
    for _, carpetaProceso := range archivos {
		if(carpetaProceso.IsDir()){
			numDir, error := strconv.Atoi(carpetaProceso.Name())
			if error == nil {
				nProcesos++;
				numDir++;
				fmt.Println("PID:", carpetaProceso.Name())
				//Recorrer los archivos para obtener: 
					//- Nombre del proceso
					//- Usuario
					//- Estado
					//- %RAM
				
			
			}
		}
	}
	fmt.Println("No. procesos: ", nProcesos)

}

*/