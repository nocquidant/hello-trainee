package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func confServerPort() int {
	port := os.Getenv("PORT")
	if port == "" {
		return 8488
	}
	num, err := strconv.Atoi(port)
	if err != nil {
		log.Println("Error during conversion: ", err)
		return 8488
	}
	return num
}

func handlerHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "It looks like you found a way to run me from a k8s service, well done :)")
}

func main() {
	port := confServerPort()

	log.Printf("HTTP server is running using port: %d\n", port)
	log.Println("Available endpoints are: '/hello' and '/back'")

	http.HandleFunc("/hello", handlerHello)

	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}
