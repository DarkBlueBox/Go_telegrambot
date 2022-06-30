package main

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

var dbInfo = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

var f []string
var s []string
var a []string
var d []string
var firstDish = []string{"wow", "ewq"}
var secondDish = []string{"kek", "sd"}
var apitizerDish = []string{"lol", "e"}
var drink = []string{"vodka", "gh"}
var wow = "kek"
var dt time.Time

func main() {
	//part := strings.Split(firstDish, ",")
	fmt.Print(createMenu(firstDish, secondDish, apitizerDish, drink))
}

//func createMenu() error {
//	db, err := sql.Open("postgres", dbInfo)
//	if err != nil {
//		return err
//	}
//	defer db.Close()
//	//dt = time.Now()
//	//convdt = dt.Format("01-02-2006 15:04:05")
//
//	//data, err := db.Prepare("INSERT INTO Menu(id,cafe) VALUES(?,?);")
//	//defer data.Close()
//	//if _, err = data.Exec(1, "vkusno"); err != nil {
//	//	return err
//	//}
//
//	data := `INSERT INTO Menu(id,cafe) VALUES(?,?);`
//
//	if _, err = db.Exec(data, 1, "vkusno"); err != nil {
//		return err
//	}
//
//	return nil
//}

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
	_, err = db.Exec(data, "vkusno", f1, s1, a1, d1, convdt)
	if err != nil {
		return err
	}

	return nil
}

var First_Dish string

var Second_Dish string

var Apitizer_Dish string

var Drinks string

func getCafe() error {
	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	rows, err := db.Query("SELECT First_Dish,Second_Dish,Apitizer_Dish,Drinks FROM menu where cafe = 'vkusno'")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&First_Dish, &Second_Dish, &Apitizer_Dish, &Drinks)
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
