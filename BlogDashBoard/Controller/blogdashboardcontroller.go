package blogdashboardcontroller

import (
	model "blogging/BlogDashBoard/Model"
	blogdashboardservice "blogging/BlogDashBoard/Service"
	"encoding/json"
	"fmt"
	"net/http"
)

type DashboardControler struct {
	DashBoardService *blogdashboardservice.DashBoardService
}

func NewDashBoardController(DashBoardService *blogdashboardservice.DashBoardService) *DashboardControler {
	return &DashboardControler{
		DashBoardService: DashBoardService,
	}
}

func (dbc *DashboardControler) WriteBlog(w http.ResponseWriter, r *http.Request) {
	var blog model.Blog
	fmt.Printf("---", r.Body)
	err := json.NewDecoder(r.Body).Decode(&blog)
	w.Header().Add("Content-Type", "application/json")
	if err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  400,
			"message": err.Error(),
		})
		return
	}
	err = dbc.DashBoardService.WriteBlog(blog)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  500,
			"message": err.Error(),
		})
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  200,
		"message": "data published successfully",
	})

}
