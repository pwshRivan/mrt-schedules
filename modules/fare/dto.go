package fare

type Station struct {
	Nid      string         `json:"nid"`
	Title    string         `json:"title"`
	Estimasi []EstimasiItem `json:"estimasi"`
}

type EstimasiItem struct {
	StasiunNid string `json:"stasiun_nid"`
	Tarif      string `json:"tarif"`
	Waktu      string `json:"waktu"`
}

type Fare struct {
	FromStationId string
	FromStation   string
	ToStationId   string
	ToStation     string
	Amount        int
	Duration      int
}

type FareResponse struct {
	FromStationId string `json:"from_station_id"`
	FromStation   string `json:"from_station"`
	ToStationId   string `json:"to_station_id"`
	ToStation     string `json:"to_station"`
	Amount        int    `json:"amount"`
	Duration      int    `json:"duration"`
}

func (f Fare) ToResponse() FareResponse {
	return FareResponse{
		FromStationId: f.FromStationId,
		FromStation:   f.FromStation,
		ToStationId:   f.ToStationId,
		ToStation:     f.ToStation,
		Amount:        f.Amount,
		Duration:      f.Duration,
	}
}
