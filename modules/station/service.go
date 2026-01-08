package station

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/pwshRivan/mrt-schedules/common/client"
)

type Service interface {
	GetAllStation() (response []StationResponse, err error)
	CheckSchedulesByStation(id string) (response []ScheduleResponse, err error)
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

func (s *service) GetAllStation() ([]StationResponse, error) {
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

	response := make([]StationResponse, 0, len(stations))
	for _, item := range stations {
		response = append(response, item.ToResponse())
	}

	return response, nil
}

func (s *service) CheckSchedulesByStation(id string) ([]ScheduleResponse, error) {
	url := "https://www.jakartamrt.co.id/id/val/stasiuns"

	byteResponse, err := client.DoRequest(s.client, url)
	if err != nil {
		return nil, err
	}

	var schedule []Schedule
	err = json.Unmarshal(byteResponse, &schedule)
	if err != nil {
		return nil, err
	}

	// schedule selected by id station
	var scheduleSeletced Schedule
	for _, item := range schedule {
		if item.StationId == id {
			scheduleSeletced = item
			break
		}
	}

	if scheduleSeletced.StationId == "" {
		return nil, errors.New("station not found")
	}

	response, err := ConvertDataToResponse(scheduleSeletced)
	if err != nil {
		return nil, err
	}

	return response, nil
}
func ConvertDataToResponse(schedule Schedule) ([]ScheduleResponse, error) {
	var (
		LebakBulusTripName = "Station Lebak Bulus Grab"
		BundaranHITripName = "Station Bundaran HI Bank DKI"
	)

	scheduleLebakBulus := schedule.ScheduleLebakBulus
	scheduleBundaranHI := schedule.ScheduleBundaranHI

	scheduleLebakBulusParsed, err := ConvertScheduleToTimeFormat(scheduleLebakBulus)
	if err != nil {
		return nil, err
	}
	scheduleBundaranHIParsed, err := ConvertScheduleToTimeFormat(scheduleBundaranHI)
	if err != nil {
		return nil, err
	}

	response := make([]ScheduleResponse, 0)

	// convert to response
	for _, item := range scheduleLebakBulusParsed {
		if item.Format("15:04") > time.Now().Format("15:04") {
			response = append(response, ScheduleResponse{
				StationName: LebakBulusTripName,
				Time:        item.Format("15:04"),
			})
		}
	}

	for _, item := range scheduleBundaranHIParsed {
		if item.Format("15:04") > time.Now().Format("15:04") {
			response = append(response, ScheduleResponse{
				StationName: BundaranHITripName,
				Time:        item.Format("15:04"),
			})
		}
	}

	return response, nil
}

func ConvertScheduleToTimeFormat(schedule string) (response []time.Time, err error) {
	var (
		parsedTime time.Time
		schedules  = strings.Split(schedule, ",")
	)

	for _, item := range schedules {
		trimmedTime := strings.TrimSpace(item)
		if trimmedTime == "" {
			continue
		}
		parsedTime, err = time.Parse("15:04:05", trimmedTime)
		if err != nil {
			err = errors.New("failed to parse schedule time " + trimmedTime)
			return nil, err
		}

		response = append(response, parsedTime)
	}
	return response, nil
}
