# MRT Jakarta API

REST API buat data MRT Jakarta. Go + Gin.

## Requirements

- Go 1.21+

## Endpoints

```
GET /v1/api/stations          - semua stasiun
GET /v1/api/stations/:id      - jadwal stasiun

GET /v1/api/trains            - semua kereta
GET /v1/api/trains/:id        - detail kereta

GET /v1/api/routes            - semua rute
GET /v1/api/routes/:id        - detail rute

GET /v1/api/fares             - semua tarif
GET /v1/api/fares/check       - cek tarif (?from=20&to=21)
```

## Run

```bash
go mod tidy
go run main.go
```

Server jalan di `localhost:8080`

## Response

Semua response punya format yang sama:

```json
{
  "success": true,
  "message": "...",
  "data": {}
}
```

Contoh cek tarif:

```bash
curl "localhost:8080/v1/api/fares/check?from=20&to=21"
```

```json
{
  "success": true,
  "message": "successfully get fare",
  "data": {
    "from_station_id": "20",
    "from_station": "Stasiun Lebak Bulus",
    "to_station_id": "21",
    "amount": 4000,
    "duration": 3
  }
}
```

Kalo error:

```json
{
  "success": false,
  "message": "fare not found",
  "data": null
}
```

## Structure

```
├── main.go
├── common/
│   ├── client/
│   └── response/
└── modules/
    ├── station/
    ├── train/
    ├── route/
    └── fare/
```

## Notes

Data dari [jakartamrt.co.id.](https://www.jakartamrt.co.id/id/val/stasiuns) Kalo API mereka berubah, endpoint ini bisa error.

---

MIT
