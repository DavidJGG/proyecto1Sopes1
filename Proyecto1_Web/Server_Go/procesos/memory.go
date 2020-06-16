package procesos
import (	
	"net/http"
	"io/ioutil"
	"strings"
	"encoding/json"
	"strconv"
	//"fmt"
)

type memStruct struct {
	Total_mem int
	Free_mem int
}

var monitor string
var RAM string


func RAMInfo(w http.ResponseWriter, r *http.Request){
	b, err := ioutil.ReadFile("/proc/meminfo")
	if err != nil {
		return
	}
	str := string(b)
	listaInfo := strings.Split(string(str), "\n")
	//fmt.Println(listaInfo)
	/*
	MemTotal:        8081752 kB
	MemFree:         1859764 kB
	MemAvailable:    3219468 kB
	...
	...
	*/
	memoriaTotal := strings.Replace((listaInfo[0])[10:24], " ", "", -1)
	memoriaLibre := strings.Replace((listaInfo[1])[10:24], " ", "", -1)
	
	ramTotalKB, err1 := strconv.Atoi(memoriaTotal)
	ramFreeKB, err2	 := strconv.Atoi(memoriaLibre) 

	if err1 == nil && err2 == nil {
		ramTotalMB := ramTotalKB / 1024
		ramFreeMB := ramFreeKB / 1024
		
		memResponse := memStruct {ramTotalMB, ramFreeMB}

		jsonResponse, errorjson := json.Marshal(memResponse)

		if errorjson != nil {
			http.Error(w, errorjson.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResponse)	
	} else {
		return
	}
}