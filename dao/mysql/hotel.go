package mysql

import (
	"ESLearn/model"
	"database/sql"
	"errors"
)

func GetHotelOne(hotel *model.HotelSql) error {
	id := hotel.ID
	strSql := "select * from hotel where id = ?"
	if err := db.Get(hotel, strSql, id); err != nil {
		if err == sql.ErrNoRows {
			return errors.New("没有这条数据")
		}
		return err
	}
	return nil
}

// GetHotel 批量查询
func GetHotel() ([]model.HotelSql, error) {
	result := make([]model.HotelSql, 0)
	strSql := "select id,name,address,price,score,brand,city,star_name,business,latitude,longitude,pic from hotel"
	if err := db.Select(&result, strSql); err != nil {
		return nil, err
	}
	return result, nil
}
