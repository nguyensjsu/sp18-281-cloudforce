package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"strconv"
	"github.com/satori/go.uuid"
	"github.com/streadway/amqp"
)

// RabbitMQ Config
var rabbitmq_server = "localhost"
var rabbitmq_port = "5672"
var rabbitmq_queue = "cloudforce"
var rabbitmq_user = "guest"
var rabbitmq_pass = "guest"

// MongoDB Config
var mongodb_server = "mongodb://cmpe281:sreedevi@ds251889.mlab.com:51889/cloudforce"
var mongodb_database = "cloudforce"
var mongodb_collection = "Burgers"
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

func getorderdetails(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		// Open MongoDB Session

		fmt.Println("inorderdetailds")

		session, err := mgo.Dial(mongodb_server)
		if err != nil {
			panic(err)
		}
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		c := session.DB(mongodb_database).C(mongodb_collection)

		params := mux.Vars(req)
		var id string = params["Id"]

		t, _ := strconv.ParseInt(id, 10, 0);
		fmt.Println(t)
		// Get Gumball Inventory
		var result bson.M
		err = c.Find(bson.M{"Id": t}).One(&result)
		if err != nil {
			log.Fatal(err)
		}
		// Return Order Status
		fmt.Println(result)
		formatter.JSON(w, http.StatusOK, result)
	}
}

// API Catalog items
