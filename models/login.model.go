package models

import (
	"database/sql"
	"espadafajar/restku/db"
	"espadafajar/restku/helpers"
	"fmt"
)

type User struct {
	Id      int    `'json:"id"`
	Usernam string `json:"usernam"`
}

func CheckLogin(usernam, password string) (bool, error) {
	var obj User
	var pwd string

	con := db.CreateCon()

	sqlstatement := "SELECT * from users where usernam = ?"

	err := con.QueryRow(sqlstatement, usernam).Scan(
		&obj.Id, &obj.Usernam, &pwd,
	)

	if err == sql.ErrNoRows {
		fmt.Println("usernam not found")
		return false, err
	}

	if err != nil {
		fmt.Println("query error")
		return false, err
	}

	match, err := helpers.CheckPasswordHash(password, pwd)

	if !match {
		fmt.Println("usernam dan password salah")
		return false, err
	}

	return true, nil
}
