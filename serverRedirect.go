// Web Server
// Redirige a la pgina de Google "Http://google.com"

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main(){
	fmt.Println("Servidor corriendo y escuchando ene l puerto 8080")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Redirect Handler recive un objeto ResponseWriter un apuntador httpe.Request 
	// la URL a la que se dirige y un  status cod de tipo 3xx(StatusMovedPermanently, StatusFound, StatusSeeOther)
	http.Redirect(w, r, "http://google.com", http.StatusMovedPermanently)
}