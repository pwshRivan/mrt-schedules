package route

type Route struct {
	Nid   string `json:"nid"`
	Title string `json:"title"`
	Path  string `json:"path"`
}

type RouteResponse struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Path string `json:"path"`
}

func (r Route) ToResponse() RouteResponse {
	return RouteResponse{
		Id:   r.Nid,
		Name: r.Title,
		Path: r.Path,
	}
}
