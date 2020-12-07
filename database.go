package main

import (
	"database/sql"
	"fmt"
	)
type Bake struct {
	Typename string
	Price string
	Madedate string
	Expdate string }
const (
	DB_USER = "postgres"
	DB_PASSWORD = "postgress"
	DB_NAME = "lab5"
	)
func DbConnect() error {
	var err error
	db,
	err = sql.Open("postgres",
		fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME))
	if err != nil {
		return err
	}
	if _,
	err := db.Exec("CREATE TABLE IF NOT EXISTS bakes (bake_type text,bake_price text,made_date text, exp_days text)");
	err != nil {
		return err }
	return nil
}
func DbAddBake(typename, price, madedate, expdate string)error {
	sqlstmt := "INSERT INTO bakes VALUES ($1, $2, $3, $4)"
	_,
	err := db.Exec(sqlstmt, typename, price, madedate, expdate)
	if err != nil {
		return err
	}
	return nil }

func DbGetBakes() ([]Bake, error) {
	var bakes []Bake
	stmt,
	err := db.Prepare("SELECT bake_type, bake_price, made_date, exp_days FROM bakes")
	if err != nil {
	return bakes, err
	}
	res,
	err := stmt.Query()
	if err != nil {
		return bakes, err
	}
	var tempBake Bake
	for res.Next() {
		err = res.Scan(&tempBake.Typename, &tempBake.Price, &tempBake.Madedate, &tempBake.Expdate)
		if err != nil {
			return bakes, err
		}
		bakes = append(bakes, tempBake) }
	return bakes, err
}

func DbGetMaximum() ([]Bake, error) {
	var bakes []Bake
	stmt,
	err := db.Prepare("SELECT * FROM bakes WHERE bake_price =(SELECT MAX (bake_price) FROM bakes)")
	if err != nil {
		return bakes, err
	}
	res,
	err := stmt.Query()
	if err != nil {
		return bakes, err
	}
	var tempBake Bake
	for res.Next() {
		err = res.Scan(&tempBake.Typename, &tempBake.Price, &tempBake.Madedate, &tempBake.Expdate)
		if err != nil {
			return bakes, err
		}
		bakes = append(bakes, tempBake) }
	return bakes, err
}
