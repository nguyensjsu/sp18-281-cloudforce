package main

import (
	"fmt"
	"log"
	"net/http"
	//"encoding/json"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	//	"github.com/satori/go.uuid"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"strconv"
	"encoding/json"
	"github.com/satori/go.uuid"
)

// MongoDB Config
var mongodb_server = "mongodb://cmpe281:sreedevi@ds251889.mlab.com:51889/cloudforce"
var mongodb_database = "cloudforce"
var mongodb_collection = "Burgers"
var mongodb_collection2 = "cart"

// NewServer configures and returns a Server.
func NewServer() *negroni.Negroni {
	formatter := render.New(render.Options{
		IndentJSON: true,
	})
	n := negroni.Classic()
	mx := mux.NewRouter()
	initRoutes(mx, formatter)
	n.UseHandler(mx)
	return n
}

// API Routes
func initRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/ping", pingHandler(formatter)).Methods("GET")

 mx.HandleFunc("/items", getCatalog(formatter)).Methods("GET")
	mx.HandleFunc("/cart", getcartitems(formatter)).Methods("GET")
	mx.HandleFunc("/storeincart/{Id}/{Name}/{Price}/{Path}", saveOrderInCart(formatter)).Methods("POST")

}

// Helper Functions
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

// API Ping Handler
func pingHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, struct{ Test string }{"API version 1.0 alive!"})
	}
}

// API Catalog items
func getCatalog(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		// Open MongoDB Session
		session, err := mgo.Dial(mongodb_server)
		if err != nil {
			panic(err)
		}
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		c := session.DB(mongodb_database).C(mongodb_collection)

		// Get Gumball Inventory
		var result []Burgers
		err = c.Find(bson.M{}).All(&result)
		fmt.Println(result)
		if err != nil {
			log.Fatal(err)
		}
		// Return Order Status
		formatter.JSON(w, http.StatusOK, result)

	}
}
