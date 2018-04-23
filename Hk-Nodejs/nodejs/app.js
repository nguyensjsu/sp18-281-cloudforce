
var post_cart = "http://localhost:3000/storeincart";




var crypto = require('crypto');
var fs = require('fs');
var express = require('express');
var Client = require('node-rest-client').Client;
var FormData = require('form-data');

var app = express();
app.use(express.bodyParser());
app.use("/images", express.static(__dirname + '/images'));
app.use("/public", express.static(__dirname + '/public'));

app.set('view engine', 'ejs');


var secretKey = "kwRg54x2Go9iEdl49jFENRM12Mp711QI";

var get_hash = function (state, ts) {


    text = state + "|" + ts + "|" + secretKey;
    hmac = crypto.createHmac("sha256", secretKey);
    hmac.setEncoding('base64');
    hmac.write(text);
    hmac.end();
    hash = hmac.read();
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
    client.post(post_cart + '/' + Id + '/' + Name + '/' + Price + '/' + Path,
        function (data, response_raw) {
            res.redirect('/cart')

        });



}




app.post('/cartorder', handle_cartorder);


console.log("Server running on Port 8080...");

app.listen(8080);
