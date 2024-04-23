package main

import (
	"log"
	"main/internal/db"
	"main/internal/router"
	"main/internal/service"
	"net/http"
)

func main() {
	connection, err := db.GetConnection()
	if err != nil {
		log.Println(err)
		return
	}
	srv := service.InitialSrv(connection)
	handler := router.InitHandler(srv)
	routs := router.GetHandlers(handler)

	server := http.Server{Addr: "localhost:8080", Handler: routs}
	err = server.ListenAndServe()
	if err != nil {
		log.Println(err)
		return
	}

}
