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
