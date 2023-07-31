package model

import "strconv"

type Hotel struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Address  string `json:"address"`
	Price    int    `json:"price"`
	Score    int    `json:"score"`
	Brand    string `json:"brand"`
	City     string `json:"city"`
	StarName string `json:"starName"`
	Business string `json:"business"`
	Location string `json:"location"`
	Pic      string `json:"pic"`
}

type HotelSql struct {
	ID        int    `db:"id"`
	Name      string `db:"name"`
	Address   string `db:"address"`
	Price     int    `db:"price"`
	Score     int    `db:"score"`
	Brand     string `db:"brand"`
	City      string `db:"city"`
	StarName  string `db:"star_name"`
	Business  string `db:"business"`
	Latitude  string `db:"latitude"`
	Longitude string `db:"longitude"`
	Pic       string `db:"pic"`
}

func HotelSqlToHotel(sql HotelSql) *Hotel {
	result := &Hotel{
		ID:       strconv.Itoa(sql.ID),
		Name:     sql.Name,
		Address:  sql.Address,
		Price:    sql.Price,
		Score:    sql.Score,
		Brand:    sql.Brand,
		City:     sql.City,
		StarName: sql.StarName,
		Business: sql.Business,
		Location: sql.Latitude + "," + sql.Longitude,
		Pic:      sql.Pic,
	}
	return result
}
