package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func init() {
	fmt.Println("=========================== Service Init ===========================")
	fmt.Println("----------------------------DashBoardService  Init-----------------------------")
	fmt.Println("-----------------------------DashboardControler Init-----------------------------")
	fmt.Println("====================================================================")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/write_blog", nil).Methods("POST")
	srv := &http.Server{
		Addr:    "localhost:8080",
		Handler: router,
		// enforce timeouts for servers you create!
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(srv.ListenAndServe())
}
