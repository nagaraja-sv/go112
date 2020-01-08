package main

import (
	"go112/app"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	//router.NotFound = http.FileServer(http.Dir("./Angular/dist"))
	router.POST("/task", app.CreateTaskHandler)
	//http.ListenAndServe(":8080", router)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
