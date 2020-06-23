const urlExists = require("url-exists");

var config = require(__dirname + '/../config.js');
var thinky = require('thinky')(config.rethinkdb);
var r = thinky.r;
var type = thinky.type;

function handleError(res) {
    return function (error) {
        console.log(error.message);
        return res.send(500, { error: error.message });
    }
}

var Url = thinky.createModel('Url', {
    id: type.string(),
    url: type.string(),
});

Url.ensureIndex("url");

exports.Url = (req, res) => {
    var id = req.params.id;
    Url.get(id).run().then(result => {
        res.json(result.url);
    }).error(handleError(res));
}

exports.addUrl = (req, res) => {
    var newUrl = new Url(req.body);


    urlExists(newUrl.url, function (err, exists) {
        if (err){
            res.json(err);
        }
        if (exists){
        Url.filter({ url: newUrl.url }).run()
            .then((result) => {
                if (result.length > 0) {
                    ids = new Array();
                    result.forEach(element => {
                        ids.push(element.id);
                    });
                    res.json(ids)
                } else {
                    newUrl.save().then((result) => {
                        res.json(
                            result
                        );
                    }).error(handleError(res));
                }
            }).error((error) => {
            });
        } else{
            res.json(null);
        }
    })

}

exports.deleteUrl = (req, res) => {
    var id = req.params.id;

    Url.get(id).delete().run().then((result) => {
        res.json(result);
    }).error(handleError(res));
}

exports.Urls = (req, res) => {
    Url.run().then((result) => {
        res.json(result);
    }).error(handleError(res));
}