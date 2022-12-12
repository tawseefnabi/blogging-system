package blogdashboardservice

import (
	"log"
	"net/http"

	"github.com/Shopify/sarama"
)

var (
	brokersUrl        = []string{"localhosr:9092"}
	topic      string = "blogs"
)

type DashBoardService struct {
	httpClient *http.Client
	conn       sarama.SyncProducer
}

func NewDashBoardService() *DashBoardService {
	conn, err := ConnectProducer(brokersUrl)
	if err != nil {
		log.Panicln(err)
	}

	return &DashBoardService{
		httpClient: &http.Client{},
		conn:       conn,
	}
}
