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
var mongodb_Deals = "deals"

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
	mx.HandleFunc("/order/{Id}", getorderdetails(formatter)).Methods("GET")
	mx.HandleFunc("/deals", getDeals(formatter)).Methods("GET")
	mx.HandleFunc("/cart", getcartitems(formatter)).Methods("GET")
	mx.HandleFunc("/storeincart/{Id}/{Name}/{Price}/{Path}", saveOrderInCart(formatter)).Methods("POST")
	mx.HandleFunc("/remove/{Id}", removefromCart(formatter)).Methods("DELETE")
	mx.HandleFunc("/payment", saveOrderInDB(formatter)).Methods("POST")
}

// Error Functions
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
func getDeals(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		// Open MongoDB Session
		session, err := mgo.Dial(mongodb_server)
		if err != nil {
			panic(err)
		}
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		c := session.DB(mongodb_database).C(mongodb_Deals)

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

func getcartitems(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		// Open MongoDB Session
		session, err := mgo.Dial(mongodb_server)
		if err != nil {
			panic(err)
		}

		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		c := session.DB(mongodb_database).C(mongodb_collection2)

		// Get Gumball Inventory
		var result []Cart
		err = c.Find(bson.M{}).All(&result)
		fmt.Println(result)
		if err != nil {
			log.Fatal(err)
		}
		// Return Order Status
		formatter.JSON(w, http.StatusOK, result)

	}
}

func saveOrderInCart(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		fmt.Println("postcart")
		params := mux.Vars(req)
		var Id string = params["Id"]
		t, _ := strconv.ParseInt(Id, 10, 0);
		var Price string = params["Price"]
		t1, _ := strconv.ParseInt(Price, 10, 0);
		var Name string = params["Name"]
		var Path string = params["Path"]

		var Cart_order = Cart{
			Id:    t,
			Name:  Name,
			Price: t1,
			Path:  Path,
		}

		session, err := mgo.Dial(mongodb_server)
		if err != nil {
			panic(err)
		}
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		c := session.DB(mongodb_database).C(mongodb_collection2)
		fmt.Printf("%T", c)

		err = c.Insert(Cart_order)
		formatter.JSON(w, http.StatusOK, Cart_order)
		//err = c.Insert(ord)
		////fmt.Println("Gumball Machine:", result )
		////formatter.JSON(w, http.StatusOK, result)
		//
		//
		//fmt.Println( "Orders: ", orders )

	}
}

func removefromCart(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		fmt.Println("removecart")

		params := mux.Vars(req)
		var Id string = params["Id"]
		t, _ := strconv.ParseInt(Id, 10, 0);
		//
		//var Price string = params["Price"]
		//t1,_ := strconv.ParseInt(Price, 10, 0);
		//
		//var Name string = params["Name"]
		//
		//
		//
		//var Path string = params["Path"]
		//
		//
		//var Cart_order = Cart {
		//	Id: t,
		//	Name: Name,
		//	Price:t1,
		//	Path:Path,
		//}

		session, err := mgo.Dial(mongodb_server)
		if err != nil {
			panic(err)
		}
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		c := session.DB(mongodb_database).C(mongodb_collection2)
		fmt.Printf("%T", c)
		err = c.Remove(bson.M{"Id": t})
		//err = c.Insert(Cart_order)
		formatter.JSON(w, http.StatusOK, struct{}{})
		//err = c.Insert(ord)
		////fmt.Println("Gumball Machine:", result )
		////formatter.JSON(w, http.StatusOK, result)
		//
		//
		//fmt.Println( "Orders: ", orders )

	}
}

func saveOrderInDB(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		//params := mux.Vars(req)
		//var Id string = params["Items"]
		//
		//fmt.Println(Id)


		var m order
		_ = json.NewDecoder(req.Body).Decode(&m)
		fmt.Println("Update Gumball Inventory To: ", m)

		uuid := uuid.NewV4()
		fmt.Println(uuid)
		//var ord = orderSave{
		//	Id:          uuid.String(),
		//	OrderStatus: "Order Placed",
		//}
		////if orders == nil {
		////	orders = make(map[string]order)
		////}
		////orders[uuid.String()] = ord
		////queue_send(uuid.String())
		//
		//session, err := mgo.Dial(mongodb_server)
		//if err != nil {
		//	panic(err)
		//}
		//defer session.Close()
		//session.SetMode(mgo.Monotonic, true)
		//c := session.DB(mongodb_database).C(mongodb_collection2)
		//fmt.Printf("%T", c)
		////var result []bson.M
		////err = c.Find(bson.M{}).Limit(10).All(&result)
		////if err != nil {
		////	log.Fatal(err)
		////}
		//
		//err = c.Insert(ord)
		////fmt.Println("Gumball Machine:", result )
		////formatter.JSON(w, http.StatusOK, result)
		//
		//fmt.Println("Orders: ", orders)
		//formatter.JSON(w, http.StatusOK, ord)
	}
}
