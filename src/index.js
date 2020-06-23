
var express = require('express');
var routes = require('./routes/url');
var config = require('./config.js');
var bodyParser = require('body-parser');

var app = express();

// parse application/x-www-form-urlencoded
app.use(bodyParser.urlencoded({ extended: false }));
// parse application/json
app.use(bodyParser.json());

app.route('/api/')
    .get(routes.Urls)
    .post(routes.addUrl);
app.route('/api/:id')
    .delete(routes.deleteUrl)
    .get(routes.Url);

// Start server
app.listen(config.express.port, function(){
    console.log("Express server listening on port %d in %s mode",
    config.express.port, app.settings.env);
});