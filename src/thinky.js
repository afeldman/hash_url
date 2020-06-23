var thinky = require('thinky');
var config = require(__dirname + '/config.js');

// Initialize thinky
// The most important thing is to initialize the pool of connection
thinky.init({
    host: config.rethinkdb.host,
    port: config.rethinkdb.port,
    db: config.rethinkdb.db
});

exports.thinky = thinky;