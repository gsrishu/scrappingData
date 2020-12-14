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

	"github.com/gocolly/colly"
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
	router.HandleFunc("/amazonData", amzData)
	router.HandleFunc("/jsonData", jsonData)
	http.ListenAndServe(":5000", router)

}

func amzData(w http.ResponseWriter, r *http.Request) {

	database, _ := sql.Open("sqlite3", "./amazonProductInfo.db")
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS data (id INTEGER PRIMARY KEY, name VARCHAR , description VARCHAR, price VARCHAR, review VARCHAR, imageUrl VARCHAR, time VARCHAR)")
	statement.Exec()

	w.Header().Set("Content-type", "application/json")
	url := r.FormValue("amazonUrl")

	fmt.Println(url)

	fName := "data.csv"
	//file, err := os.OpenFile(fName,os.O_APPEND|os.O_WRONLY,os.ModeAppend)
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("Not, err : %q", err)
		return
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()
	//writer.Write([]string{"Name", "Description","Price","Review","ImageUrl","Time"})
	c := colly.NewCollector()
	var x string
	var y string

	c.OnHTML("#imgTagWrapperId", func(e *colly.HTMLElement) {

		//log.Printf("Scrapping Compleate")
		x = e.ChildAttr("img", "data-old-hires")
	})
	c.OnHTML("#reviewsMedley", func(e *colly.HTMLElement) {

		log.Printf("Scrapping Compleatettttt")

		y = e.ChildText(".a-color-secondary")

		fmt.Println(e.ChildText("#reviewsMedley[a-color-secondary]"))
	})
	//	y = y[0:4]
	cTime := time.Now().Local()
	c.OnHTML("#ppd", func(e *colly.HTMLElement) {

		// Print link
		//	log.Printf("Scrapping Compleate")
		var testStrings [][]string
		testStrings = append(testStrings, []string{e.ChildText("#productTitle"), e.ChildText("#feature-bullets"), e.ChildText("#priceblock_ourprice"), y, x, cTime.String()})
		writer.WriteAll(testStrings)
		statement, _ = database.Prepare("INSERT INTO data (name, description, price, review,imageUrl, time) VALUES (?, ?, ?, ?, ?, ?)")
		statement.Exec(e.ChildText("#productTitle"), e.ChildText("#feature-bullets"), e.ChildText("#priceblock_ourprice"), y, x, cTime.String())

		rows, _ := database.Query("SELECT id,name, description, price, review,imageUrl, time FROM data")
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

	})
	c.OnRequest(func(e *colly.Request) {
		fmt.Println("Visiting")
	})
	c.Visit(url)

	json.NewEncoder(w).Encode("SCRAPPING COMPLEATED AND DATA IS SAVED IN DB")
}

func jsonData(w http.ResponseWriter, r *http.Request) {


	database, _ := sql.Open("sqlite3", "./amazonProductInfo.db")
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS data (id INTEGER PRIMARY KEY, name VARCHAR , description VARCHAR, price VARCHAR, review VARCHAR, imageUrl VARCHAR, time VARCHAR)")
	statement.Exec()

	var user User
	// json.Unmarshal([]byte(birdJson), &birds)
	// fmt.Printf(bird)
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		panic(err)
	}
	userJson, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}

	fName := "data.csv"
	file, err := os.OpenFile(fName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	//file, err := os.Create(fName)
	return
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
	statement, _ = database.Prepare("INSERT INTO data (name, description, price, review,imageUrl, time) VALUES (?, ?, ?, ?, ?, ?)")
	statement.Exec(user.Product.Name,user.Product.Description, user.Product.Price, user.Product.TotalReview, user.Product.ImageUrl, cTime.String())

	rows, _ := database.Query("SELECT id,name, description, price, review,imageUrl, time FROM data")
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

	fmt.Printf(user.Product.Name)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	//Write json response back to response
	w.Write(userJson)
}
