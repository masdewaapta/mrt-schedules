package station

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/masdewaapta/mrt-schedules/common/client"
)

type Service interface {
	GetAllStation() (response []StationResponse, err error)
}

type service struct {
	client *http.Client
}

func NewService() Service {
	return &service{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (s *service) GetAllStation() (response []StationResponse, err error) {
	// Layer Service
	url := "https://www.jakartamrt.co.id/id/val/stasiuns"

	// hit url
	byteReponses, err := client.DoRequest(s.client, url)
	if err != nil {
		return
	}

	var stations []Station
	err = json.Unmarshal(byteReponses, &stations)

	// keluarkan response
	for _, item := range stations {
		response = append(response, StationResponse{
			Id:   item.Id,
			Name: item.Name,
		})
	}
	return
}
