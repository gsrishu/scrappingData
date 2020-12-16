# scrappingData
scrapping data online

There is two API's

To run the first api
# go run ascrapping.go

 # (POST) localhost:5000/amazonData

Parameters

amazonUrl:  URL of amazon product

FormData

amazonUrl   https://www.amazon.in/Fancy-Storage-Additional-Exchange-Offers/dp/B089MTKP2P/?_encoding=UTF8&pd_rd_w=mYIXE&pf_rd_p=d67307d4-aeda-4774-98a8-bd9f829b8ea9&pf_rd_r=0JYR73TJG4JKVP89S9KF&pd_rd_r=9a454544-c132-4529-b5fa-edc27fba7869&pd_rd_wg=t9oQ8&ref_=pd_gw_ci_mcx_mr_hp_d

This API will save the name, image, price, etc. of the product in the DataBase.

HEADERS Content-Type:application/json

Sample Response

202 
"SCRAPPING COMPLEATED AND DATA IS SAVED IN DB"

To run the first api
# go run jscrapping.go

#  (POST) localhost:5000/jsonData

Parameters
url: ProductUrl
name: Name of the product
imageUrl: product image
description: product description
price: the price of the product
totelReviews: a total review of the product

Body
{
	"url": "https://www.amazon.com/PlayStation-4-Pro-1TB-Console/dp/B01LOP8EZC/",
	"product": {
		"name": "iphonex",
		"imageURL": "https://images-na.ssl-images-amazon.com/images/I/41GGPRqTZtL._AC_.jpg",
		"description": "Heighten your experiences.\n Enrich your adventures.",
		"price": "$348.00",
		"totalReviews": 4436
	}
}

This API takes the JSON data as input and save the value in DB

Sample Response

202 
"Data Saved in DB"


# Requirements

1. gorilla/mux(for routing)
 installation 
  go get -u github.com/gorilla/mux
  
2. colly framework (for scrapping)
 installation 
 go get -u github.com/gocolly/colly/...

3. postman (for testing the API)
installation 
 https://www.postman.com/downloads/
 
 4. Sqlite3(database)
 go get github.com/mattn/go-sqlite3
 
 go-sqlite3 is a cgo package. If you want to build your app using go-sqlite3, you need GCC. However, after you have built and installed go-sqlite3 with go install github.com/mattn/go-sqlite3 (which requires GCC), you can build your app without relying on GCC in the future.

Important: Because this is a CGO enabled package you are required to set the environment variable CGO_ENABLED=1 and have a GCC compile present within your path.

The data saved in the DB can be seen on the terminal


