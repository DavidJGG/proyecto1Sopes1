package main

/*
   #include <stdio.h>
   #include <stdlib.h>
   #include <sys/sysinfo.h>

   struct sysinfo info;

   void llenar(){
       int val = sysinfo(&info);
   }
*/
import "C"
import "fmt"

func main() {
	var megas float64 = 1024 * 1024

	C.llenar()

	var total = float64(C.info.totalram) / megas
	var libre = float64(C.info.freeram) / megas
	var compartida = float64(C.info.sharedram) / megas
	var enBuffer = float64(C.info.bufferram) / megas
	var porUsada = (total - libre + compartida + enBuffer) / total

	fmt.Println("Carné: 201610648")
	fmt.Println("Nomrbe: David González")

	fmt.Printf("Memoria total: %3.2f MB \n", total)
	fmt.Printf("Memoria libre: %3.2f MB \n", libre)
	fmt.Printf("Memoria usada: %3.2f %c \n", porUsada*100, '%')
}
