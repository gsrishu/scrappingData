package main

import (
	"database/sql"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"

)
type Product struct {
	Name        string
	ImageUrl    string
	Description string
	Price       string
	TotalReview string
}
type User struct {
	Url       string
	Product   Product
	CreatedAt time.Time
}

func main() {

	router := mux.NewRouter()
	go router.HandleFunc("/jsonData",data)
	http.ListenAndServe(":5000", router)

}
func data(w http.ResponseWriter, r *http.Request) {


	database, _ := sql.Open("sqlite3", "./jsonProductInfo.db")
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS jsondata (id INTEGER PRIMARY KEY, name VARCHAR , description VARCHAR, price VARCHAR, review VARCHAR, imageUrl VARCHAR, time VARCHAR)")
	statement.Exec()

	var user User
	// json.Unmarshal([]byte(birdJson), &birds)
	// fmt.Printf(bird)
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		panic(err)
	}
	// userJson, err := json.Marshal(user)
	// if err != nil {
	// 	panic(err)
	// }

	fName := "data.csv"
	file, err := os.OpenFile(fName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	//file, err := os.Create(fName)
	if err != nil {

		log.Fatalf("Not, err : %q", err)

		return
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()
	//c := colly.NewCollector()
	cTime := time.Now().Local()
	var testStrings [][]string
	testStrings = append(testStrings, []string{user.Product.Name, user.Product.Description, user.Product.Price, user.Product.TotalReview, user.Product.ImageUrl, cTime.String()})
	writer.WriteAll(testStrings)
	statement, _ = database.Prepare("INSERT INTO jsondata (name, description, price, review,imageUrl, time) VALUES (?, ?, ?, ?, ?, ?)")
	statement.Exec(user.Product.Name,user.Product.Description, user.Product.Price, user.Product.TotalReview, user.Product.ImageUrl, cTime.String())

	rows, _ := database.Query("SELECT id,name, description, price, review,imageUrl, time FROM jsondata")
	var id int
	var name string
	var description string
	var price string
	var review string
	var imageUrl string
	var time string

	for rows.Next() {
		rows.Scan(&id, &name, &description, &price, &review, &imageUrl, &time)
		fmt.Println(strconv.Itoa(id) + ": " + name + " " + description + " " + price + " " + review + " " + imageUrl + " " + time)
	}

	//fmt.Printf(userJson)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	//Write json response back to response
	res := "Data Saved in DB"
	json.NewEncoder(w).Encode(res)
	//w.Write(res)

}