package train

import "strconv"

type Train struct {
	Nid           string `json:"nid"`
	Title         string `json:"title"`
	Urutan        string `json:"urutan"`
	IsBig         string `json:"isbig"`
	JadwalHiBiasa string `json:"jadwal_hi_biasa"`
	JadwalHiLibur string `json:"jadwal_hi_libur"`
	JadwalLbBiasa string `json:"jadwal_lb_biasa"`
	JadwalLbLibur string `json:"jadwal_lb_libur"`
}

type TrainResponse struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	Order         int    `json:"order"`
	IsBig         int    `json:"is_big"`
	JadwalHiBiasa string `json:"jadwal_hi_biasa"`
	JadwalHiLibur string `json:"jadwal_hi_libur"`
	JadwalLbBiasa string `json:"jadwal_lb_biasa"`
	JadwalLbLibur string `json:"jadwal_lb_libur"`
}

func (t Train) ToResponse() TrainResponse {
	order, _ := strconv.Atoi(t.Urutan)
	isBig, _ := strconv.Atoi(t.IsBig)
	return TrainResponse{
		Id:            t.Nid,
		Name:          t.Title,
		Order:         order,
		IsBig:         isBig,
		JadwalHiBiasa: t.JadwalHiBiasa,
		JadwalHiLibur: t.JadwalHiLibur,
		JadwalLbBiasa: t.JadwalLbBiasa,
		JadwalLbLibur: t.JadwalLbLibur,
	}
}
