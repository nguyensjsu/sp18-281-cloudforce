var express = require('express');
var path = require('path');
var logger = require('morgan');
var cookieParser = require('cookie-parser');
var bodyParser = require('body-parser');
var passport = require('passport');
var cors = require('cors');
var createmotboard=require('./routes/createmotboard');
var motboard = require('./routes/motboard');
var routes = require('./routes/index');
var login = require('./routes/login');
var users = require('./routes/users');
var mongoSessionURL = "mongodb://localhost:27017/sessions";
var expressSessions = require("express-session");
var mongoStore = require("connect-mongo")(expressSessions);

var app = express();

// view engine setup
app.set('views', path.join(__dirname, 'views'));
app.set('view engine', 'ejs');

// uncomment after placing your favicon in /public
//app.use(favicon(path.join(__dirname, 'public', 'favicon.ico')));
app.use(logger('dev'));
app.use(bodyParser.json());
app.use(bodyParser.urlencoded({extended: false}));
app.use(cookieParser());
app.use(express.static(path.join(__dirname, 'public')));
app.use(expressSessions({
    secret: "CMPE280_passport",
    resave: false,
    //Forces the session to be saved back to the session store, even if the session was never modified during the request
    saveUninitialized: false, //force to save uninitialized session to db.
    //A session is uninitialized when it is new but not modified.
    duration: 30 * 60 * 1000,
    activeDuration: 5 * 6 * 1000,
    store: new mongoStore({
        url: mongoSessionURL
    })
}));


app.use(function (req, res, next) {
    // Website you wish to allow to connect
    res.setHeader('Access-Control-Allow-Origin', 'http://localhost:8080');

    // Request methods you wish to allow
    res.setHeader('Access-Control-Allow-Methods', 'GET, POST, OPTIONS, PUT, PATCH, DELETE');

    // Request headers you wish to allow
    res.setHeader('Access-Control-Allow-Headers', 'X-Requested-With,content-type');

    // Set to true if you need the website to include cookies in the requests sent
    // to the API (e.g. in case you use sessions)
    res.setHeader('Access-Control-Allow-Credentials', true);

    // Pass to next layer of middleware
    next();
});


app.use(passport.initialize());
app.use('/', routes);
app.use('/users', users);
app.post('/logout', function (req, res) {
    console.log(req.session.user);
    req.session.destroy();
    console.log('Session Destroyed');
    res.status(200).send();
});

var handle_get_menu = function(req, res, next){
    console.log( "In Get Menu" ) ;

    var client = new Client();
    client.get( get_items,
        function(data, response_raw){
            console.log( "In client.Get Menu" ) ;
            jsdata = JSON.parse(data);
            console.log(jsdata.length);
            res.render('menu',{ data: jsdata })

        });

    console.log( "In out Get Menu" ) ;

};


var handle= function(req, res, next){
    console.log( "In Get Menu" ) ;

    var client = new Client();
    client.get( get_items,
        function(data, response_raw){
            console.log(jsdata.length);
            res.render('menu',{ data: jsdata })

        });

    console.log( "In out Get Menu" ) ;

    console.log( "In out Get Menu" ) ;

    console.log( "In out Get Menu" ) ;

    console.log( "In out Get Menu" ) ;

};


bodyParser = require('body-parser').json();
app.post('/login',bodyParser, login);
app.post('/signup', users);

module.exports = app;
