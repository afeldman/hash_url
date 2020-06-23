var r = require('rethinkdb');
var config = require('./config.js');

var connection;

// Connect
var connect = function() {
    r.connect({
        host: config.rethinkdb.host,
        port: config.rethinkdb.port,
        db: config.rethinkdb.db
    }, function(error, conn) {
        if (error) throw error;
        connection = conn;
        createDatabase();
    });
}

// Create the database
var createDatabase = function() {
    r.dbCreate(config.rethinkdb.db).run(connection, function(error, result) {
        if (error) console.log(error);
        if ((result != null) && (result.dbs_created === 1)) {
            console.log('Database `blog` created');
        }
        else {
            console.log('Error: Database `blog` not created');
        }
        createUrlTable()
    })
}

// Create the table Post
var createUrlTable = function() {
    r.db(config.rethinkdb.db).tableCreate(config.rethinkdb.table).run(connection, function(error, result) {
        if (error) console.log(error);
    
        if ((result != null) && (result.tables_created === 1)) {
            console.log('Table `url` created');
        }
        else {
            console.log('Error: Table `url` not created');
        }
    });
}

// Create the index postId on the table Comment
var createUrlIndex = function() {
    r.db(config.rethinkdb.db).table(config.rethinkdb.table).indexCreate('url').run(connection, function(error, result) {
        if (error) console.log(error);
    
        if ((result != null) && (result.created === 1)) {
            console.log('Index `url` created on `Comment`');
        }
        else {
            console.log('Error: Index `url` not created');
        }
    });
}

// Insert authors
var insertURL = function(url) {
    r.db(config.rethinkdb.db).table(config.rethinkdb.table).insert(url).run(connection, function(error, result) {
        if (error) console.log(error);
    
        if ((result != null) && (result.errors === 0)) {
            console.log('Added url data');
        }
        else {
            console.log('Error: Failed to add url data.');
        }
    });
}
