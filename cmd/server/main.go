package main

import (
	"eoda/intermal/backend"
	"eoda/intermal/db"
	"log"
	"net/http"
)

func main() {
	s := backend.NewServer(db.Connect())
	http.HandleFunc("/api/setContainers", s.SetContainers)
	http.HandleFunc("/api/checkContainers", s.CheckContainers)

	log.Fatal(http.ListenAndServe(":8000", nil))

}
