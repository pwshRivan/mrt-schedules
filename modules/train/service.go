package train

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/pwshRivan/mrt-schedules/common/client"
)

type Service interface {
	GetAllTrains() ([]TrainResponse, error)
	GetTrainPosition(id string) (TrainResponse, error)
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

func (s *service) GetAllTrains() ([]TrainResponse, error) {
	url := "https://www.jakartamrt.co.id/id/val/stasiuns"

	byteResponse, err := client.DoRequest(s.client, url)
	if err != nil {
		return nil, err
	}

	var trains []Train
	err = json.Unmarshal(byteResponse, &trains)
	if err != nil {
		return nil, err
	}

	response := make([]TrainResponse, 0, len(trains))
	for _, item := range trains {
		response = append(response, item.ToResponse())
	}

	return response, nil
}

func (s *service) GetTrainPosition(id string) (TrainResponse, error) {
	url := "https://www.jakartamrt.co.id/id/val/stasiuns"

	byteResponse, err := client.DoRequest(s.client, url)
	if err != nil {
		return TrainResponse{}, err
	}

	var trains []Train
	err = json.Unmarshal(byteResponse, &trains)
	if err != nil {
		return TrainResponse{}, err
	}

	for _, item := range trains {
		if item.Nid == id {
			return item.ToResponse(), nil
		}
	}

	return TrainResponse{}, errors.New("train not found")
}
