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

func removefromCart(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		fmt.Println("removecart")

		params := mux.Vars(req)
		var Id string = params["Id"]
		var user string = params["User"]
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
		err = c.Remove(bson.M{"Id": t, "UserName": user})
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




func queue_send(message string) {
	conn, err := amqp.Dial("amqp://" + rabbitmq_user + ":" + rabbitmq_pass + "@" + rabbitmq_server + ":" + rabbitmq_port + "/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		rabbitmq_queue, // name
		false,          // durable
		false,          // delete when unused
		false,          // exclusive
		false,          // no-wait
		nil,            // arguments
	)
	failOnError(err, "Failed to declare a queue")

	body := message
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	log.Printf(" [x] Sent %s", body)
	failOnError(err, "Failed to publish a message")
}

// Receive Order from Queue to Process
func queue_receive() []string {
	conn, err := amqp.Dial("amqp://" + rabbitmq_user + ":" + rabbitmq_pass + "@" + rabbitmq_server + ":" + rabbitmq_port + "/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		rabbitmq_queue, // name
		false,          // durable
		false,          // delete when usused
		false,          // exclusive
		false,          // no-wait
		nil,            // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name,   // queue
		"orders", // consumer
		true,     // auto-ack
		false,    // exclusive
		false,    // no-local
		false,    // no-wait
		nil,      // args
	)
	failOnError(err, "Failed to register a consumer")
	//fmt.Println("jhgfghjkljhgf",msgs)

	order_ids := make(chan string)
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			order_ids <- string(d.Body)
		}
		close(order_ids)
	}()

	err = ch.Cancel("orders", false)
	if err != nil {
		log.Fatalf("basic.cancel: %v", err)
	}

	var order_ids_array []string
	for n := range order_ids {
		order_ids_array = append(order_ids_array, n)
	}

	return order_ids_array
}


// API Catalog items
