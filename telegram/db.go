package telegram

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"strings"
	"time"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "9880"
	dbname   = "telegrambot"
)

var cafe string
var cafepassword string
var cafename []string
var cafepass []string

var admin string
var adminpassword string
var adminname []string
var adminpass []string

//type firstDish []string
//type secondDish []string
//type apitizerDish []string
//type drink []string
var first string
var second string
var apitizer string
var drinkname string

var firstDish []string
var secondDish []string
var apitizerDish []string
var drink []string

var dt time.Time
var convdt string

var dbInfo = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

//var cafe []string

//Collecting data from bot
//func collectData(username string, chatid int64, message string, answer []string) error {

//Connecting to database
//db, err := sql.Open("postgres", dbInfo)
//if err != nil {
//	return err
//}
//defer db.Close()

//Converting slice with answer to string
//answ := strings.Join(answer, ", ")

//Creating SQL command
//data := `INSERT INTO users(username, chat_id, message, answer) VALUES($1, $2, $3, $4);`

//Execute SQL command in database
//	if _, err = db.Exec(data, `@`+username, chatid, message, answ); err != nil {
//		return err
//	}
//
//	return nil
//}

//Creating users table in database

func createTable() error {

	//Connecting to database
	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		return err
	}
	defer db.Close()

	//Creating users Table
	if _, err = db.Exec(`CREATE TABLE AdminCafe(ID int PRIMARY KEY, Cafe VARCHAR,Password VARCHAR);`); err != nil {
		return err
	}

	return nil
}

func getCafe() error {
	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	rows, err := db.Query("SELECT cafe, password FROM admincafe ")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&cafe, &cafepassword)
		if err != nil {
			panic(err)
		}
		cafename = append(cafename, cafe)
		cafepass = append(cafepass, cafepassword)

	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}

	return nil
}

func getCompany() error {

	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	rows, err := db.Query("SELECT name, password FROM admincompany")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&admin, &adminpassword)
		if err != nil {
			panic(err)
		}

	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	return nil
}

func createMenu(f []string, s []string, a []string, d []string) error {
	dt = time.Now()
	convdt := dt.Format("01-02-2006 15:04:05")
	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		return err
	}
	f1 := strings.Join(f, ",")
	s1 := strings.Join(s, ",")
	a1 := strings.Join(a, ",")
	d1 := strings.Join(d, ", ")
	defer db.Close()
	data := `INSERT INTO Menu(cafe,First_Dish,Second_Dish,Apitizer_Dish,Drinks,created) VALUES ($1,$2,$3,$4,$5,$6);`
	if err != nil {
		return err
	}
	//_, err = db.Exec(`INSERT INTO Menu(cafe,First_Dish,Second_Dish,Apitizer_Dish,Drinks) VALUES('vkusno', f, pq.Array(s), pq.Array(a), pq.Array(d));`)
	//_, err = db.Exec(data, "vkusno", pq.Array(f), pq.Array(s), pq.Array(a), pq.Array(d))
	_, err = db.Exec(data, cafeName, f1, s1, a1, d1, convdt)
	if err != nil {
		return err
	}

	return nil
}
func showMenu() {
	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	rows, err := db.Query("SELECT First_Dish,Second_Dish,Apitizer_Dish,Drinks FROM menu where created > now() - interval '14 hour';")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&first, &second, &apitizer, &drinkname)
		if err != nil {
			panic(err)
		}
		firstDish = append(firstDish, first)
		secondDish = append(secondDish, second)
		apitizerDish = append(apitizerDish, apitizer)
		drink = append(drink, drinkname)

	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}

}

//Getting number of users who using bot
//func getNumberOfUsers() (int64, error) {
//
//	var count int64

//Connecting to database
//db, err := sql.Open("postgres", dbInfo)
//if err != nil {
//	return 0, err
//}
//defer db.Close()

//Counting number of users
//	row := db.QueryRow("SELECT COUNT(DISTINCT username) FROM users;")
//	err = row.Scan(&count)
//	if err != nil {
//		return 0, err
//	}
//
//	return count, nil
//}
