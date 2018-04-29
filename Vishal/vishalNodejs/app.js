/**

 Mighty Gumball, Inc.
 Version 5.0

 - Refactored Previous REST Client Approach to Transaction Based REST API
 - (i.e. instead of the Scaffolded REST API based on Domain Object Annotation)
 - Handlebars Page Templates
 - Client State Validation using HMAC Key-Based Hash

 NodeJS-Enabled Standing Gumball
 Model# M102988
 Serial# 1234998871109

 **/

// added in v3: handlebars
// https://www.npmjs.org/package/express3-handlebars
// npm install express3-handlebars

// added in v2: crypto
// crypto functions:  http://nodejs.org/api/crypto.html


var crypto = require('crypto');
var fs = require('fs');
var express = require('express');
var Client = require('node-rest-client').Client;


var app = express();
app.use(express.bodyParser());
app.use("/images", express.static(__dirname + '/images'));
app.use("/public", express.static(__dirname + '/public'));
// handlebars  = require('express3-handlebars');
// hbs = handlebars.create();
// app.engine('handlebars', hbs.engine);
app.set('view engine', 'ejs');


var secretKey = "kwRg54x2Go9iEdl49jFENRM12Mp711QI";

var get_hash = function (state, ts) {

    // http://nodejs.org/api/crypto.html#crypto_crypto_createhmac_algorithm_key
    text = state + "|" + ts + "|" + secretKey;
    hmac = crypto.createHmac("sha256", secretKey);
    hmac.setEncoding('base64');
    hmac.write(text);
    hmac.end();
    hash = hmac.read();
    //console.log( "HASH: " + hash )
    return hash;

}


var error = function (req, res, msg, ts) {

    var result = new Object();
    state = "error";
    hash = get_hash(state, ts);

    result.msg = msg;
    result.ts = ts;
    result.hash = hash;
    result.state = state;

    res.render('gumball', {
        state: result.state,
        ts: result.ts,
        hash: result.hash,
        message: result.msg
    });

}


var page = function (req, res, state, ts, status) {

    var result = new Object();
    hash = get_hash(state, ts);
    console.log(state);

    var client = new Client();
    var count = "";
    client.get(machine,
        function (data, response_raw) {
            console.log(data);
            //for(var key in data) {
            //    console.log( "key:" + key + ", value:" + data[key] );
            //}
            jsdata = JSON.parse(data)
            for (var key in jsdata) {
                console.log("key:" + key + ", value:" + jsdata[key]);
            }
            count = jsdata.CountGumballs
            console.log("count = " + count);
            var msg = "\n\nMighty Gumball, Inc.\n\nNodeJS-Enabled Standing Gumball\nModel# " +
                jsdata.ModelNumber + "\n" +
                "Serial# " + jsdata.SerialNumber + "\n" +
                "\n" + state + "\n";
            if (status) {
                msg = msg + "\n" + status + "\n\n";
            }
            result.msg = msg;
            result.ts = ts;
            result.hash = hash;
            result.state = state;

            res.render('gumball', {
                state: result.state,
                ts: result.ts,
                hash: result.hash,
                message: result.msg
            });
        });
}


var order = function (req, res, state, ts) {

    var client = new Client();
    var count = 0;
    client.post(endpoint,
        function (data, response_raw) {
            jsdata = JSON.parse(data)
            for (var key in jsdata) {
                console.log("key:" + key + ", value:" + jsdata[key]);
            }
            id = jsdata.Id;
            status = jsdata.OrderStatus;
            console.log("order id: " + id);
            console.log("order status: " + status);
            status_msg = "order id: " + id + " order status: " + status;
            page(req, res, state, ts, status_msg);
        });
}


var handle_cartorder = function (req, res, next) {

    var Id = req.param("id");
    var Name = req.param("name");
    var Price = req.param("price");
    var Path = req.param("path");


    console.log("in cart order");
    console.log(Id);
    console.log(Name);
    console.log(Price);
    console.log(Path);


    var client = new Client();
    client.post(post_cart + '/' + Id + '/' + Name + '/' + Price + '/' + Path + '/' + localStorage.getItem('user'),
        function (data, response_raw) {
            res.redirect('/cart')
        });
}

var handle_deleteCartItem = function (req, res, next) {

    var Id = req.param("Id");
    console.log("In cart delete", Id);
    console.log(Id);

    var client = new Client();
    client.delete(delete_cart_item + '/' + Id,
        function (data, response_raw) {
            res.redirect('/cart')
            // client.get(get_cart,
            //     function (data1, response_raw) {
            //         jsdata = JSON.parse(data1)
            //         console.log(jsdata.length);
            //         res.render('cart', {data: jsdata})
            //     });

        });


}


var handle_get_cart = function (req, res, next) {
    var client = new Client();
    console.log(localStorage.getItem('user') == undefined);

    if (localStorage.getItem('user') == undefined) {
        console.log("No")
    } else {
        client.get(get_cart + '/' + localStorage.getItem('user'),
            function (data, response_raw) {
                jsdata = JSON.parse(data)
                var temp;
                if (localStorage.getItem('user') == undefined) {
                    temp = 'no';
                } else
                    temp = 'yes';
                res.render('cart', {data: jsdata, logged: temp})
            });
    }
};


var handle_get = function (req, res, next) {
    console.log("Get: ...");
    ts = new Date().getTime()
    console.log(ts)
    res.render('home', {title: 'Express'});
}



/*  Handlebars Test using Home template

app.get('/', function (req, res, next) {
    res.render('home', {
        showTitle: true,
        helpers: {
            foo: function () { return 'foo!'; },
            bar: function () { return 'bar!'; }
        }
    });
});

*/

var handle_payment = function (req, res, next) {

    var items = req.param("items");
    var price = req.param("price");

    console.log("in cart order");
    console.log(items);
    console.log(price);

    var formData = new FormData();
    formData.append('Items', items);
    formData.append('Price', price);
    formData.append('OrderStatus', "FGHJ");

    formData.Items = items;
    formData.Price = price;
    formData.OrderStatus = 'Placed';

    var args = {
        data: {Items: items, OrderStatus: 'Placed', Price: price, UserName: 'vishal@gmail.com'},
        headers: {"Content-Type": "application/json"}
    };


    var client = new Client();
    client.post(payment, args,
        function (data, response_raw) {
            console.log(data)
            jsdata = JSON.parse(data)
            res.render('success', {data: jsdata})
            // client.get(get_cart,
            //     function (data1, response_raw) {
            //         jsdata = JSON.parse(data1)
            //         console.log(jsdata.length);
            //         res.redirect('cart', {data: jsdata})
            //     });

        });


}


app.get('/', handle_get);


app.post('/', handle_post);

app.post('/cartorder', handle_cartorder);

app.post('/payment', handle_payment);

console.log("Server running on Port 8080...");

app.listen(8080);


/**

 Mighty Gumball, Inc.

 NodeJS-Enabled Standing Gumball
 Model# M102988
 Serial# 1234998871109

 **/