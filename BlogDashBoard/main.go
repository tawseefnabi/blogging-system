package main

import (
	blogdashboardcontroller "blogging/BlogDashBoard/Controller"
	blogdashboardservice "blogging/BlogDashBoard/Service"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var (
	DashBoardService   *blogdashboardservice.DashBoardService
	DashboardControler *blogdashboardcontroller.DashboardControler
)

func init() {
	fmt.Println("=========================== Service Init ===========================")
	fmt.Println("----------------------------DashBoardService  Init-----------------------------")
	DashBoardService = blogdashboardservice.NewDashBoardService()
	fmt.Println("-----------------------------DashboardControler Init-----------------------------")
	DashboardControler = blogdashboardcontroller.NewDashBoardController(DashBoardService)
	fmt.Println("====================================================================")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/write_blog", DashboardControler.WriteBlog).Methods("POST")
	router.HandleFunc("/api/read_blog", DashboardControler.ReadBlog).Methods("GET")
	srv := &http.Server{
		Addr:    "localhost:8080",
		Handler: router,
		// enforce timeouts for servers you create!
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	fmt.Println("Server running at port", srv.Addr)
	log.Fatal(srv.ListenAndServe())

}
