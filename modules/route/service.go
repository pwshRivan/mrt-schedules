package route

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/pwshRivan/mrt-schedules/common/client"
)

type Service interface {
	GetAllRoutes() ([]RouteResponse, error)
	GetRouteById(id string) (RouteResponse, error)
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

func (s *service) GetAllRoutes() ([]RouteResponse, error) {
	url := "https://www.jakartamrt.co.id/id/val/stasiuns"

	byteResponse, err := client.DoRequest(s.client, url)
	if err != nil {
		return nil, err
	}

	var routes []Route
	err = json.Unmarshal(byteResponse, &routes)
	if err != nil {
		return nil, err
	}

	response := make([]RouteResponse, 0, len(routes))
	for _, item := range routes {
		response = append(response, item.ToResponse())
	}

	return response, nil
}

func (s *service) GetRouteById(id string) (RouteResponse, error) {
	url := "https://www.jakartamrt.co.id/id/val/stasiuns"

	byteResponse, err := client.DoRequest(s.client, url)
	if err != nil {
		return RouteResponse{}, err
	}

	var routes []Route
	err = json.Unmarshal(byteResponse, &routes)
	if err != nil {
		return RouteResponse{}, err
	}

	for _, item := range routes {
		if item.Nid == id {
			return item.ToResponse(), nil
		}
	}

	return RouteResponse{}, errors.New("route not found")
}
