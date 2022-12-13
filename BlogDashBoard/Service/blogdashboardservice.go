package blogdashboardservice

import (
	model "blogging/BlogDashBoard/Model"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Shopify/sarama"
)

var (
	brokersUrl        = []string{"localhost:9092"}
	topic      string = "blogs"
)

type DashBoardService struct {
	httpClient *http.Client
	conn       sarama.SyncProducer
}

func NewDashBoardService() *DashBoardService {
	fmt.Print("NewDashBoardService")

	conn, err := ConnectProducer(brokersUrl)
	if err != nil {
		log.Panicln(err)
	}

	return &DashBoardService{
		httpClient: &http.Client{},
		conn:       conn,
	}
}

func (dbs *DashBoardService) WriteBlog(blog model.Blog) error {
	fmt.Printf("%+v\n", blog)
	requestBody := RequestBody{
		RequestType: "create",
		Body:        blog,
		CallBackUrl: "bttp://localhost:9200",
	}
	data, err := json.Marshal(requestBody)
	if err != nil {
		return err

	}
	return dbs.PushCommentToQueue(topic, data)
}
