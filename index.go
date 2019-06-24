package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("/build")))
	//http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Println("Hello World")
	//	http.FileServer(http.Dir("/build"))
	//}))
	Port := GetPort()
	log.Println("Now server is running on port" + Port)
	if err := http.ListenAndServe(Port, nil); err != nil {
		log.Fatal(err)
	}
	//http.Handle("/", http.FileServer(http.Dir("./build")))
}

func GetPort() string {
	var port = os.Getenv("PORT")
	// Set a default port if there is nothing in the environment
	if port == "" {
		port = "3000"
		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}
	return ":" + port
}
