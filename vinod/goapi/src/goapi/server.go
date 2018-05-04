package main

import (
	"fmt"
	"log"
	"net/http"
	//"encoding/json"
	"github.com/codegangsta/negroni"
	//"github.com/streadway/amqp"
	"github.com/sqs/goreturns"
	"github.com/unrolled/render"
	//"github.com/satori/go.uuid"
	"gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
)

// MongoDB Config
var mongodb_server = "localhost:27017"
var mongodb_database = "cmpe281project"
var mongodb_collection = "users"



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
	 mx.HandleFunc("/login", loginCheck(formatter)).Methods("POST")
}

// Helper Functions
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

// API validate login credentials. 

func login(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		params := mux.Vars(req)
		// Open MongoDB Session
		//session, err := mgo.Dial(mongodb_server)
		session, err := mgo.DialWithInfo(&mgo.DialInfo{
			Addrs:          []string{"54.201.235.180:27017", "54.69.252.51:27017", "52.38.221.198:27017", "34.212.32.182:27017", "34.217.64.2:27017"},
			Timeout:        60 * time.Second,
			Database:       "cloudforce",
			Username:       "cloudforce",
			Password:       "cmpe281",
			ReplicaSetName: "vkcloud",
		})
		if err != nil {
			panic(err)
		}
		defer session.Close()
		fmt.Println("Session has the live Servers", session.LiveServers())
		session.SetMode(mgo.Eventual, true)
		c := session.DB(mongodb_database).C(mongodb_collection)
		var username string = params["username"]
		fmt.Println(username)
		var password string = params["password"]
		fmt.Println(password)

		// Get user details
		var result []bson.M
		err = c.Find(bson.M{"username": username}).All(&result)
		fmt.Println(result[0])
		m := result[0]
		passwor := m["password"]
		fmt.Println(passwor)
		if passwor == password {
			// Return login success
			formatter.JSON(w, http.StatusOK, result)
		} else {
			if err != nil {
				log.Fatal(err)
			}
			// Return login failure
			formatter.JSON(w, http.StatusBadRequest, result)
		}
	}
}

func signup(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		params := mux.Vars(req)
		// Open MongoDB Session
		//session, err := mgo.Dial(mongodb_server)
		session, err := mgo.DialWithInfo(&mgo.DialInfo{
			Addrs:          []string{"54.201.235.180:27017", "54.69.252.51:27017", "52.38.221.198:27017", "34.212.32.182:27017", "34.217.64.2:27017"},
			Timeout:        60 * time.Second,
			Database:       "cloudforce",
			Username:       "cloudforce",
			Password:       "cmpe281",
			ReplicaSetName: "vkcloud",
		})
		if err != nil {
			panic(err)
		}
		defer session.Close()
		fmt.Println("Session has the live Servers", session.LiveServers())
		session.SetMode(mgo.Eventual, true)

		c := session.DB(mongodb_database).C(mongodb_collection)
		var username string = params["username"]
		fmt.Println(username)
		var password string = params["password"]
		fmt.Println(password)
		var firstname string = params["firstname"]
		fmt.Println(username)
		var lastname string = params["lastname"]
		fmt.Println(password)
		if err = c.Insert(bson.M{"username": username,
			"password":  password,
			"firstname": firstname,
			"lastname":  lastname}); err != nil {
			if err != nil {
				log.Fatal(err)
			}
			//Return signup error
			formatter.JSON(w, http.StatusBadRequest, err)
		}
		// Return signup success
		formatter.JSON(w, http.StatusOK, username)

	}
}

