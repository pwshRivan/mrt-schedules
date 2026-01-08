package fare

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/pwshRivan/mrt-schedules/common/client"
)

type Service interface {
	GetFare(fromStationId string, toStationId string) (FareResponse, error)
	GetAllFares() ([]FareResponse, error)
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

func (s *service) GetAllFares() ([]FareResponse, error) {
	url := "https://www.jakartamrt.co.id/id/val/stasiuns"

	byteResponse, err := client.DoRequest(s.client, url)
	if err != nil {
		return nil, err
	}

	var stations []Station
	err = json.Unmarshal(byteResponse, &stations)
	if err != nil {
		return nil, err
	}

	response := make([]FareResponse, 0)

	// Convert estimasi dari setiap station menjadi fare
	for _, station := range stations {
		for _, estimasi := range station.Estimasi {
			amount, _ := strconv.Atoi(estimasi.Tarif)
			duration, _ := strconv.Atoi(estimasi.Waktu)
			fare := Fare{
				FromStationId: station.Nid,
				FromStation:   station.Title,
				ToStationId:   estimasi.StasiunNid,
				Amount:        amount,
				Duration:      duration,
			}
			response = append(response, fare.ToResponse())
		}
	}

	return response, nil
}

func (s *service) GetFare(fromStationId string, toStationId string) (FareResponse, error) {
	url := "https://www.jakartamrt.co.id/id/val/stasiuns"

	byteResponse, err := client.DoRequest(s.client, url)
	if err != nil {
		return FareResponse{}, err
	}

	var stations []Station
	err = json.Unmarshal(byteResponse, &stations)
	if err != nil {
		return FareResponse{}, err
	}

	// Find the source station
	for _, station := range stations {
		if station.Nid == fromStationId {
			// Find the destination station in estimasi
			for _, estimasi := range station.Estimasi {
				if estimasi.StasiunNid == toStationId {
					amount, _ := strconv.Atoi(estimasi.Tarif)
					duration, _ := strconv.Atoi(estimasi.Waktu)
					fare := Fare{
						FromStationId: station.Nid,
						FromStation:   station.Title,
						ToStationId:   estimasi.StasiunNid,
						Amount:        amount,
						Duration:      duration,
					}
					return fare.ToResponse(), nil
				}
			}
		}
	}

	return FareResponse{}, errors.New("fare not found")
}
